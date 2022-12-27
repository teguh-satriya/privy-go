package repositories

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/teguh-satriya/privy-go/models"
	repositories "github.com/teguh-satriya/privy-go/repository"
	"google.golang.org/grpc/grpclog"
)

type CakesRepositoryImpl struct {
	db     *sqlx.DB
	logger grpclog.LoggerV2
}

func NewCakesRepository(
	db *sqlx.DB,
	logger grpclog.LoggerV2,
) repositories.CakesRepository {
	return &CakesRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *CakesRepositoryImpl) List(ctx context.Context) ([]models.Cakes, error) {
	cakes := []models.Cakes{}

	rows, err := r.db.QueryxContext(ctx, CAKES_REPOSITORY_LIST_SQL)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		cake := models.Cakes{}

		if err := rows.Scan(
			&cake.ID,
			&cake.Title,
			&cake.Description,
			&cake.Rating,
			&cake.Image,
			&cake.CreatedAt,
			&cake.UpdatedAt,
		); err != nil {
			return nil, err
		}

		cakes = append(cakes, cake)
	}

	return cakes, nil
}

func (r *CakesRepositoryImpl) Get(ctx context.Context, id int) (*models.Cakes, error) {
	var cake *models.Cakes

	rows, err := r.db.QueryxContext(ctx, CAKES_REPOSITORY_GET_SQL, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		cake = new(models.Cakes)
		err := rows.StructScan(&cake)
		if err != nil {
			return nil, err
		}
	}

	if cake == nil {
		return nil, nil
	}

	return cake, nil
}

func (s *CakesRepositoryImpl) Create(ctx context.Context, cake *models.Cakes) (*int64, error) {
	res, err := s.db.NamedExecContext(ctx, CAKES_REPOSITORY_CREATE_SQL, cake)
	if err != nil {
		return nil, err
	}

	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	return &lastId, nil
}

func (s *CakesRepositoryImpl) Update(ctx context.Context, cake *models.Cakes) error {
	_, err := s.db.NamedExecContext(ctx, CAKES_REPOSITORY_UPDATE_SQL, cake)
	if err != nil {
		return err
	}

	return nil
}

func (s *CakesRepositoryImpl) Delete(ctx context.Context, id int) error {
	_, err := s.db.ExecContext(ctx, CAKES_REPOSITORY_DELETE_SQL, id)
	if err != nil {
		return err
	}

	return nil
}
