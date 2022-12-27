package server

import (
	"context"

	cakesv1 "github.com/teguh-satriya/privy-go/proto/cakes/v1"
	"github.com/teguh-satriya/privy-go/services"
)

func (s *CakesServer) DeleteCake(ctx context.Context, req *cakesv1.DeleteCakeRequest) (*cakesv1.DeleteCakeResponse, error) {
	_, err := s.deleteCakeService.Call(ctx, &services.DeleteCakeParams{
		ID: int(req.GetId()),
	})
	if err != nil {
		return nil, err
	}

	res := cakesv1.DeleteCakeResponse{}

	return &res, nil
}
