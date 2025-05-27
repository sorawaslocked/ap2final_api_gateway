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
	server         *gin.Engine
	cfg            config.HTTPServer
	log            *slog.Logger
	addr           string
	movieHandler   *handler.Movie
	userHandler    *handler.User
	actorHandler   *handler.Actor
	sessionHandler *handler.Session
}

func New(
	cfg config.HTTPServer,
	log *slog.Logger,
	movieUseCase MovieUseCase,
	userUseCase UserUseCase,
	actorUseCase ActorUseCase,
	sessionUseCase SessionUseCase,
) *API {
	gin.SetMode(cfg.GinMode)

	server := gin.New()

	// Middleware
	server.Use(gin.Recovery())

	// Injecting presenters
	movieHandler := handler.NewMovie(log, movieUseCase)
	userHandler := handler.NewUser(log, userUseCase)
	actorHandler := handler.NewActor(log, actorUseCase)
	sessionHandler := handler.NewSession(log, sessionUseCase)

	api := &API{
		server:         server,
		cfg:            cfg,
		log:            log,
		addr:           cfg.Address,
		movieHandler:   movieHandler,
		userHandler:    userHandler,
		actorHandler:   actorHandler,
		sessionHandler: sessionHandler,
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
		users := v1.Group("/users")
		{
			users.POST("/register", a.userHandler.Register)
			users.POST("/login", a.userHandler.Login)
			users.POST("/refreshToken", a.userHandler.RefreshToken)
			users.GET("/:id", a.userHandler.Get)
			users.PATCH("/:id", a.userHandler.Update)
			users.DELETE("/:id", a.userHandler.Delete)
		}
		actors := v1.Group("/actors")
		{
			actors.POST("/", a.actorHandler.Create)
			actors.GET("/:id", a.actorHandler.Get)
			actors.GET("/", a.actorHandler.GetAll)
			actors.GET("/filter", a.actorHandler.GetAllWithFilter)
			actors.GET("/movie/:movieId", a.actorHandler.GetByMovieID)
			actors.PATCH("/:id", a.actorHandler.Update)
			actors.DELETE("/:id", a.actorHandler.Delete)
		}
		sessions := v1.Group("/sessions")
		{
			sessions.POST("/", a.sessionHandler.Create)
			sessions.GET("/:id", a.sessionHandler.Get)
			sessions.GET("/", a.sessionHandler.GetAll)
			sessions.GET("/filter", a.sessionHandler.GetAllWithFilter)
			sessions.GET("/movie/:movieId", a.sessionHandler.GetByMovieID)
			sessions.GET("/cinema-hall/:cinemaHallId", a.sessionHandler.GetByCinemaHallID)
			sessions.GET("/time-range", a.sessionHandler.GetByTimeRange)
			sessions.GET("/available", a.sessionHandler.GetAvailableSessions)
			sessions.PATCH("/:id", a.sessionHandler.Update)
			sessions.PATCH("/:id/seats", a.sessionHandler.UpdateSeatAvailability)
			sessions.PATCH("/:id/status", a.sessionHandler.ChangeStatus)
			sessions.DELETE("/:id", a.sessionHandler.Delete)
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
