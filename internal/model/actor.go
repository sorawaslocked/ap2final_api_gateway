package model

import "time"

// Actor represents a person who acts in movies or works as crew
type Actor struct {
	ID          string
	FirstName   string
	LastName    string
	BirthDate   time.Time
	DeathDate   *time.Time // Pointer since it could be nil for living actors
	Nationality string
	Biography   string
	ImageURL    string             // URL to actor's image
	Filmography []FilmographyEntry // Movies they participated in with their roles
	CreatedAt   time.Time
	UpdatedAt   time.Time
	IsDeleted   bool
}

// FilmographyEntry represents an actor's participation in a specific movie
type FilmographyEntry struct {
	MovieID   string
	Character string // The character they played (empty for crew)
	Role      string // "Actor" or crew role (Director, Producer, etc.)
}

// ActorFilter for searching actors based on criteria
type ActorFilter struct {
	ID          *string
	FirstName   *string
	LastName    *string
	AgeRange    *AgeRange
	Nationality *string
	MovieID     *string // Filter actors by a specific movie they were in
	Role        *string // Filter by role type
	IsDeleted   *bool
}

// AgeRange is inclusive
type AgeRange struct {
	AgeFrom uint8
	AgeTo   uint8
}

// ActorUpdateData for partial updates to actor records
type ActorUpdateData struct {
	FirstName   *string
	LastName    *string
	BirthDate   *time.Time
	DeathDate   *time.Time
	Nationality *string
	Biography   *string
	ImageURL    *string
	Filmography *[]FilmographyEntry // Using pointer for optional update
	IsDeleted   *bool
}
