package services_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	mockPackage "github.com/teguh-satriya/privy-go/mocks/package"
	mockRepo "github.com/teguh-satriya/privy-go/mocks/repository"
	"github.com/teguh-satriya/privy-go/models"
	"github.com/teguh-satriya/privy-go/services"
	"github.com/teguh-satriya/privy-go/trouble"
)

type MockListCakesService struct {
	cakeRepository *mockRepo.CakesRepository
	logger         *mockPackage.LoggerV2
}

func TestListCakesService_Call(t *testing.T) {
	type input struct {
		ctx    context.Context
		params *services.ListCakeParams
	}

	type output struct {
		result *services.ListCakeResult
		err    error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockListCakesService, *input, *output)
	}

	listParams := &services.ListCakeParams{}

	cakeId := 1
	cakeModel := &models.Cakes{
		ID:          cakeId,
		Title:       "test",
		Description: "test desc",
		Rating:      10,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}

	listCakeModels := []models.Cakes{}
	listCakeModels = append(listCakeModels, *cakeModel)

	listResponse := &services.ListCakeResult{
		Cakes: listCakeModels,
	}

	for _, tt := range []testcase{
		{
			name: "ERROR_ON_GET",
			in: &input{
				ctx:    context.Background(),
				params: listParams,
			},
			out: &output{
				err:    trouble.INTERNAL_SERVER_ERROR,
				result: nil,
			},
			on: func(mc *MockListCakesService, i *input, o *output) {
				mc.cakeRepository.On("List", i.ctx).Return(nil, o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				params: listParams,
			},
			out: &output{
				err:    nil,
				result: listResponse,
			},
			on: func(mc *MockListCakesService, i *input, o *output) {

				mc.cakeRepository.On("List", i.ctx).Return(listCakeModels, nil)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockListCakesService{
				cakeRepository: &mockRepo.CakesRepository{},
				logger:         &mockPackage.LoggerV2{},
			}

			m.logger.On("Warning", mock.Anything)
			m.logger.On("Error", mock.Anything)
			m.logger.On("Errorf", mock.Anything, mock.Anything)

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := services.NewListCakesService(m.cakeRepository, m.logger)
			result, err := subject.Call(tt.in.ctx)

			if tt.out.err != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.out.err, err)
			}

			if tt.out.result != nil {
				assert.NotNil(t, result)
				assert.Equal(t, tt.out.result, result)
			}
		})
	}
}
