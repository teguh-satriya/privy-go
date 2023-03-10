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

type MockCreateCakesService struct {
	cakeRepository *mockRepo.CakesRepository
	logger         *mockPackage.LoggerV2
}

func TestCreateCakesService_Call(t *testing.T) {
	type input struct {
		ctx    context.Context
		params *services.CreateCakeParams
	}

	type output struct {
		result *services.CreateCakeResult
		err    error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockCreateCakesService, *input, *output)
	}

	createParams := &services.CreateCakeParams{
		Title:       "test",
		Description: "test desc",
		Rating:      10,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
	}

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

	createResponse := &services.CreateCakeResult{
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
			name: "ERROR_ON_CREATE",
			in: &input{
				ctx:    context.Background(),
				params: createParams,
			},
			out: &output{
				err:    trouble.INTERNAL_SERVER_ERROR,
				result: nil,
			},
			on: func(mc *MockCreateCakesService, i *input, o *output) {
				cakeModelPost := &models.Cakes{
					Title:       createParams.Title,
					Description: createParams.Description,
					Rating:      createParams.Rating,
					Image:       createParams.Image,
				}

				mc.cakeRepository.On("Create", i.ctx, cakeModelPost).Return(nil, o.err)
			},
		},
		{
			name: "ERROR_ON_GET",
			in: &input{
				ctx:    context.Background(),
				params: createParams,
			},
			out: &output{
				err:    trouble.INTERNAL_SERVER_ERROR,
				result: nil,
			},
			on: func(mc *MockCreateCakesService, i *input, o *output) {
				cakeModelPost := &models.Cakes{
					Title:       createParams.Title,
					Description: createParams.Description,
					Rating:      createParams.Rating,
					Image:       createParams.Image,
				}
				lastInsertCakeID := int64(cakeId)
				mc.cakeRepository.On("Create", i.ctx, cakeModelPost).Return(&lastInsertCakeID, nil)
				mc.cakeRepository.On("Get", i.ctx, cakeId).Return(nil, o.err)
			},
		},
		{
			name: "DATA_NOT_FOUND",
			in: &input{
				ctx:    context.Background(),
				params: createParams,
			},
			out: &output{
				err:    trouble.CAKE_NOT_FOUND,
				result: nil,
			},
			on: func(mc *MockCreateCakesService, i *input, o *output) {
				cakeModelPost := &models.Cakes{
					Title:       createParams.Title,
					Description: createParams.Description,
					Rating:      createParams.Rating,
					Image:       createParams.Image,
				}
				lastInsertCakeID := int64(cakeId)
				mc.cakeRepository.On("Create", i.ctx, cakeModelPost).Return(&lastInsertCakeID, nil)
				mc.cakeRepository.On("Get", i.ctx, cakeId).Return(nil, nil)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				params: createParams,
			},
			out: &output{
				err:    nil,
				result: createResponse,
			},
			on: func(mc *MockCreateCakesService, i *input, o *output) {
				cakeModelPost := &models.Cakes{
					Title:       createParams.Title,
					Description: createParams.Description,
					Rating:      createParams.Rating,
					Image:       createParams.Image,
				}
				lastInsertCakeID := int64(cakeId)
				mc.cakeRepository.On("Create", i.ctx, cakeModelPost).Return(&lastInsertCakeID, nil)
				mc.cakeRepository.On("Get", i.ctx, cakeId).Return(cakeModel, nil)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockCreateCakesService{
				cakeRepository: &mockRepo.CakesRepository{},
				logger:         &mockPackage.LoggerV2{},
			}

			m.logger.On("Warning", mock.Anything)
			m.logger.On("Error", mock.Anything)
			m.logger.On("Errorf", mock.Anything, mock.Anything)

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := services.NewCreateCakesService(m.cakeRepository, m.logger)
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
