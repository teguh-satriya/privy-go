package services

import (
	"context"
	"time"

	repositories "github.com/teguh-satriya/privy-go/repository"
	"github.com/teguh-satriya/privy-go/trouble"
)

type DeleteCakesService interface {
	Call(ctx context.Context, params *DeleteCakeParams) (*DeleteCakeResult, error)
}

type DeleteCakesServiceImpl struct {
	repo repositories.CakesRepository
}

type DeleteCakeParams struct {
	ID int
}

type DeleteCakeResult struct {
	ID          int
	Title       string
	Description string
	Rating      int
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (s *DeleteCakesServiceImpl) Call(ctx context.Context, params *DeleteCakeParams) (res *DeleteCakeResult, err error) {
	cakeData, err := s.repo.Get(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	if cakeData == nil {
		return nil, trouble.CAKE_NOT_FOUND
	}

	err = s.repo.Delete(ctx, params.ID)
	if err != nil {
		return nil, err
	}

	res = &DeleteCakeResult{
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

func NewDeleteCakesService(
	repo repositories.CakesRepository,
) DeleteCakesService {
	return &DeleteCakesServiceImpl{
		repo: repo,
	}
}
