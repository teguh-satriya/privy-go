package server_test

import (
	mockService "github.com/teguh-satriya/privy-go/mocks/services"
)

type MockCakesServer struct {
	createCakesService *mockService.CreateCakesService
	updateCakesService *mockService.UpdateCakesService
	deleteCakesService *mockService.DeleteCakesService
	getCakesService    *mockService.GetCakesService
	listCakesService   *mockService.ListCakesService
}
