package dto

import (
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
	"github.com/sorawaslocked/ap2final_protos_gen/base"
	svc "github.com/sorawaslocked/ap2final_protos_gen/service/user"
)

func FromBaseUser(res *base.User) model.User {
	return model.User{
		ID:           res.ID,
		FirstName:    res.FirstName,
		LastName:     res.LastName,
		Email:        res.Email,
		PhoneNumber:  res.PhoneNumber,
		PasswordHash: res.PasswordHash,
		Role:         res.Role,
		CreatedAt:    res.CreatedAt.AsTime(),
		UpdatedAt:    res.UpdatedAt.AsTime(),
		IsDeleted:    res.IsDeleted,
		IsActive:     res.IsActive,
	}
}

func ToRegisterUserRequest(user model.User) *svc.RegisterRequest {
	return &svc.RegisterRequest{
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToLoginUserRequest(user model.User) *svc.LoginRequest {
	return &svc.LoginRequest{
		Email:    user.Email,
		Password: user.Password,
	}
}

func ToRefreshTokenUserRequest(refreshToken string) *svc.RefreshTokenRequest {
	return &svc.RefreshTokenRequest{
		RefreshToken: refreshToken,
	}
}

func ToGetUserRequest(id string) *svc.GetRequest {
	return &svc.GetRequest{
		ID: id,
	}
}

func ToUpdateUserRequest(
	id string,
	update model.UserUpdateData,
	credentialsUpdate model.UserCredentialsUpdateData,
) *svc.UpdateRequest {
	req := &svc.UpdateRequest{
		ID:          id,
		FirstName:   update.FirstName,
		LastName:    update.LastName,
		Email:       update.Email,
		PhoneNumber: update.PhoneNumber,
		Role:        update.Role,
		IsDeleted:   update.IsDeleted,
		IsActive:    update.IsActive,
	}

	if credentialsUpdate.CurrentPassword != "" && credentialsUpdate.NewPassword != "" {
		req.CurrentPassword = &credentialsUpdate.CurrentPassword
		req.NewPassword = &credentialsUpdate.NewPassword
	}

	return req
}

func ToDeleteUserRequest(id string) *svc.DeleteRequest {
	return &svc.DeleteRequest{
		ID: id,
	}
}
