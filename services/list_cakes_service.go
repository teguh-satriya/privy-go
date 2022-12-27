package services

import (
	"context"

	"github.com/teguh-satriya/privy-go/models"
	repositories "github.com/teguh-satriya/privy-go/repository"
	"github.com/teguh-satriya/privy-go/trouble"
	"google.golang.org/grpc/grpclog"
)

type ListCakesService interface {
	Call(ctx context.Context) (res *ListCakeResult, err error)
}

type ListCakesServiceImpl struct {
	repo   repositories.CakesRepository
	logger grpclog.LoggerV2
}

type ListCakeParams struct{}

type ListCakeResult struct {
	Cakes []models.Cakes
}

func (s *ListCakesServiceImpl) Call(ctx context.Context) (res *ListCakeResult, err error) {
	data, err := s.repo.List(ctx)

	if err != nil {
		s.logger.Errorf("Failed to get cakes: %v", err)
		return nil, trouble.INTERNAL_SERVER_ERROR
	}

	res = &ListCakeResult{
		Cakes: data,
	}

	return res, nil
}

func NewListCakesService(
	repo repositories.CakesRepository,
	logger grpclog.LoggerV2,
) ListCakesService {
	return &ListCakesServiceImpl{
		repo:   repo,
		logger: logger,
	}
}
