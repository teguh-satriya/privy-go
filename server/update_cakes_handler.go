package server

import (
	"context"
	"database/sql"

	"github.com/teguh-satriya/privy-go/library/troublemaker"
	cakesv1 "github.com/teguh-satriya/privy-go/proto/cakes/v1"
	"github.com/teguh-satriya/privy-go/services"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *CakesServer) UpdateCake(ctx context.Context, req *cakesv1.UpdateCakeRequest) (*cakesv1.UpdateCakeResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, troublemaker.FromValidationError(err)
	}

	params := &services.UpdateCakeParams{
		ID:          int(req.Id),
		Title:       sql.NullString{},
		Description: sql.NullString{},
		Rating:      sql.NullInt32{},
		Image:       sql.NullString{},
	}

	if req.GetTitle() != "" {
		params.Title.Valid = true
		params.Title.String = req.GetTitle()
	}

	if req.GetDescription() != "" {
		params.Description.Valid = true
		params.Description.String = req.GetDescription()
	}

	if req.GetRating() != 0 {
		params.Rating.Valid = true
		params.Rating.Int32 = req.GetRating()
	}

	if req.GetImage() != "" {
		params.Image.Valid = true
		params.Image.String = req.GetImage()
	}

	update, err := s.updateCakeService.Call(ctx, params)
	if err != nil {
		return nil, err
	}

	res := &cakesv1.UpdateCakeResponse{
		Data: &cakesv1.Cake{
			Id:          int32(update.ID),
			Title:       update.Title,
			Description: update.Description,
			Rating:      int32(update.Rating),
			Image:       update.Image,
			CreatedAt:   timestamppb.New(update.CreatedAt),
			UpdatedAt:   timestamppb.New(update.UpdatedAt),
		},
	}

	return res, nil
}
