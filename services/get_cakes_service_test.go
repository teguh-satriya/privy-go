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

type MockGetCakesService struct {
	cakeRepository *mockRepo.CakesRepository
	logger         *mockPackage.LoggerV2
}

func TestGetCakesService_Call(t *testing.T) {
	type input struct {
		ctx    context.Context
		params *services.GetCakeParams
	}

	type output struct {
		result *services.GetCakeResult
		err    error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockGetCakesService, *input, *output)
	}

	cakeId := 1
	getParams := &services.GetCakeParams{
		ID: cakeId,
	}

	cakeModel := &models.Cakes{
		ID:          cakeId,
		Title:       "test",
		Description: "test desc",
		Rating:      10,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}

	getResponse := &services.GetCakeResult{
		ID:          cakeId,
		Title:       "test",
		Description: "test desc",
		Rating:      10,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
		UpdatedAt:   cakeModel.UpdatedAt,
		CreatedAt:   cakeModel.CreatedAt,
	}

	for _, tt := range []testcase{
		{
			name: "ERROR_ON_GET",
			in: &input{
				ctx:    context.Background(),
				params: getParams,
			},
			out: &output{
				err:    trouble.INTERNAL_SERVER_ERROR,
				result: nil,
			},
			on: func(mc *MockGetCakesService, i *input, o *output) {
				mc.cakeRepository.On("Get", i.ctx, cakeId).Return(nil, o.err)
			},
		},
		{
			name: "NOT_FOUND",
			in: &input{
				ctx:    context.Background(),
				params: getParams,
			},
			out: &output{
				err:    trouble.CAKE_NOT_FOUND,
				result: nil,
			},
			on: func(mc *MockGetCakesService, i *input, o *output) {
				mc.cakeRepository.On("Get", i.ctx, cakeId).Return(nil, nil)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				params: getParams,
			},
			out: &output{
				err:    nil,
				result: getResponse,
			},
			on: func(mc *MockGetCakesService, i *input, o *output) {

				mc.cakeRepository.On("Get", i.ctx, cakeId).Return(cakeModel, nil)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockGetCakesService{
				cakeRepository: &mockRepo.CakesRepository{},
				logger:         &mockPackage.LoggerV2{},
			}

			m.logger.On("Warning", mock.Anything)
			m.logger.On("Error", mock.Anything)
			m.logger.On("Errorf", mock.Anything, mock.Anything)

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := services.NewGetCakesService(m.cakeRepository, m.logger)
			result, err := subject.Call(tt.in.ctx, tt.in.params)

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
