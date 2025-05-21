package http

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/adapter/http/handler"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/config"
	"log/slog"
	"net/http"
)

type API struct {
	server       *gin.Engine
	cfg          config.HTTPServer
	log          *slog.Logger
	addr         string
	movieHandler *handler.Movie
}

func New(cfg config.HTTPServer, log *slog.Logger, movieUseCase MovieUseCase) *API {
	gin.SetMode(cfg.GinMode)

	server := gin.New()

	// Middleware
	server.Use(gin.Recovery())

	// Injecting presenters
	movieHandler := handler.NewMovie(log, movieUseCase)

	api := &API{
		server:       server,
		cfg:          cfg,
		log:          log,
		addr:         cfg.Address,
		movieHandler: movieHandler,
	}

	api.setupRoutes()

	return api
}

func (a *API) setupRoutes() {
	v1 := a.server.Group("/api/v1")
	{
		movies := v1.Group("/movies")
		{
			movies.POST("/", a.movieHandler.Create)
			movies.GET("/:id", a.movieHandler.Get)
			movies.GET("/", a.movieHandler.GetAll)
			movies.PATCH("/:id", a.movieHandler.Update)
			movies.DELETE("/:id", a.movieHandler.Delete)
		}
	}
}

func (a *API) MustRun() {
	a.log.Info("starting http server", slog.String("addr", a.addr))

	go func() {
		err := a.server.Run(a.addr)

		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
}

func (a *API) Stop(ctx context.Context) {
	a.log.Info("stopping http server")

	<-ctx.Done()
}
