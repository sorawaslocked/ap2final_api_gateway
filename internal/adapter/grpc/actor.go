package grpc

import (
	"context"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/adapter/grpc/dto"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
	svc "github.com/sorawaslocked/ap2final_protos_gen/service/actor"
)

type Actor struct {
	client svc.ActorServiceClient
}

func NewActor(client svc.ActorServiceClient) *Actor {
	return &Actor{
		client: client,
	}
}

func (c *Actor) Create(ctx context.Context, actor model.Actor) (model.Actor, error) {
	res, err := c.client.Create(ctx, dto.ToCreateActorRequest(actor))
	if err != nil {
		return model.Actor{}, wrapError(err)
	}

	createdActor := dto.FromBaseActor(res.Actor)

	return createdActor, nil
}

func (c *Actor) GetByID(ctx context.Context, id string) (model.Actor, error) {
	res, err := c.client.Get(ctx, dto.ToGetActorRequest(id))
	if err != nil {
		return model.Actor{}, wrapError(err)
	}

	actor := dto.FromBaseActor(res.Actor)

	return actor, nil
}

func (c *Actor) GetAll(ctx context.Context) ([]model.Actor, error) {
	res, err := c.client.GetAll(ctx, &svc.GetAllRequest{})
	if err != nil {
		return []model.Actor{}, wrapError(err)
	}

	var actors []model.Actor

	for _, actor := range res.Actors {
		actors = append(actors, dto.FromBaseActor(actor))
	}

	return actors, nil
}

func (c *Actor) GetAllWithFilter(ctx context.Context, filter model.ActorFilter) ([]model.Actor, error) {
	res, err := c.client.GetAllWithFilter(ctx, dto.ToActorFilterRequest(filter))
	if err != nil {
		return []model.Actor{}, wrapError(err)
	}

	var actors []model.Actor

	for _, actor := range res.Actors {
		actors = append(actors, dto.FromBaseActor(actor))
	}

	return actors, nil
}

func (c *Actor) GetByMovieID(ctx context.Context, movieID string, role *string) ([]model.Actor, error) {
	res, err := c.client.GetByMovieID(ctx, dto.ToGetActorsByMovieIDRequest(movieID, role))
	if err != nil {
		return []model.Actor{}, wrapError(err)
	}

	var actors []model.Actor

	for _, actor := range res.Actors {
		actors = append(actors, dto.FromBaseActor(actor))
	}

	return actors, nil
}

func (c *Actor) UpdateByID(ctx context.Context, id string, update model.ActorUpdateData) (model.Actor, error) {
	res, err := c.client.Update(ctx, dto.ToUpdateActorRequest(id, update))
	if err != nil {
		return model.Actor{}, wrapError(err)
	}

	updatedActor := dto.FromBaseActor(res.Actor)

	return updatedActor, nil
}

func (c *Actor) DeleteByID(ctx context.Context, id string) (model.Actor, error) {
	res, err := c.client.Delete(ctx, dto.ToDeleteActorRequest(id))
	if err != nil {
		return model.Actor{}, wrapError(err)
	}

	deletedActor := dto.FromBaseActor(res.Actor)

	return deletedActor, nil
}
