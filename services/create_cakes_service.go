package services

import (
	"context"
	"time"

	"github.com/teguh-satriya/privy-go/models"
	repositories "github.com/teguh-satriya/privy-go/repository"
	"github.com/teguh-satriya/privy-go/trouble"
)

type CreateCakesService interface {
	Call(ctx context.Context, params *CreateCakeParams) (*CreateCakeResult, error)
}

type CreateCakesServiceImpl struct {
	repo repositories.CakesRepository
}

type CreateCakeParams struct {
	Title       string
	Description string
	Rating      int
	Image       string
}

type CreateCakeResult struct {
	ID          int
	Title       string
	Description string
	Rating      int
	Image       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (s *CreateCakesServiceImpl) Call(ctx context.Context, params *CreateCakeParams) (res *CreateCakeResult, err error) {
	cake := &models.Cakes{
		Title:       params.Title,
		Description: params.Description,
		Rating:      params.Rating,
		Image:       params.Image,
	}

	id, err := s.repo.Create(ctx, cake)
	if err != nil {
		return nil, err
	}

	cakeData, err := s.repo.Get(ctx, int(*id))
	if err != nil {
		return nil, err
	}

	if cakeData == nil {
		return nil, trouble.CAKE_NOT_FOUND
	}

	res = &CreateCakeResult{
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

func NewCreateCakesService(
	repo repositories.CakesRepository,
) CreateCakesService {
	return &CreateCakesServiceImpl{
		repo: repo,
	}
}
