package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
)

type Movie struct {
	ID               string   `json:"id"`
	AgeRating        string   `json:"ageRating"`
	PrimaryTitle     string   `json:"primaryTitle"`
	OriginalTitle    string   `json:"originalTitle"`
	ReleaseYear      uint16   `json:"releaseYear"`
	RuntimeInMinutes uint16   `json:"runtimeInMinutes"`
	Genres           []string `json:"genres"`
	IsDeleted        bool     `json:"isDeleted"`
}

type CreateRequest struct {
	AgeRating        string   `json:"ageRating"`
	PrimaryTitle     string   `json:"primaryTitle"`
	OriginalTitle    string   `json:"originalTitle"`
	ReleaseYear      uint16   `json:"releaseYear"`
	RuntimeInMinutes uint16   `json:"runtimeInMinutes"`
	Genres           []string `json:"genres"`
}

type UpdateRequest struct {
	AgeRating        *string  `json:"ageRating"`
	PrimaryTitle     *string  `json:"primaryTitle"`
	OriginalTitle    *string  `json:"originalTitle"`
	ReleaseYear      *uint16  `json:"releaseYear"`
	RuntimeInMinutes *uint16  `json:"runtimeInMinutes"`
	Genres           []string `json:"genres"`
	IsDeleted        *bool    `json:"isDeleted"`
}

func FromCreateRequest(ctx *gin.Context) (model.Movie, error) {
	var req CreateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return model.Movie{}, ErrJSONBinding
	}

	return model.Movie{
		AgeRating:        req.AgeRating,
		PrimaryTitle:     req.PrimaryTitle,
		OriginalTitle:    req.OriginalTitle,
		ReleaseYear:      req.ReleaseYear,
		RuntimeInMinutes: req.RuntimeInMinutes,
		Genres:           req.Genres,
	}, nil
}

func FromUpdateRequest(ctx *gin.Context) (string, model.MovieUpdateData, error) {
	id := ctx.Param("id")

	var req UpdateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return "", model.MovieUpdateData{}, ErrJSONBinding
	}

	return id, model.MovieUpdateData{
		AgeRating:        req.AgeRating,
		PrimaryTitle:     req.PrimaryTitle,
		OriginalTitle:    req.OriginalTitle,
		ReleaseYear:      req.ReleaseYear,
		RuntimeInMinutes: req.RuntimeInMinutes,
		Genres:           req.Genres,
		IsDeleted:        req.IsDeleted,
	}, nil
}

func ToMovie(movie model.Movie) Movie {
	return Movie{
		ID:               movie.ID,
		AgeRating:        movie.AgeRating,
		PrimaryTitle:     movie.PrimaryTitle,
		OriginalTitle:    movie.OriginalTitle,
		ReleaseYear:      movie.ReleaseYear,
		RuntimeInMinutes: movie.RuntimeInMinutes,
		Genres:           movie.Genres,
		IsDeleted:        movie.IsDeleted,
	}
}
