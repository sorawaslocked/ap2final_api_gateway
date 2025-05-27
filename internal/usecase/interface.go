package usecase

import (
	"context"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
	"time"
)

type MoviePresenter interface {
	Create(ctx context.Context, movie model.Movie) (model.Movie, error)
	GetByID(ctx context.Context, id string) (model.Movie, error)
	GetAll(ctx context.Context) ([]model.Movie, error)
	UpdateByID(ctx context.Context, id string, update model.MovieUpdateData) (model.Movie, error)
	DeleteByID(ctx context.Context, id string) (model.Movie, error)
}

type UserPresenter interface {
	Register(ctx context.Context, user model.User) (model.User, error)
	Login(ctx context.Context, user model.User) (model.Token, error)
	RefreshToken(ctx context.Context, refreshToken string) (model.Token, error)
	GetByID(ctx context.Context, token model.Token, id string) (model.User, error)
	UpdateByID(ctx context.Context, token model.Token, id string, credentialsUpdate model.UserCredentialsUpdateData, update model.UserUpdateData) (model.User, error)
	DeleteByID(ctx context.Context, token model.Token, id string) (model.User, error)
}

type ActorPresenter interface {
	Create(ctx context.Context, actor model.Actor) (model.Actor, error)
	GetByID(ctx context.Context, id string) (model.Actor, error)
	GetAll(ctx context.Context) ([]model.Actor, error)
	GetAllWithFilter(ctx context.Context, filter model.ActorFilter) ([]model.Actor, error)
	GetByMovieID(ctx context.Context, movieID string, role *string) ([]model.Actor, error)
	UpdateByID(ctx context.Context, id string, update model.ActorUpdateData) (model.Actor, error)
	DeleteByID(ctx context.Context, id string) (model.Actor, error)
}

type SessionPresenter interface {
	Create(ctx context.Context, session model.Session) (model.Session, error)
	GetByID(ctx context.Context, id string) (model.Session, error)
	GetAll(ctx context.Context) ([]model.Session, error)
	GetAllWithFilter(ctx context.Context, filter model.SessionFilter) ([]model.Session, error)
	GetByMovieID(ctx context.Context, movieID string) ([]model.Session, error)
	GetByCinemaHallID(ctx context.Context, cinemaHallID string) ([]model.Session, error)
	GetByTimeRange(ctx context.Context, startTime, endTime time.Time) ([]model.Session, error)
	GetAvailableSessions(ctx context.Context) ([]model.Session, error)
	UpdateByID(ctx context.Context, id string, update model.SessionUpdateData) (model.Session, error)
	UpdateSeatAvailability(ctx context.Context, id string, bookedSeatsCount int) (model.Session, error)
	ChangeStatus(ctx context.Context, id string, status model.SessionStatus) (model.Session, error)
	DeleteByID(ctx context.Context, id string) (model.Session, error)
}
