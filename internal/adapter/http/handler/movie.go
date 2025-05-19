package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/adapter/http"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/adapter/http/dto"
	"log/slog"
)

type Movie struct {
	log *slog.Logger
	uc  MovieUseCase
}

func NewMovie(log *slog.Logger, uc MovieUseCase) *Movie {
	return &Movie{
		log: log,
		uc:  uc,
	}
}

func (h *Movie) Create(ctx *gin.Context) {
	const op = "handler.Movie.Create"
	log := h.log.With("op", op)

	movie, err := dto.FromCreateRequest(ctx)
	if err != nil {
		handleError(ctx, err)
		logError(h.log, err)

		return
	}

	newMovie, err := h.uc.Create(ctx, movie)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	http.Created(ctx, dto.ToMovie(newMovie))
}

func (h *Movie) Get(ctx *gin.Context) {
	const op = "handler.Movie.Get"
	log := h.log.With("op", op)

	id := ctx.Param("id")

	movie, err := h.uc.GetByID(ctx, id)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	http.Ok(ctx, dto.ToMovie(movie))
}

func (h *Movie) GetAll(ctx *gin.Context) {
	const op = "handler.Movie.GetAll"
	log := h.log.With("op", op)

	movies, err := h.uc.GetAll(ctx)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	var movieDtos []dto.Movie
	for _, movie := range movies {
		movieDtos = append(movieDtos, dto.ToMovie(movie))
	}

	http.Ok(ctx, movieDtos)
}

func (h *Movie) Update(ctx *gin.Context) {
	const op = "handler.Movie.Update"
	log := h.log.With("op", op)

	id, update, err := dto.FromUpdateRequest(ctx)
	if err != nil {
		handleError(ctx, err)
		logError(log, err)

		return
	}

	movie, err := h.uc.UpdateByID(ctx, id, update)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	http.Ok(ctx, dto.ToMovie(movie))
}

func (h *Movie) Delete(ctx *gin.Context) {
	const op = "handler.Movie.Delete"
	log := h.log.With("op", op)

	id := ctx.Param("id")

	movie, err := h.uc.DeleteByID(ctx, id)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	http.Ok(ctx, dto.ToMovie(movie))
}
