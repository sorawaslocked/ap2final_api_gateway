package usecase

import (
	"context"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
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
