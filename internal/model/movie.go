package model

import "time"

type Movie struct {
	ID               string
	AgeRating        string
	PrimaryTitle     string
	OriginalTitle    string
	ReleaseYear      uint16
	RuntimeInMinutes uint16
	Genres           []string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	IsDeleted        bool
}

type MovieUpdateData struct {
	AgeRating        *string
	PrimaryTitle     *string
	OriginalTitle    *string
	ReleaseYear      *uint16
	RuntimeInMinutes *uint16
	Genres           []string
	UpdatedAt        time.Time
	IsDeleted        *bool
}
