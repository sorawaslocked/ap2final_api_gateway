package grpc

import (
	"context"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/adapter/grpc/dto"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
	svc "github.com/sorawaslocked/ap2final_protos_gen/service/movie"
)

type Movie struct {
	client svc.MovieServiceClient
}

func NewMovie(client svc.MovieServiceClient) *Movie {
	return &Movie{
		client: client,
	}
}

func (c *Movie) Create(ctx context.Context, movie model.Movie) (model.Movie, error) {
	res, err := c.client.Create(ctx, dto.ToCreateMovieRequest(movie))
	if err != nil {
		return model.Movie{}, wrapError(err)
	}

	createdMovie := dto.FromBaseMovie(res.Movie)

	return createdMovie, nil
}

func (c *Movie) GetByID(ctx context.Context, id string) (model.Movie, error) {
	res, err := c.client.Get(ctx, dto.ToGetMovieRequest(id))
	if err != nil {
		return model.Movie{}, wrapError(err)
	}

	movie := dto.FromBaseMovie(res.Movie)

	return movie, nil
}

func (c *Movie) GetAll(ctx context.Context) ([]model.Movie, error) {
	res, err := c.client.GetAll(ctx, &svc.GetAllRequest{})
	if err != nil {
		return []model.Movie{}, wrapError(err)
	}

	var movies []model.Movie

	for _, movie := range res.Movies {
		movies = append(movies, dto.FromBaseMovie(movie))
	}

	return movies, nil
}

func (c *Movie) UpdateByID(ctx context.Context, id string, update model.MovieUpdateData) (model.Movie, error) {
	res, err := c.client.Update(ctx, dto.ToUpdateMovieRequest(id, update))
	if err != nil {
		return model.Movie{}, wrapError(err)
	}

	updatedMovie := dto.FromBaseMovie(res.Movie)

	return updatedMovie, nil
}

func (c *Movie) DeleteByID(ctx context.Context, id string) (model.Movie, error) {
	res, err := c.client.Delete(ctx, dto.ToDeleteMovieRequest(id))
	if err != nil {
		return model.Movie{}, wrapError(err)
	}

	deletedMovie := dto.FromBaseMovie(res.Movie)

	return deletedMovie, nil
}
