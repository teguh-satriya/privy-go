package server_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	mockService "github.com/teguh-satriya/privy-go/mocks/services"
	cakesv1 "github.com/teguh-satriya/privy-go/proto/cakes/v1"
	"github.com/teguh-satriya/privy-go/server"
	"github.com/teguh-satriya/privy-go/services"
	"github.com/teguh-satriya/privy-go/trouble"
)

func TestCakesServer_DeleteCake(t *testing.T) {
	type input struct {
		ctx context.Context
		req *cakesv1.DeleteCakeRequest
	}

	type output struct {
		res *cakesv1.DeleteCakeResponse
		err error
	}

	cakeId := 1
	request := &cakesv1.DeleteCakeRequest{
		Id: int32(cakeId),
	}

	params := &services.DeleteCakeParams{
		ID: int(request.Id),
	}

	serviceRes := &services.DeleteCakeResult{
		ID:          cakeId,
		Title:       "Test",
		Description: "Test Desc",
		Rating:      10,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result := &cakesv1.DeleteCakeResponse{}

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

				cs.deleteCakesService.On("Call", i.ctx, params).Return(nil, o.err)
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
				cs.deleteCakesService.On("Call", i.ctx, params).Return(serviceRes, nil)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockCakesServer{
				deleteCakesService: &mockService.DeleteCakesService{},
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := server.NewCakesServer(
				server.WithDeleteCakesService(m.deleteCakesService),
			)

			res, err := subject.DeleteCake(tt.in.ctx, tt.in.req)

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
