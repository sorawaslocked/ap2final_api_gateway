package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
	"time"
)

type Actor struct {
	ID          string             `json:"id"`
	FirstName   string             `json:"firstName"`
	LastName    string             `json:"lastName"`
	BirthDate   time.Time          `json:"birthDate"`
	DeathDate   *time.Time         `json:"deathDate,omitempty"`
	Nationality string             `json:"nationality"`
	Biography   string             `json:"biography"`
	ImageURL    string             `json:"imageUrl"`
	Filmography []FilmographyEntry `json:"filmography"`
	IsDeleted   bool               `json:"isDeleted"`
}

type FilmographyEntry struct {
	MovieID   string `json:"movieId"`
	Character string `json:"character"`
	Role      string `json:"role"`
}

type CreateActorRequest struct {
	FirstName   string             `json:"firstName"`
	LastName    string             `json:"lastName"`
	BirthDate   time.Time          `json:"birthDate"`
	DeathDate   *time.Time         `json:"deathDate,omitempty"`
	Nationality string             `json:"nationality"`
	Biography   string             `json:"biography"`
	ImageURL    string             `json:"imageUrl"`
	Filmography []FilmographyEntry `json:"filmography"`
}

type UpdateActorRequest struct {
	FirstName   *string             `json:"firstName"`
	LastName    *string             `json:"lastName"`
	BirthDate   *time.Time          `json:"birthDate"`
	DeathDate   *time.Time          `json:"deathDate"`
	Nationality *string             `json:"nationality"`
	Biography   *string             `json:"biography"`
	ImageURL    *string             `json:"imageUrl"`
	Filmography *[]FilmographyEntry `json:"filmography"`
	IsDeleted   *bool               `json:"isDeleted"`
}

type ActorFilterRequest struct {
	ID          *string `form:"id"`
	FirstName   *string `form:"firstName"`
	LastName    *string `form:"lastName"`
	AgeFrom     *uint8  `form:"ageFrom"`
	AgeTo       *uint8  `form:"ageTo"`
	Nationality *string `form:"nationality"`
	MovieID     *string `form:"movieId"`
	Role        *string `form:"role"`
	IsDeleted   *bool   `form:"isDeleted"`
}

func FromCreateActorRequest(ctx *gin.Context) (model.Actor, error) {
	var req CreateActorRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return model.Actor{}, ErrJSONBinding
	}

	filmography := make([]model.FilmographyEntry, len(req.Filmography))
	for i, entry := range req.Filmography {
		filmography[i] = model.FilmographyEntry{
			MovieID:   entry.MovieID,
			Character: entry.Character,
			Role:      entry.Role,
		}
	}

	return model.Actor{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		BirthDate:   req.BirthDate,
		DeathDate:   req.DeathDate,
		Nationality: req.Nationality,
		Biography:   req.Biography,
		ImageURL:    req.ImageURL,
		Filmography: filmography,
	}, nil
}

func FromUpdateActorRequest(ctx *gin.Context) (string, model.ActorUpdateData, error) {
	id := ctx.Param("id")

	var req UpdateActorRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return "", model.ActorUpdateData{}, ErrJSONBinding
	}

	var filmography *[]model.FilmographyEntry
	if req.Filmography != nil {
		entries := make([]model.FilmographyEntry, len(*req.Filmography))
		for i, entry := range *req.Filmography {
			entries[i] = model.FilmographyEntry{
				MovieID:   entry.MovieID,
				Character: entry.Character,
				Role:      entry.Role,
			}
		}
		filmography = &entries
	}

	return id, model.ActorUpdateData{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		BirthDate:   req.BirthDate,
		DeathDate:   req.DeathDate,
		Nationality: req.Nationality,
		Biography:   req.Biography,
		ImageURL:    req.ImageURL,
		Filmography: filmography,
		IsDeleted:   req.IsDeleted,
	}, nil
}

func FromActorFilterRequest(ctx *gin.Context) model.ActorFilter {
	var req ActorFilterRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		return model.ActorFilter{}
	}

	var ageRange *model.AgeRange
	if req.AgeFrom != nil && req.AgeTo != nil {
		ageRange = &model.AgeRange{
			AgeFrom: *req.AgeFrom,
			AgeTo:   *req.AgeTo,
		}
	}

	return model.ActorFilter{
		ID:          req.ID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		AgeRange:    ageRange,
		Nationality: req.Nationality,
		MovieID:     req.MovieID,
		Role:        req.Role,
		IsDeleted:   req.IsDeleted,
	}
}

func FromGetActorsByMovieRequest(ctx *gin.Context) (string, *string) {
	movieID := ctx.Param("movieId")
	role := ctx.Query("role")

	if role == "" {
		return movieID, nil
	}

	return movieID, &role
}

func ToActor(actor model.Actor) Actor {
	filmography := make([]FilmographyEntry, len(actor.Filmography))
	for i, entry := range actor.Filmography {
		filmography[i] = FilmographyEntry{
			MovieID:   entry.MovieID,
			Character: entry.Character,
			Role:      entry.Role,
		}
	}

	return Actor{
		ID:          actor.ID,
		FirstName:   actor.FirstName,
		LastName:    actor.LastName,
		BirthDate:   actor.BirthDate,
		DeathDate:   actor.DeathDate,
		Nationality: actor.Nationality,
		Biography:   actor.Biography,
		ImageURL:    actor.ImageURL,
		Filmography: filmography,
		IsDeleted:   actor.IsDeleted,
	}
}
