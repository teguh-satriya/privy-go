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

type MockDeleteCakesService struct {
	cakeRepository *mockRepo.CakesRepository
	logger         *mockPackage.LoggerV2
}

func TestDeleteCakesService_Call(t *testing.T) {
	type input struct {
		ctx    context.Context
		params *services.DeleteCakeParams
	}

	type output struct {
		result *services.DeleteCakeResult
		err    error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockDeleteCakesService, *input, *output)
	}

	cakeId := 1
	deleteParams := &services.DeleteCakeParams{
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

	deleteResponse := &services.DeleteCakeResult{
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
				params: deleteParams,
			},
			out: &output{
				err:    trouble.INTERNAL_SERVER_ERROR,
				result: nil,
			},
			on: func(mc *MockDeleteCakesService, i *input, o *output) {
				mc.cakeRepository.On("Get", i.ctx, cakeId).Return(nil, o.err)
			},
		},
		{
			name: "NOT_FOUND",
			in: &input{
				ctx:    context.Background(),
				params: deleteParams,
			},
			out: &output{
				err:    trouble.CAKE_NOT_FOUND,
				result: nil,
			},
			on: func(mc *MockDeleteCakesService, i *input, o *output) {
				mc.cakeRepository.On("Get", i.ctx, cakeId).Return(nil, nil)
			},
		},
		{
			name: "ERROR_ON_DELETE",
			in: &input{
				ctx:    context.Background(),
				params: deleteParams,
			},
			out: &output{
				err:    trouble.INTERNAL_SERVER_ERROR,
				result: nil,
			},
			on: func(mc *MockDeleteCakesService, i *input, o *output) {
				mc.cakeRepository.On("Get", i.ctx, cakeId).Return(cakeModel, nil)
				mc.cakeRepository.On("Delete", i.ctx, cakeId).Return(o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				params: deleteParams,
			},
			out: &output{
				err:    nil,
				result: deleteResponse,
			},
			on: func(mc *MockDeleteCakesService, i *input, o *output) {
				mc.cakeRepository.On("Get", i.ctx, cakeId).Return(cakeModel, nil)
				mc.cakeRepository.On("Delete", i.ctx, cakeId).Return(nil)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockDeleteCakesService{
				cakeRepository: &mockRepo.CakesRepository{},
				logger:         &mockPackage.LoggerV2{},
			}

			m.logger.On("Warning", mock.Anything)
			m.logger.On("Error", mock.Anything)
			m.logger.On("Errorf", mock.Anything, mock.Anything)

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := services.NewDeleteCakesService(m.cakeRepository, m.logger)
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
