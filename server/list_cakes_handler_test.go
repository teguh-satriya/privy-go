package server_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	mockService "github.com/teguh-satriya/privy-go/mocks/services"
	"github.com/teguh-satriya/privy-go/models"
	cakesv1 "github.com/teguh-satriya/privy-go/proto/cakes/v1"
	"github.com/teguh-satriya/privy-go/server"
	"github.com/teguh-satriya/privy-go/services"
	"github.com/teguh-satriya/privy-go/trouble"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCakesServer_ListCake(t *testing.T) {
	type input struct {
		ctx context.Context
		req *cakesv1.ListCakesRequest
	}

	type output struct {
		res *cakesv1.ListCakesResponse
		err error
	}

	cakeId := 1
	request := &cakesv1.ListCakesRequest{}

	cakeModel := models.Cakes{
		ID:          cakeId,
		Title:       "Test",
		Description: "Test Desc",
		Rating:      10,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	cakesModel := []models.Cakes{}
	cakesModel = append(cakesModel, cakeModel)
	serviceRes := &services.ListCakeResult{Cakes: cakesModel}

	cakeProto := &cakesv1.Cake{
		Id:          int32(cakeModel.ID),
		Title:       cakeModel.Title,
		Description: cakeModel.Description,
		Rating:      int32(cakeModel.Rating),
		Image:       cakeModel.Image,
		CreatedAt:   timestamppb.New(cakeModel.CreatedAt),
		UpdatedAt:   timestamppb.New(cakeModel.UpdatedAt),
	}
	cakesProto := []*cakesv1.Cake{}
	cakesProto = append(cakesProto, cakeProto)
	result := &cakesv1.ListCakesResponse{
		Data: cakesProto,
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockCakesServer, *input, *output)
	}

	for _, tt := range []testcase{
		{
			name: "INTERNAL_SERVER_ERROR",
			in: &input{
				ctx: context.Background(),
				req: request,
			},
			out: &output{
				err: trouble.INTERNAL_SERVER_ERROR,
				res: nil,
			},
			on: func(cs *MockCakesServer, i *input, o *output) {

				cs.listCakesService.On("Call", i.ctx).Return(nil, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx: context.Background(),
				req: request,
			},
			out: &output{
				err: nil,
				res: result,
			},
			on: func(cs *MockCakesServer, i *input, o *output) {
				cs.listCakesService.On("Call", i.ctx).Return(serviceRes, nil)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockCakesServer{
				listCakesService: &mockService.ListCakesService{},
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := server.NewCakesServer(
				server.WithListCakesService(m.listCakesService),
			)

			res, err := subject.ListCakes(tt.in.ctx, tt.in.req)

			if tt.out.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.out.err, err)
			}

			if tt.out.res != nil {
				assert.NotNil(t, res)
				assert.Equal(t, tt.out.res, res)
			}
		})
	}
}
