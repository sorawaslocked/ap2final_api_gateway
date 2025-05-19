package usecase

import (
	"context"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
)

type MovieUseCase struct {
	presenter MoviePresenter
}

func NewMovie(presenter MoviePresenter) *MovieUseCase {
	return &MovieUseCase{
		presenter: presenter,
	}
}

func (uc *MovieUseCase) Create(ctx context.Context, movie model.Movie) (model.Movie, error) {
	return uc.presenter.Create(ctx, movie)
}

func (uc *MovieUseCase) GetByID(ctx context.Context, id string) (model.Movie, error) {
	return uc.presenter.GetByID(ctx, id)
}

func (uc *MovieUseCase) GetAll(ctx context.Context) ([]model.Movie, error) {
	return uc.presenter.GetAll(ctx)
}

func (uc *MovieUseCase) UpdateByID(ctx context.Context, id string, update model.MovieUpdateData) (model.Movie, error) {
	return uc.presenter.UpdateByID(ctx, id, update)
}

func (uc *MovieUseCase) DeleteByID(ctx context.Context, id string) (model.Movie, error) {
	return uc.DeleteByID(ctx, id)
}
