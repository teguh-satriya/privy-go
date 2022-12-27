package server

import (
	"context"

	cakesv1 "github.com/teguh-satriya/privy-go/proto/cakes/v1"
	"github.com/teguh-satriya/privy-go/services"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *CakesServer) GetCake(ctx context.Context, req *cakesv1.GetCakeRequest) (*cakesv1.GetCakeResponse, error) {
	result, err := s.getCakesService.Call(ctx, &services.GetCakeParams{
		ID: int(req.GetId()),
	})
	if err != nil {
		return nil, err
	}

	res := cakesv1.GetCakeResponse{
		Data: &cakesv1.Cake{
			Id:          int32(result.ID),
			Title:       result.Title,
			Description: result.Description,
			Rating:      int32(result.Rating),
			Image:       result.Image,
			CreatedAt:   timestamppb.New(result.CreatedAt),
			UpdatedAt:   timestamppb.New(result.UpdatedAt),
		},
	}

	return &res, nil
}
