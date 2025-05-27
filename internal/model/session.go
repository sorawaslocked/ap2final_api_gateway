package model

import "time"

// SessionStatus represents the possible states of a session
type SessionStatus string

const (
	StatusScheduled SessionStatus = "scheduled"
	StatusActive    SessionStatus = "active"
	StatusFinished  SessionStatus = "finished"
	StatusCancelled SessionStatus = "cancelled"
)

// Session represents a movie session in a cinema hall
type Session struct {
	ID             string
	MovieID        string
	CinemaHallID   string
	StartTime      time.Time
	EndTime        time.Time // Calculated based on movie duration
	Price          float64
	AvailableSeats int
	TotalSeats     int
	Status         SessionStatus
	CreatedAt      time.Time
	UpdatedAt      time.Time
	IsDeleted      bool
}

// SessionFilter for searching sessions based on criteria
type SessionFilter struct {
	ID           *string
	MovieID      *string
	CinemaHallID *string
	StartTimeGTE *time.Time // Greater than or equal to
	StartTimeLTE *time.Time // Less than or equal to
	Status       *SessionStatus
	IsDeleted    *bool
	HasAvailable *bool // Filter for sessions with available seats
}

// SessionUpdateData for partial updates to session records
type SessionUpdateData struct {
	MovieID        *string
	CinemaHallID   *string
	StartTime      *time.Time
	EndTime        *time.Time
	Price          *float64
	AvailableSeats *int
	TotalSeats     *int
	Status         *SessionStatus
	IsDeleted      *bool
}
