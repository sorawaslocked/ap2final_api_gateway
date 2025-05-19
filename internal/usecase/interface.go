package usecase

import (
	"context"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
)

type MoviePresenter interface {
	Create(ctx context.Context, movie model.Movie) (model.Movie, error)
	GetByID(ctx context.Context, id string) (model.Movie, error)
	GetAll(ctx context.Context) ([]model.Movie, error)
	UpdateByID(ctx context.Context, id string, update model.MovieUpdateData) (model.Movie, error)
	DeleteByID(ctx context.Context, id string) (model.Movie, error)
}
