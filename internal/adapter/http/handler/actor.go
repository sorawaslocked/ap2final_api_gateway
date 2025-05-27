package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/adapter/http/dto"
	"log/slog"
)

type Actor struct {
	log *slog.Logger
	uc  ActorUseCase
}

func NewActor(log *slog.Logger, uc ActorUseCase) *Actor {
	return &Actor{
		log: log,
		uc:  uc,
	}
}

func (h *Actor) Create(ctx *gin.Context) {
	const op = "handler.Actor.Create"
	log := h.log.With("op", op)

	actor, err := dto.FromCreateActorRequest(ctx)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	newActor, err := h.uc.Create(ctx, actor)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	Created(ctx, dto.ToActor(newActor))
}

func (h *Actor) Get(ctx *gin.Context) {
	const op = "handler.Actor.Get"
	log := h.log.With("op", op)

	id := ctx.Param("id")

	actor, err := h.uc.GetByID(ctx, id)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	Ok(ctx, dto.ToActor(actor))
}

func (h *Actor) GetAll(ctx *gin.Context) {
	const op = "handler.Actor.GetAll"
	log := h.log.With("op", op)

	actors, err := h.uc.GetAll(ctx)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	var actorDtos []dto.Actor
	for _, actor := range actors {
		actorDtos = append(actorDtos, dto.ToActor(actor))
	}

	Ok(ctx, actorDtos)
}

func (h *Actor) GetAllWithFilter(ctx *gin.Context) {
	const op = "handler.Actor.GetAllWithFilter"
	log := h.log.With("op", op)

	filter := dto.FromActorFilterRequest(ctx)

	actors, err := h.uc.GetAllWithFilter(ctx, filter)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	var actorDtos []dto.Actor
	for _, actor := range actors {
		actorDtos = append(actorDtos, dto.ToActor(actor))
	}

	Ok(ctx, actorDtos)
}

func (h *Actor) GetByMovieID(ctx *gin.Context) {
	const op = "handler.Actor.GetByMovieID"
	log := h.log.With("op", op)

	movieID, role := dto.FromGetActorsByMovieRequest(ctx)

	actors, err := h.uc.GetByMovieID(ctx, movieID, role)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	var actorDtos []dto.Actor
	for _, actor := range actors {
		actorDtos = append(actorDtos, dto.ToActor(actor))
	}

	Ok(ctx, actorDtos)
}

func (h *Actor) Update(ctx *gin.Context) {
	const op = "handler.Actor.Update"
	log := h.log.With("op", op)

	id, update, err := dto.FromUpdateActorRequest(ctx)
	if err != nil {
		handleError(ctx, err)
		logError(log, err)

		return
	}

	actor, err := h.uc.UpdateByID(ctx, id, update)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	Ok(ctx, dto.ToActor(actor))
}

func (h *Actor) Delete(ctx *gin.Context) {
	const op = "handler.Actor.Delete"
	log := h.log.With("op", op)

	id := ctx.Param("id")

	actor, err := h.uc.DeleteByID(ctx, id)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	Ok(ctx, dto.ToActor(actor))
}
