package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
	"strings"
)

type User struct {
	ID           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	PhoneNumber  string `json:"phoneNumber"`
	PasswordHash string `json:"passwordHash"`
	Role         string `json:"role"`
	IsDeleted    bool   `json:"isDeleted"`
	IsActive     bool   `json:"isActive"`
}

type Token struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RegisterUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshTokenUserRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type UpdateUserRequest struct {
	FirstName       *string `json:"firstName"`
	LastName        *string `json:"lastName"`
	Email           *string `json:"email"`
	PhoneNumber     *string `json:"phoneNumber"`
	CurrentPassword *string `json:"currentPassword"`
	NewPassword     *string `json:"newPassword"`
	Role            *string `json:"role"`
	IsDeleted       *bool   `json:"isDeleted"`
	IsActive        *bool   `json:"isActive"`
}

func FromRegisterUserRequest(ctx *gin.Context) (model.User, error) {
	var req RegisterUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return model.User{}, ErrJSONBinding
	}

	return model.User{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}

func FromLoginUserRequest(ctx *gin.Context) (model.User, error) {
	var req LoginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return model.User{}, ErrJSONBinding
	}

	return model.User{
		Email:    req.Email,
		Password: req.Password,
	}, nil
}

func FromRefreshTokenUserRequest(ctx *gin.Context) (string, error) {
	var req RefreshTokenUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return "", ErrJSONBinding
	}

	return req.RefreshToken, nil
}

func FromGetUserRequest(ctx *gin.Context) (model.Token, string) {
	id := ctx.Param("id")
	token := getBearerToken(ctx)

	return model.Token{
		AccessToken: token,
	}, id
}

func FromUpdateUserRequest(ctx *gin.Context) (
	model.Token,
	string,
	model.UserCredentialsUpdateData,
	model.UserUpdateData,
	error,
) {
	id := ctx.Param("id")
	token := getBearerToken(ctx)

	var req UpdateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return model.Token{}, "", model.UserCredentialsUpdateData{}, model.UserUpdateData{}, ErrJSONBinding
	}

	credentialsUpdate := model.UserCredentialsUpdateData{}

	if req.CurrentPassword != nil && req.NewPassword != nil {
		credentialsUpdate.CurrentPassword = *req.CurrentPassword
		credentialsUpdate.NewPassword = *req.NewPassword
	}

	return model.Token{
			AccessToken: token,
		},
		id,
		credentialsUpdate,
		model.UserUpdateData{
			FirstName:   req.FirstName,
			LastName:    req.LastName,
			Email:       req.Email,
			PhoneNumber: req.PhoneNumber,
			Role:        req.Role,
			IsDeleted:   req.IsDeleted,
			IsActive:    req.IsActive,
		}, nil
}

func FromDeleteUserRequest(ctx *gin.Context) (model.Token, string) {
	id := ctx.Param("id")
	token := getBearerToken(ctx)

	return model.Token{
		AccessToken: token,
	}, id
}

func ToUser(user model.User) User {
	return User{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		PhoneNumber:  user.PhoneNumber,
		PasswordHash: user.PasswordHash,
		Role:         user.Role,
		IsDeleted:    user.IsDeleted,
		IsActive:     user.IsActive,
	}
}

func ToToken(token model.Token) Token {
	return Token{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}
}

func getBearerToken(ctx *gin.Context) string {
	header := ctx.GetHeader("Authorization")

	return strings.TrimPrefix(header, "Bearer ")
}
