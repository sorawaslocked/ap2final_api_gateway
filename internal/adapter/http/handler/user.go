package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/adapter/http/dto"
	"log/slog"
)

type User struct {
	log *slog.Logger
	uc  UserUseCase
}

func NewUser(log *slog.Logger, uc UserUseCase) *User {
	return &User{
		log: log,
		uc:  uc,
	}
}

func (h *User) Register(ctx *gin.Context) {
	const op = "handler.User.Register"
	log := h.log.With(slog.String("op", op))

	user, err := dto.FromRegisterUserRequest(ctx)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	registeredUser, err := h.uc.Register(ctx, user)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	Created(ctx, dto.ToUser(registeredUser))
}

func (h *User) Login(ctx *gin.Context) {
	const op = "handler.User.Login"

	log := h.log.With(slog.String("op", op))

	user, err := dto.FromLoginUserRequest(ctx)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	token, err := h.uc.Login(ctx, user)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	Ok(ctx, dto.ToToken(token))
}

func (h *User) RefreshToken(ctx *gin.Context) {
	const op = "handler.User.RefreshToken"

	log := h.log.With(slog.String("op", op))

	refreshToken, err := dto.FromRefreshTokenUserRequest(ctx)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	token, err := h.uc.RefreshToken(ctx, refreshToken)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	Ok(ctx, dto.ToToken(token))
}

func (h *User) Get(ctx *gin.Context) {
	const op = "handler.User.Get"

	log := h.log.With(slog.String("op", op))

	token, id := dto.FromGetUserRequest(ctx)

	user, err := h.uc.GetByID(ctx, token, id)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	Ok(ctx, dto.ToUser(user))
}

func (h *User) Update(ctx *gin.Context) {
	const op = "handler.User.Update"

	log := h.log.With(slog.String("op", op))

	token, id, credentialsUpdate, update, err := dto.FromUpdateUserRequest(ctx)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	updatedUser, err := h.uc.UpdateByID(ctx, token, id, credentialsUpdate, update)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	Ok(ctx, dto.ToUser(updatedUser))
}

func (h *User) Delete(ctx *gin.Context) {
	const op = "handler.User.Delete"

	log := h.log.With(slog.String("op", op))

	token, id := dto.FromDeleteUserRequest(ctx)

	deletedUser, err := h.uc.DeleteByID(ctx, token, id)
	if err != nil {
		logError(log, err)
		handleError(ctx, err)

		return
	}

	Ok(ctx, dto.ToUser(deletedUser))
}
