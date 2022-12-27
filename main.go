package main

import (
	"context"
	"flag"
	"io/ioutil"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/teguh-satriya/privy-go/config"
	mysql "github.com/teguh-satriya/privy-go/library/driver"
	cakesv1 "github.com/teguh-satriya/privy-go/proto/cakes/v1"
	sql "github.com/teguh-satriya/privy-go/repository/mysql"
	"github.com/teguh-satriya/privy-go/server"
	"github.com/teguh-satriya/privy-go/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/encoding/protojson"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", ":8080", "gRPC server endpoint") // NOTE: grpc server endpoint options
	logger             = grpclog.NewLoggerV2(os.Stdout, ioutil.Discard, ioutil.Discard)       // NOTE: Initialize Logger
)

func main() {
	flag.Parse()

	db, err := mysql.ConnectDB(config.GetDatabaseConnectionString())
	if err != nil {
		logger.Fatalf("Failed to connect to the database: %v", err)
	}

	repo := sql.NewCakesRepository(db, logger)

	listCakesService := services.NewListCakesService(repo)
	getCakesService := services.NewGetCakesService(repo)

	cakesServiceServer := server.NewCakesServer(
		server.WithListCakesService(listCakesService),
		server.WithGetCakesService(getCakesService),
	)

	// NOTE: Initialize gRPC Dial Option
	dialOptions := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// NOTE: Initialize TCP Connection
	tcp, err := net.Listen("tcp", *grpcServerEndpoint)
	if err != nil {
		logger.Fatalf("net.Listen: cannot initialize tcp connection")
	}

	// NOTE: Create gRPC Server
	srv := grpc.NewServer()

	// NOTE: Create Mux Handler
	mux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames:   true,
				AllowPartial:    true,
				EmitUnpopulated: false,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)

	grpclog.SetLoggerV2(logger)

	// NOTE: Setup context, so the requets can be cancelled
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// NOTE: Run grpc server as go routine
	go func() {
		// NOTE: Register internal servers
		cakesv1.RegisterCakesServiceServer(srv, cakesServiceServer)

		srv.Serve(tcp)
	}()

	// NOTE: Start HTTP server (and proxy calls to gRPC server endpoint)
	// NOTE: Regsiter request servers
	err = cakesv1.RegisterCakesServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, dialOptions)
	if err != nil {
		logger.Fatalf(err.Error())
	}

	httpServer := &http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	httpServer.ListenAndServe()
}
