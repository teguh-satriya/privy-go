package services

import (
	"context"
	"time"

	repositories "github.com/teguh-satriya/privy-go/repository"
	"github.com/teguh-satriya/privy-go/trouble"
	"google.golang.org/grpc/grpclog"
)

type GetCakesService interface {
	Call(ctx context.Context, params *GetCakeParams) (*GetCakeResult, error)
}

type GetCakesServiceImpl struct {
	repo   repositories.CakesRepository
	logger grpclog.LoggerV2
}

type GetCakeParams struct {
	ID int
}

type GetCakeResult struct {
	ID          int
	Title       string
	Description string
	Rating      int
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (s *GetCakesServiceImpl) Call(ctx context.Context, params *GetCakeParams) (res *GetCakeResult, err error) {
	data, err := s.repo.Get(ctx, params.ID)

	if err != nil {
		s.logger.Errorf("Failed to get cakes: %v", err)
		return nil, trouble.INTERNAL_SERVER_ERROR
	}

	if data == nil {
		return nil, trouble.CAKE_NOT_FOUND
	}

	res = &GetCakeResult{
		ID:          data.ID,
		Title:       data.Title,
		Description: data.Description,
		Rating:      data.Rating,
		Image:       data.Image,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}

	return res, nil
}

func NewGetCakesService(
	repo repositories.CakesRepository,
	logger grpclog.LoggerV2,
) GetCakesService {
	return &GetCakesServiceImpl{
		repo:   repo,
		logger: logger,
	}
}
