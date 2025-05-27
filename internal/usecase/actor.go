package usecase

import (
	"context"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
)

type ActorUseCase struct {
	presenter ActorPresenter
}

func NewActor(presenter ActorPresenter) *ActorUseCase {
	return &ActorUseCase{
		presenter: presenter,
	}
}

func (uc *ActorUseCase) Create(ctx context.Context, actor model.Actor) (model.Actor, error) {
	return uc.presenter.Create(ctx, actor)
}

func (uc *ActorUseCase) GetByID(ctx context.Context, id string) (model.Actor, error) {
	return uc.presenter.GetByID(ctx, id)
}

func (uc *ActorUseCase) GetAll(ctx context.Context) ([]model.Actor, error) {
	return uc.presenter.GetAll(ctx)
}

func (uc *ActorUseCase) GetAllWithFilter(ctx context.Context, filter model.ActorFilter) ([]model.Actor, error) {
	return uc.presenter.GetAllWithFilter(ctx, filter)
}

func (uc *ActorUseCase) GetByMovieID(ctx context.Context, movieID string, role *string) ([]model.Actor, error) {
	return uc.presenter.GetByMovieID(ctx, movieID, role)
}

func (uc *ActorUseCase) UpdateByID(ctx context.Context, id string, update model.ActorUpdateData) (model.Actor, error) {
	return uc.presenter.UpdateByID(ctx, id, update)
}

func (uc *ActorUseCase) DeleteByID(ctx context.Context, id string) (model.Actor, error) {
	return uc.presenter.DeleteByID(ctx, id)
}
