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
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCakesServer_GetCake(t *testing.T) {
	type input struct {
		ctx context.Context
		req *cakesv1.GetCakeRequest
	}

	type output struct {
		res *cakesv1.GetCakeResponse
		err error
	}

	cakeId := 1
	request := &cakesv1.GetCakeRequest{
		Id: int32(cakeId),
	}

	params := &services.GetCakeParams{
		ID: int(request.Id),
	}

	serviceRes := &services.GetCakeResult{
		ID:          cakeId,
		Title:       "Test",
		Description: "Test Desc",
		Rating:      10,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result := &cakesv1.GetCakeResponse{
		Data: &cakesv1.Cake{
			Id:          int32(serviceRes.ID),
			Title:       serviceRes.Title,
			Description: serviceRes.Description,
			Rating:      int32(serviceRes.Rating),
			Image:       serviceRes.Image,
			CreatedAt:   timestamppb.New(serviceRes.CreatedAt),
			UpdatedAt:   timestamppb.New(serviceRes.UpdatedAt),
		},
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

				cs.getCakesService.On("Call", i.ctx, params).Return(nil, o.err)
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
				cs.getCakesService.On("Call", i.ctx, params).Return(serviceRes, nil)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockCakesServer{
				getCakesService: &mockService.GetCakesService{},
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := server.NewCakesServer(
				server.WithGetCakesService(m.getCakesService),
			)

			res, err := subject.GetCake(tt.in.ctx, tt.in.req)

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
