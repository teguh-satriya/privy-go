package server

import (
	cakesv1 "github.com/teguh-satriya/privy-go/proto/cakes/v1"
	"github.com/teguh-satriya/privy-go/services"
)

type CakesServer struct {
	cakesv1.UnimplementedCakesServiceServer
	listCakesService  services.ListCakesService
	getCakesService   services.GetCakesService
	createCakeService services.CreateCakesService
	updateCakeService services.UpdateCakesService
	deleteCakeService services.DeleteCakesService
}

type CakesServerSetter func(server *CakesServer)

func NewCakesServer(setters ...CakesServerSetter) *CakesServer {
	s := &CakesServer{}

	for _, set := range setters {
		set(s)
	}

	return s
}

func WithListCakesService(listCakesService services.ListCakesService) CakesServerSetter {
	return func(as *CakesServer) {
		as.listCakesService = listCakesService
	}
}

func WithGetCakesService(getCakesService services.GetCakesService) CakesServerSetter {
	return func(as *CakesServer) {
		as.getCakesService = getCakesService
	}
}

func WithUpdateCakesService(updateCakeService services.UpdateCakesService) CakesServerSetter {
	return func(as *CakesServer) {
		as.updateCakeService = updateCakeService
	}
}

func WithCreateCakesService(createCakeService services.CreateCakesService) CakesServerSetter {
	return func(as *CakesServer) {
		as.createCakeService = createCakeService
	}
}

func WithDeleteCakesService(deleteCakeService services.DeleteCakesService) CakesServerSetter {
	return func(as *CakesServer) {
		as.deleteCakeService = deleteCakeService
	}
}
