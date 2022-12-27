package server

import (
	"context"

	"github.com/teguh-satriya/privy-go/library/troublemaker"
	cakesv1 "github.com/teguh-satriya/privy-go/proto/cakes/v1"
	"github.com/teguh-satriya/privy-go/services"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *CakesServer) CreateCake(ctx context.Context, req *cakesv1.CreateCakeRequest) (*cakesv1.CreateCakeResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, troublemaker.FromValidationError(err)
	}

	params := &services.CreateCakeParams{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Rating:      int(req.GetRating()),
		Image:       req.GetImage(),
	}

	create, err := s.createCakeService.Call(ctx, params)
	if err != nil {
		return nil, err
	}

	res := &cakesv1.CreateCakeResponse{
		Data: &cakesv1.Cake{
			Id:          int32(create.ID),
			Title:       create.Title,
			Description: create.Description,
			Rating:      int32(create.Rating),
			Image:       create.Image,
			CreatedAt:   timestamppb.New(create.CreatedAt),
			UpdatedAt:   timestamppb.New(create.UpdatedAt),
		},
	}

	return res, nil
}
