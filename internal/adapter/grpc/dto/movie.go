package dto

import (
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
	"github.com/sorawaslocked/ap2final_protos_gen/base"
	svc "github.com/sorawaslocked/ap2final_protos_gen/service/movie"
)

func FromBaseMovie(res *base.Movie) model.Movie {
	return model.Movie{
		ID:               res.ID,
		AgeRating:        res.AgeRating,
		PrimaryTitle:     res.PrimaryTitle,
		OriginalTitle:    res.OriginalTitle,
		ReleaseYear:      uint16(res.ReleaseYear),
		RuntimeInMinutes: uint16(res.RuntimeInMinutes),
		Genres:           res.Genres,
		CreatedAt:        res.CreatedAt.AsTime(),
		UpdatedAt:        res.UpdatedAt.AsTime(),
		IsDeleted:        res.IsDeleted,
	}
}

func ToCreateMovieRequest(movie model.Movie) *svc.CreateRequest {
	return &svc.CreateRequest{
		AgeRating:        movie.AgeRating,
		PrimaryTitle:     movie.PrimaryTitle,
		OriginalTitle:    movie.OriginalTitle,
		ReleaseYear:      uint32(movie.ReleaseYear),
		RuntimeInMinutes: uint32(movie.RuntimeInMinutes),
		Genres:           movie.Genres,
	}
}

func ToGetMovieRequest(id string) *svc.GetRequest {
	return &svc.GetRequest{
		ID: id,
	}
}

func ToUpdateMovieRequest(id string, update model.MovieUpdateData) *svc.UpdateRequest {
	var releaseYear, runtimeInMinutes *uint32

	if update.ReleaseYear != nil {
		value := uint32(*update.ReleaseYear)
		releaseYear = &value
	}

	if update.RuntimeInMinutes != nil {
		value := uint32(*update.RuntimeInMinutes)
		runtimeInMinutes = &value
	}

	return &svc.UpdateRequest{
		ID:               id,
		AgeRating:        update.AgeRating,
		PrimaryTitle:     update.PrimaryTitle,
		OriginalTitle:    update.OriginalTitle,
		ReleaseYear:      releaseYear,
		RuntimeInMinutes: runtimeInMinutes,
		Genres:           update.Genres,
		IsDeleted:        update.IsDeleted,
	}
}

func ToDeleteMovieRequest(id string) *svc.DeleteRequest {
	return &svc.DeleteRequest{
		ID: id,
	}
}
