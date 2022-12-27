package services

import (
	"context"

	"github.com/teguh-satriya/privy-go/models"
	repositories "github.com/teguh-satriya/privy-go/repository"
)

type ListCakesService interface {
	Call(ctx context.Context) (res *ListCakeResult, err error)
}

type ListCakesServiceImpl struct {
	repo repositories.CakesRepository
}

type ListCakeParams struct{}

type ListCakeResult struct {
	Cakes []models.Cakes
}

func (s *ListCakesServiceImpl) Call(ctx context.Context) (res *ListCakeResult, err error) {
	data, err := s.repo.List(ctx)

	if err != nil {
		return nil, err
	}

	res = &ListCakeResult{
		Cakes: data,
	}

	return res, nil
}

func NewListCakesService(
	repo repositories.CakesRepository,
) ListCakesService {
	return &ListCakesServiceImpl{
		repo: repo,
	}
}
