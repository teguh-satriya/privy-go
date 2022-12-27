package services

import (
	"context"
	"time"

	repositories "github.com/teguh-satriya/privy-go/repository"
	"github.com/teguh-satriya/privy-go/trouble"
)

type GetCakesService interface {
	Call(ctx context.Context, params *GetCakeParams) (*GetCakeResult, error)
}

type GetCakesServiceImpl struct {
	repo repositories.CakesRepository
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
		return nil, err
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
) GetCakesService {
	return &GetCakesServiceImpl{
		repo: repo,
	}
}
