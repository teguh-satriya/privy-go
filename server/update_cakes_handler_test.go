package server_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/teguh-satriya/privy-go/library/troublemaker"
	mockService "github.com/teguh-satriya/privy-go/mocks/services"
	cakesv1 "github.com/teguh-satriya/privy-go/proto/cakes/v1"
	"github.com/teguh-satriya/privy-go/server"
	"github.com/teguh-satriya/privy-go/services"
	"github.com/teguh-satriya/privy-go/trouble"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestCakesServer_UpdateCake(t *testing.T) {
	type input struct {
		ctx context.Context
		req *cakesv1.UpdateCakeRequest
	}

	type output struct {
		res *cakesv1.UpdateCakeResponse
		err error
	}

	type testcase struct {
		name string
		in   *input
		out  *output
		on   func(*MockCakesServer, *input, *output)
	}

	cakeId := 1
	request := &cakesv1.UpdateCakeRequest{
		Id:          int32(cakeId),
		Title:       "test",
		Description: "test desc",
		Rating:      1,
		Image:       "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
	}

	params := &services.UpdateCakeParams{
		ID:          int(request.Id),
		Title:       sql.NullString{Valid: true, String: request.Title},
		Description: sql.NullString{Valid: true, String: request.Description},
		Rating:      sql.NullInt32{Valid: true, Int32: request.Rating},
		Image:       sql.NullString{Valid: true, String: request.Image},
	}

	serviceRes := &services.UpdateCakeResult{
		ID:          cakeId,
		Title:       request.Title,
		Description: request.Description,
		Rating:      int(request.Rating),
		Image:       request.Image,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	result := &cakesv1.UpdateCakeResponse{
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

	for _, tt := range []testcase{
		{
			name: "BAD_REQUEST",
			in: &input{
				ctx: context.Background(),
				req: &cakesv1.UpdateCakeRequest{
					Id:          int32(cakeId),
					Title:       "test fail",
					Description: "test desc",
					Rating:      1,
					Image:       "tes123 fail url",
				},
			},
			out: &output{},
			on: func(cs *MockCakesServer, i *input, o *output) {
				o.err = troublemaker.FromValidationError(i.req.Validate())
			},
		},
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

				cs.updateCakesService.On("Call", i.ctx, params).Return(nil, o.err)
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
				cs.updateCakesService.On("Call", i.ctx, params).Return(serviceRes, nil)
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			m := &MockCakesServer{
				updateCakesService: &mockService.UpdateCakesService{},
			}

			if tt.on != nil {
				tt.on(m, tt.in, tt.out)
			}

			subject := server.NewCakesServer(
				server.WithUpdateCakesService(m.updateCakesService),
			)

			res, err := subject.UpdateCake(tt.in.ctx, tt.in.req)

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
