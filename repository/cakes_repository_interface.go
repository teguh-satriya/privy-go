package repositories

import (
	"context"

	"github.com/teguh-satriya/privy-go/models"
)

type CakesRepository interface {
	Get(ctx context.Context, id int) (*models.Cakes, error)
	List(ctx context.Context) ([]models.Cakes, error)
}
