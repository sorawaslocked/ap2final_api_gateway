package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/adapter/http/dto"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
	"github.com/sorawaslocked/ap2final_base/pkg/logger"
	"log/slog"
)

func logError(log *slog.Logger, err error) {
	switch {
	case errors.Is(err, model.ErrNotFound):
		return
	default:
		log.Error("http error", logger.Err(err))
	}
}

func handleError(ctx *gin.Context, err error) {
	switch {
	case errors.Is(err, dto.ErrJSONBinding):
		BadRequest(ctx, err)
	case errors.Is(err, model.ErrNotFound):
		NotFound(ctx)
	default:
		InternalServerError(ctx, model.ErrDefault)
	}
}
