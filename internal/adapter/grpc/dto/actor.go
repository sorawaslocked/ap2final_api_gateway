package dto

import (
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
	"github.com/sorawaslocked/ap2final_protos_gen/base"
	svc "github.com/sorawaslocked/ap2final_protos_gen/service/actor"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func FromBaseActor(res *base.Actor) model.Actor {
	var deathDate *time.Time
	if res.DeathDate != nil {
		death := res.DeathDate.AsTime()
		deathDate = &death
	}

	filmography := make([]model.FilmographyEntry, len(res.Filmography))
	for i, entry := range res.Filmography {
		filmography[i] = model.FilmographyEntry{
			MovieID:   entry.MovieID,
			Character: entry.Character,
			Role:      entry.Role,
		}
	}

	return model.Actor{
		ID:          res.ID,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		BirthDate:   res.BirthDate.AsTime(),
		DeathDate:   deathDate,
		Nationality: res.Nationality,
		Biography:   res.Biography,
		ImageURL:    res.ImageURL,
		Filmography: filmography,
		CreatedAt:   res.CreatedAt.AsTime(),
		UpdatedAt:   res.UpdatedAt.AsTime(),
		IsDeleted:   res.IsDeleted,
	}
}

func ToCreateActorRequest(actor model.Actor) *svc.CreateRequest {
	var deathDate *timestamppb.Timestamp
	if actor.DeathDate != nil {
		deathDate = timestamppb.New(*actor.DeathDate)
	}

	filmography := make([]*base.FilmographyEntry, len(actor.Filmography))
	for i, entry := range actor.Filmography {
		filmography[i] = &base.FilmographyEntry{
			MovieID:   entry.MovieID,
			Character: entry.Character,
			Role:      entry.Role,
		}
	}

	return &svc.CreateRequest{
		FirstName:   actor.FirstName,
		LastName:    actor.LastName,
		BirthDate:   timestamppb.New(actor.BirthDate),
		DeathDate:   deathDate,
		Nationality: actor.Nationality,
		Biography:   actor.Biography,
		ImageURL:    actor.ImageURL,
		Filmography: filmography,
	}
}

func ToGetActorRequest(id string) *svc.GetRequest {
	return &svc.GetRequest{
		ID: id,
	}
}

func ToUpdateActorRequest(id string, update model.ActorUpdateData) *svc.UpdateRequest {
	var birthDate, deathDate *timestamppb.Timestamp

	if update.BirthDate != nil {
		birthDate = timestamppb.New(*update.BirthDate)
	}

	if update.DeathDate != nil {
		deathDate = timestamppb.New(*update.DeathDate)
	}

	var filmography []*base.FilmographyEntry
	if update.Filmography != nil {
		filmography = make([]*base.FilmographyEntry, len(*update.Filmography))
		for i, entry := range *update.Filmography {
			filmography[i] = &base.FilmographyEntry{
				MovieID:   entry.MovieID,
				Character: entry.Character,
				Role:      entry.Role,
			}
		}
	}

	return &svc.UpdateRequest{
		ID:          id,
		FirstName:   update.FirstName,
		LastName:    update.LastName,
		BirthDate:   birthDate,
		DeathDate:   deathDate,
		Nationality: update.Nationality,
		Biography:   update.Biography,
		ImageURL:    update.ImageURL,
		Filmography: filmography,
		IsDeleted:   update.IsDeleted,
	}
}

func ToDeleteActorRequest(id string) *svc.DeleteRequest {
	return &svc.DeleteRequest{
		ID: id,
	}
}

func ToGetActorsByMovieIDRequest(movieID string, role *string) *svc.GetByMovieIDRequest {
	return &svc.GetByMovieIDRequest{
		MovieID: movieID,
		Role:    role,
	}
}

func ToActorFilterRequest(filter model.ActorFilter) *svc.GetAllWithFilterRequest {
	if filter == (model.ActorFilter{}) {
		return &svc.GetAllWithFilterRequest{}
	}

	var ageRange *svc.AgeRange
	if filter.AgeRange != nil {
		ageRange = &svc.AgeRange{
			AgeFrom: uint32(filter.AgeRange.AgeFrom),
			AgeTo:   uint32(filter.AgeRange.AgeTo),
		}
	}

	return &svc.GetAllWithFilterRequest{
		Filter: &svc.ActorFilter{
			ID:          filter.ID,
			FirstName:   filter.FirstName,
			LastName:    filter.LastName,
			AgeRange:    ageRange,
			Nationality: filter.Nationality,
			MovieID:     filter.MovieID,
			Role:        filter.Role,
			IsDeleted:   filter.IsDeleted,
		},
	}
}
