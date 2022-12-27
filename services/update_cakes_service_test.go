package services_test

import (
	"context"
	"database/sql"
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

type MockUpdateCakesService struct {
	cakeRepository *mockRepo.CakesRepository
	logger         *mockPackage.LoggerV2
}

func TestUpdateCakesService_Call(t *testing.T) {
	type input struct {
		ctx    context.Context
		params *services.UpdateCakeParams
	}

	type output struct {
		result *services.UpdateCakeResult
		err    error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockUpdateCakesService, *input, *output)
	}
	cakeId := 1
	cakeModelBefore := &models.Cakes{
		ID:          cakeId,
		Title:       "test",
		Description: "test desc",
		Rating:      10,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}

	updateParams := &services.UpdateCakeParams{
		ID:          cakeId,
		Title:       sql.NullString{String: "test update", Valid: true},
		Description: sql.NullString{},
		Rating:      sql.NullInt32{Int32: 7, Valid: true},
		Image:       sql.NullString{},
	}

	cakeModelAfter := &models.Cakes{
		ID:          cakeId,
		Title:       "test update",
		Description: "test desc",
		Rating:      7,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
		UpdatedAt:   cakeModelBefore.UpdatedAt,
		CreatedAt:   cakeModelBefore.CreatedAt,
	}

	updateResponse := &services.UpdateCakeResult{
		ID:          cakeId,
		Title:       "test update",
		Description: "test desc",
		Rating:      7,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
		UpdatedAt:   cakeModelBefore.UpdatedAt,
		CreatedAt:   cakeModelBefore.CreatedAt,
	}

	for _, tt := range []testcase{
		{
			name: "ERROR_ON_GET",
			in: &input{
				ctx:    context.Background(),
				params: updateParams,
			},
			out: &output{
				err:    trouble.INTERNAL_SERVER_ERROR,
				result: nil,
			},
			on: func(mc *MockUpdateCakesService, i *input, o *output) {
				mc.cakeRepository.On("Get", i.ctx, cakeId).Return(nil, o.err)
			},
		},
		{
			name: "NOT_FOUND",
			in: &input{
				ctx:    context.Background(),
				params: updateParams,
			},
			out: &output{
				err:    trouble.CAKE_NOT_FOUND,
				result: nil,
			},
			on: func(mc *MockUpdateCakesService, i *input, o *output) {
				mc.cakeRepository.On("Get", i.ctx, cakeId).Return(nil, nil)
			},
		},
		{
			name: "ERROR_ON_UPDATE",
			in: &input{
				ctx:    context.Background(),
				params: updateParams,
			},
			out: &output{
				err:    trouble.INTERNAL_SERVER_ERROR,
				result: nil,
			},
			on: func(mc *MockUpdateCakesService, i *input, o *output) {
				mc.cakeRepository.On("Get", i.ctx, cakeId).Return(cakeModelBefore, nil)
				mc.cakeRepository.On("Update", i.ctx, cakeModelAfter).Return(o.err)
			},
		},
		{
			name: "OK",
			in: &input{
				ctx:    context.Background(),
				params: updateParams,
			},
			out: &output{
				err:    nil,
				result: updateResponse,
			},
			on: func(mc *MockUpdateCakesService, i *input, o *output) {
				mc.cakeRepository.On("Get", i.ctx, cakeId).Return(cakeModelBefore, nil)
				mc.cakeRepository.On("Update", i.ctx, cakeModelAfter).Return(nil)

			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockUpdateCakesService{
				cakeRepository: &mockRepo.CakesRepository{},
				logger:         &mockPackage.LoggerV2{},
			}

			m.logger.On("Warning", mock.Anything)
			m.logger.On("Error", mock.Anything)
			m.logger.On("Errorf", mock.Anything, mock.Anything)

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := services.NewUpdateCakesService(m.cakeRepository, m.logger)
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
