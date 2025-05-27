package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/adapter/http/dto"
	"log/slog"
)

type Session struct {
	log *slog.Logger
	uc  SessionUseCase
}

func NewSession(log *slog.Logger, uc SessionUseCase) *Session {
	return &Session{
		log: log,
		uc:  uc,
	}
}

func (h *Session) Create(ctx *gin.Context) {
	const op = "handler.Session.Create"
	log := h.log.With("op", op)

	session, err := dto.FromCreateSessionRequest(ctx)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	newSession, err := h.uc.Create(ctx, session)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	Created(ctx, dto.ToSession(newSession))
}

func (h *Session) Get(ctx *gin.Context) {
	const op = "handler.Session.Get"
	log := h.log.With("op", op)

	id := ctx.Param("id")

	session, err := h.uc.GetByID(ctx, id)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	Ok(ctx, dto.ToSession(session))
}

func (h *Session) GetAll(ctx *gin.Context) {
	const op = "handler.Session.GetAll"
	log := h.log.With("op", op)

	sessions, err := h.uc.GetAll(ctx)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	var sessionDtos []dto.Session
	for _, session := range sessions {
		sessionDtos = append(sessionDtos, dto.ToSession(session))
	}

	Ok(ctx, sessionDtos)
}

func (h *Session) GetAllWithFilter(ctx *gin.Context) {
	const op = "handler.Session.GetAllWithFilter"
	log := h.log.With("op", op)

	filter := dto.FromSessionFilterRequest(ctx)

	sessions, err := h.uc.GetAllWithFilter(ctx, filter)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	var sessionDtos []dto.Session
	for _, session := range sessions {
		sessionDtos = append(sessionDtos, dto.ToSession(session))
	}

	Ok(ctx, sessionDtos)
}

func (h *Session) GetByMovieID(ctx *gin.Context) {
	const op = "handler.Session.GetByMovieID"
	log := h.log.With("op", op)

	movieID := ctx.Param("movieId")

	sessions, err := h.uc.GetByMovieID(ctx, movieID)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	var sessionDtos []dto.Session
	for _, session := range sessions {
		sessionDtos = append(sessionDtos, dto.ToSession(session))
	}

	Ok(ctx, sessionDtos)
}

func (h *Session) GetByCinemaHallID(ctx *gin.Context) {
	const op = "handler.Session.GetByCinemaHallID"
	log := h.log.With("op", op)

	cinemaHallID := ctx.Param("cinemaHallId")

	sessions, err := h.uc.GetByCinemaHallID(ctx, cinemaHallID)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	var sessionDtos []dto.Session
	for _, session := range sessions {
		sessionDtos = append(sessionDtos, dto.ToSession(session))
	}

	Ok(ctx, sessionDtos)
}

func (h *Session) GetByTimeRange(ctx *gin.Context) {
	const op = "handler.Session.GetByTimeRange"
	log := h.log.With("op", op)

	startTime, endTime, err := dto.FromGetSessionsByTimeRangeRequest(ctx)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	sessions, err := h.uc.GetByTimeRange(ctx, startTime, endTime)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	var sessionDtos []dto.Session
	for _, session := range sessions {
		sessionDtos = append(sessionDtos, dto.ToSession(session))
	}

	Ok(ctx, sessionDtos)
}

func (h *Session) GetAvailableSessions(ctx *gin.Context) {
	const op = "handler.Session.GetAvailableSessions"
	log := h.log.With("op", op)

	sessions, err := h.uc.GetAvailableSessions(ctx)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	var sessionDtos []dto.Session
	for _, session := range sessions {
		sessionDtos = append(sessionDtos, dto.ToSession(session))
	}

	Ok(ctx, sessionDtos)
}

func (h *Session) Update(ctx *gin.Context) {
	const op = "handler.Session.Update"
	log := h.log.With("op", op)

	id, update, err := dto.FromUpdateSessionRequest(ctx)
	if err != nil {
		handleError(ctx, err)
		logError(log, err)

		return
	}

	session, err := h.uc.UpdateByID(ctx, id, update)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	Ok(ctx, dto.ToSession(session))
}

func (h *Session) UpdateSeatAvailability(ctx *gin.Context) {
	const op = "handler.Session.UpdateSeatAvailability"
	log := h.log.With("op", op)

	id, bookedSeatsCount, err := dto.FromUpdateSeatAvailabilityRequest(ctx)
	if err != nil {
		handleError(ctx, err)
		logError(log, err)

		return
	}

	session, err := h.uc.UpdateSeatAvailability(ctx, id, bookedSeatsCount)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	Ok(ctx, dto.ToSession(session))
}

func (h *Session) ChangeStatus(ctx *gin.Context) {
	const op = "handler.Session.ChangeStatus"
	log := h.log.With("op", op)

	id, status, err := dto.FromChangeStatusRequest(ctx)
	if err != nil {
		handleError(ctx, err)
		logError(log, err)

		return
	}

	session, err := h.uc.ChangeStatus(ctx, id, status)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	Ok(ctx, dto.ToSession(session))
}

func (h *Session) Delete(ctx *gin.Context) {
	const op = "handler.Session.Delete"
	log := h.log.With("op", op)

	id := ctx.Param("id")

	session, err := h.uc.DeleteByID(ctx, id)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	Ok(ctx, dto.ToSession(session))
}
