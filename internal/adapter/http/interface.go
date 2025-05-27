package http

import "github.com/sorawaslocked/ap2final_api_gateway/internal/adapter/http/handler"

type MovieUseCase interface {
	handler.MovieUseCase
}

type UserUseCase interface {
	handler.UserUseCase
}

type ActorUseCase interface {
	handler.ActorUseCase
}

type SessionUseCase interface {
	handler.SessionUseCase
}
