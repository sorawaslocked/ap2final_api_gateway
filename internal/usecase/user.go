package usecase

import (
	"context"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
)

type User struct {
	presenter UserPresenter
}

func NewUser(presenter UserPresenter) *User {
	return &User{presenter: presenter}
}

func (u *User) Register(ctx context.Context, user model.User) (model.User, error) {
	return u.presenter.Register(ctx, user)
}

func (u *User) Login(ctx context.Context, user model.User) (model.Token, error) {
	return u.presenter.Login(ctx, user)
}

func (u *User) RefreshToken(ctx context.Context, refreshToken string) (model.Token, error) {
	return u.presenter.RefreshToken(ctx, refreshToken)
}

func (u *User) GetByID(ctx context.Context, token model.Token, id string) (model.User, error) {
	return u.presenter.GetByID(ctx, token, id)
}

func (u *User) UpdateByID(
	ctx context.Context,
	token model.Token,
	id string,
	credentialsUpdate model.UserCredentialsUpdateData,
	update model.UserUpdateData,
) (model.User, error) {
	return u.presenter.UpdateByID(ctx, token, id, credentialsUpdate, update)
}

func (u *User) DeleteByID(ctx context.Context, token model.Token, id string) (model.User, error) {
	return u.presenter.DeleteByID(ctx, token, id)
}
