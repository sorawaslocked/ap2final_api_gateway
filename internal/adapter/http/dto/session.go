package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
	"time"
)

type Session struct {
	ID             string    `json:"id"`
	MovieID        string    `json:"movieId"`
	CinemaHallID   string    `json:"cinemaHallId"`
	StartTime      time.Time `json:"startTime"`
	EndTime        time.Time `json:"endTime"`
	Price          float64   `json:"price"`
	AvailableSeats int       `json:"availableSeats"`
	TotalSeats     int       `json:"totalSeats"`
	Status         string    `json:"status"`
	IsDeleted      bool      `json:"isDeleted"`
}

type CreateSessionRequest struct {
	MovieID      string    `json:"movieId"`
	CinemaHallID string    `json:"cinemaHallId"`
	StartTime    time.Time `json:"startTime"`
	EndTime      time.Time `json:"endTime"`
	Price        float64   `json:"price"`
	TotalSeats   int       `json:"totalSeats"`
}

type UpdateSessionRequest struct {
	MovieID        *string    `json:"movieId"`
	CinemaHallID   *string    `json:"cinemaHallId"`
	StartTime      *time.Time `json:"startTime"`
	EndTime        *time.Time `json:"endTime"`
	Price          *float64   `json:"price"`
	AvailableSeats *int       `json:"availableSeats"`
	TotalSeats     *int       `json:"totalSeats"`
	Status         *string    `json:"status"`
	IsDeleted      *bool      `json:"isDeleted"`
}

type SessionFilterRequest struct {
	ID           *string    `form:"id"`
	MovieID      *string    `form:"movieId"`
	CinemaHallID *string    `form:"cinemaHallId"`
	StartTimeGTE *time.Time `form:"startTimeGte" time_format:"2006-01-02T15:04:05Z07:00"`
	StartTimeLTE *time.Time `form:"startTimeLte" time_format:"2006-01-02T15:04:05Z07:00"`
	Status       *string    `form:"status"`
	IsDeleted    *bool      `form:"isDeleted"`
	HasAvailable *bool      `form:"hasAvailable"`
}

type UpdateSeatAvailabilityRequest struct {
	BookedSeatsCount int `json:"bookedSeatsCount"`
}

type ChangeStatusRequest struct {
	Status string `json:"status"`
}

func FromCreateSessionRequest(ctx *gin.Context) (model.Session, error) {
	var req CreateSessionRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return model.Session{}, ErrJSONBinding
	}

	return model.Session{
		MovieID:      req.MovieID,
		CinemaHallID: req.CinemaHallID,
		StartTime:    req.StartTime,
		EndTime:      req.EndTime,
		Price:        req.Price,
		TotalSeats:   req.TotalSeats,
	}, nil
}

func FromUpdateSessionRequest(ctx *gin.Context) (string, model.SessionUpdateData, error) {
	id := ctx.Param("id")

	var req UpdateSessionRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return "", model.SessionUpdateData{}, ErrJSONBinding
	}

	var status *model.SessionStatus
	if req.Status != nil {
		stat := model.SessionStatus(*req.Status)
		status = &stat
	}

	return id, model.SessionUpdateData{
		MovieID:        req.MovieID,
		CinemaHallID:   req.CinemaHallID,
		StartTime:      req.StartTime,
		EndTime:        req.EndTime,
		Price:          req.Price,
		AvailableSeats: req.AvailableSeats,
		TotalSeats:     req.TotalSeats,
		Status:         status,
		IsDeleted:      req.IsDeleted,
	}, nil
}

func FromSessionFilterRequest(ctx *gin.Context) model.SessionFilter {
	var req SessionFilterRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		return model.SessionFilter{}
	}

	var status *model.SessionStatus
	if req.Status != nil {
		stat := model.SessionStatus(*req.Status)
		status = &stat
	}

	return model.SessionFilter{
		ID:           req.ID,
		MovieID:      req.MovieID,
		CinemaHallID: req.CinemaHallID,
		StartTimeGTE: req.StartTimeGTE,
		StartTimeLTE: req.StartTimeLTE,
		Status:       status,
		IsDeleted:    req.IsDeleted,
		HasAvailable: req.HasAvailable,
	}
}

func FromUpdateSeatAvailabilityRequest(ctx *gin.Context) (string, int, error) {
	id := ctx.Param("id")

	var req UpdateSeatAvailabilityRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return "", 0, ErrJSONBinding
	}

	return id, req.BookedSeatsCount, nil
}

func FromChangeStatusRequest(ctx *gin.Context) (string, model.SessionStatus, error) {
	id := ctx.Param("id")

	var req ChangeStatusRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		return "", "", ErrJSONBinding
	}

	return id, model.SessionStatus(req.Status), nil
}

func FromGetSessionsByTimeRangeRequest(ctx *gin.Context) (time.Time, time.Time, error) {
	startTimeStr := ctx.Query("startTime")
	endTimeStr := ctx.Query("endTime")

	startTime, err := time.Parse(time.RFC3339, startTimeStr)
	if err != nil {
		return time.Time{}, time.Time{}, ErrJSONBinding
	}

	endTime, err := time.Parse(time.RFC3339, endTimeStr)
	if err != nil {
		return time.Time{}, time.Time{}, ErrJSONBinding
	}

	return startTime, endTime, nil
}

func ToSession(session model.Session) Session {
	return Session{
		ID:             session.ID,
		MovieID:        session.MovieID,
		CinemaHallID:   session.CinemaHallID,
		StartTime:      session.StartTime,
		EndTime:        session.EndTime,
		Price:          session.Price,
		AvailableSeats: session.AvailableSeats,
		TotalSeats:     session.TotalSeats,
		Status:         string(session.Status),
		IsDeleted:      session.IsDeleted,
	}
}
