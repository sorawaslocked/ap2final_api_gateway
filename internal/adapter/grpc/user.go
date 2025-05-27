package grpc

import (
	"context"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/adapter/grpc/dto"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
	svc "github.com/sorawaslocked/ap2final_protos_gen/service/user"
	"google.golang.org/grpc/metadata"
)

type User struct {
	client svc.UserServiceClient
}

func NewUser(client svc.UserServiceClient) *User {
	return &User{
		client: client,
	}
}

func (c *User) Register(ctx context.Context, user model.User) (model.User, error) {
	res, err := c.client.Register(ctx, dto.ToRegisterUserRequest(user))
	if err != nil {
		return model.User{}, wrapError(err)
	}

	registeredUser := dto.FromBaseUser(res.User)

	return registeredUser, nil
}

func (c *User) Login(ctx context.Context, user model.User) (model.Token, error) {
	res, err := c.client.Login(ctx, dto.ToLoginUserRequest(user))
	if err != nil {
		return model.Token{}, wrapError(err)
	}

	return model.Token{
		AccessToken:  res.Token.AccessToken,
		RefreshToken: res.Token.RefreshToken,
	}, nil
}

func (c *User) RefreshToken(ctx context.Context, refreshToken string) (model.Token, error) {
	res, err := c.client.RefreshToken(ctx, dto.ToRefreshTokenUserRequest(refreshToken))
	if err != nil {
		return model.Token{}, wrapError(err)
	}

	return model.Token{
		AccessToken:  res.Token.AccessToken,
		RefreshToken: res.Token.RefreshToken,
	}, nil
}

func (c *User) GetByID(
	ctx context.Context,
	token model.Token,
	id string,
) (model.User, error) {
	md := setAuthHeader(token)

	ctxC := metadata.NewOutgoingContext(ctx, md)

	res, err := c.client.Get(ctxC, dto.ToGetUserRequest(id))
	if err != nil {
		return model.User{}, wrapError(err)
	}

	return dto.FromBaseUser(res.User), nil
}

func (c *User) UpdateByID(
	ctx context.Context,
	token model.Token,
	id string,
	credentialsUpdate model.UserCredentialsUpdateData,
	update model.UserUpdateData,
) (model.User, error) {
	md := setAuthHeader(token)

	ctxC := metadata.NewOutgoingContext(ctx, md)

	res, err := c.client.Update(ctxC, dto.ToUpdateUserRequest(id, update, credentialsUpdate))
	if err != nil {
		return model.User{}, wrapError(err)
	}

	return dto.FromBaseUser(res.User), nil
}

func (c *User) DeleteByID(
	ctx context.Context,
	token model.Token,
	id string,
) (model.User, error) {
	md := setAuthHeader(token)

	ctxC := metadata.NewOutgoingContext(ctx, md)

	res, err := c.client.Delete(ctxC, dto.ToDeleteUserRequest(id))
	if err != nil {
		return model.User{}, wrapError(err)
	}

	return dto.FromBaseUser(res.User), nil
}
