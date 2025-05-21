package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequest(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}

func InternalServerError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
}

func NotFound(ctx *gin.Context) {
	ctx.JSON(http.StatusNotFound, gin.H{"error": "model with this id not found"})
}

func Ok(ctx *gin.Context, v any) {
	ctx.JSON(http.StatusOK, v)
}

func Created(ctx *gin.Context, v any) {
	ctx.JSON(http.StatusCreated, v)
}
