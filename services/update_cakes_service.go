package services

import (
	"context"
	"database/sql"
	"time"

	repositories "github.com/teguh-satriya/privy-go/repository"
	"github.com/teguh-satriya/privy-go/trouble"
	"google.golang.org/grpc/grpclog"
)

type UpdateCakesService interface {
	Call(ctx context.Context, params *UpdateCakeParams) (*UpdateCakeResult, error)
}

type UpdateCakesServiceImpl struct {
	repo   repositories.CakesRepository
	logger grpclog.LoggerV2
}

type UpdateCakeParams struct {
	ID          int
	Title       sql.NullString
	Description sql.NullString
	Rating      sql.NullInt32
	Image       sql.NullString
}

type UpdateCakeResult struct {
	ID          int
	Title       string
	Description string
	Rating      int
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (s *UpdateCakesServiceImpl) Call(ctx context.Context, params *UpdateCakeParams) (res *UpdateCakeResult, err error) {
	cakeData, err := s.repo.Get(ctx, params.ID)
	if err != nil {
		s.logger.Errorf("Failed to get cake: %v", err)
		return nil, trouble.INTERNAL_SERVER_ERROR
	}

	if cakeData == nil {
		return nil, trouble.CAKE_NOT_FOUND
	}

	if params.Title.Valid {
		cakeData.Title = params.Title.String
	}

	if params.Description.Valid {
		cakeData.Description = params.Description.String
	}

	if params.Rating.Valid {
		cakeData.Rating = int(params.Rating.Int32)
	}

	if params.Image.Valid {
		cakeData.Image = params.Image.String
	}

	err = s.repo.Update(ctx, cakeData)
	if err != nil {
		s.logger.Errorf("Failed to update cake: %v", err)
		return nil, trouble.INTERNAL_SERVER_ERROR
	}

	res = &UpdateCakeResult{
		ID:          cakeData.ID,
		Title:       cakeData.Title,
		Description: cakeData.Description,
		Rating:      cakeData.Rating,
		Image:       cakeData.Image,
		CreatedAt:   cakeData.CreatedAt,
		UpdatedAt:   cakeData.UpdatedAt,
	}

	return res, nil

}

func NewUpdateCakesService(
	repo repositories.CakesRepository,
	logger grpclog.LoggerV2,
) UpdateCakesService {
	return &UpdateCakesServiceImpl{
		repo:   repo,
		logger: logger,
	}
}
