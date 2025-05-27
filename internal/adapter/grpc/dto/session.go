package dto

import (
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
	"github.com/sorawaslocked/ap2final_protos_gen/base"
	svc "github.com/sorawaslocked/ap2final_protos_gen/service/session"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

func FromBaseSession(res *base.Session) model.Session {
	return model.Session{
		ID:             res.ID,
		MovieID:        res.MovieID,
		CinemaHallID:   res.CinemaHallID,
		StartTime:      res.StartTime.AsTime(),
		EndTime:        res.EndTime.AsTime(),
		Price:          res.Price,
		AvailableSeats: int(res.AvailableSeats),
		TotalSeats:     int(res.TotalSeats),
		Status:         model.SessionStatus(res.Status),
		CreatedAt:      res.CreatedAt.AsTime(),
		UpdatedAt:      res.UpdatedAt.AsTime(),
		IsDeleted:      res.IsDeleted,
	}
}

func ToCreateSessionRequest(session model.Session) *svc.CreateRequest {
	return &svc.CreateRequest{
		MovieID:      session.MovieID,
		CinemaHallID: session.CinemaHallID,
		StartTime:    timestamppb.New(session.StartTime),
		EndTime:      timestamppb.New(session.EndTime),
		Price:        session.Price,
		TotalSeats:   int32(session.TotalSeats),
	}
}

func ToGetSessionRequest(id string) *svc.GetRequest {
	return &svc.GetRequest{
		ID: id,
	}
}

func ToUpdateSessionRequest(id string, update model.SessionUpdateData) *svc.UpdateRequest {
	var startTime, endTime *timestamppb.Timestamp
	var price *float64
	var availableSeats, totalSeats *int32
	var status *string

	if update.StartTime != nil {
		startTime = timestamppb.New(*update.StartTime)
	}

	if update.EndTime != nil {
		endTime = timestamppb.New(*update.EndTime)
	}

	if update.Price != nil {
		price = update.Price
	}

	if update.AvailableSeats != nil {
		seats := int32(*update.AvailableSeats)
		availableSeats = &seats
	}

	if update.TotalSeats != nil {
		seats := int32(*update.TotalSeats)
		totalSeats = &seats
	}

	if update.Status != nil {
		stat := string(*update.Status)
		status = &stat
	}

	return &svc.UpdateRequest{
		ID:             id,
		MovieID:        update.MovieID,
		CinemaHallID:   update.CinemaHallID,
		StartTime:      startTime,
		EndTime:        endTime,
		Price:          price,
		AvailableSeats: availableSeats,
		TotalSeats:     totalSeats,
		Status:         status,
		IsDeleted:      update.IsDeleted,
	}
}

func ToDeleteSessionRequest(id string) *svc.DeleteRequest {
	return &svc.DeleteRequest{
		ID: id,
	}
}

func ToGetSessionsByMovieIDRequest(movieID string) *svc.GetByMovieIDRequest {
	return &svc.GetByMovieIDRequest{
		MovieID: movieID,
	}
}

func ToGetSessionsByCinemaHallIDRequest(cinemaHallID string) *svc.GetByCinemaHallIDRequest {
	return &svc.GetByCinemaHallIDRequest{
		CinemaHallID: cinemaHallID,
	}
}

func ToGetSessionsByTimeRangeRequest(startTime, endTime time.Time) *svc.GetByTimeRangeRequest {
	return &svc.GetByTimeRangeRequest{
		StartTime: timestamppb.New(startTime),
		EndTime:   timestamppb.New(endTime),
	}
}

func ToSessionFilterRequest(filter model.SessionFilter) *svc.GetAllWithFilterRequest {
	if filter == (model.SessionFilter{}) {
		return &svc.GetAllWithFilterRequest{}
	}

	var startTimeGTE, startTimeLTE *timestamppb.Timestamp
	var status *string

	if filter.StartTimeGTE != nil {
		startTimeGTE = timestamppb.New(*filter.StartTimeGTE)
	}

	if filter.StartTimeLTE != nil {
		startTimeLTE = timestamppb.New(*filter.StartTimeLTE)
	}

	if filter.Status != nil {
		stat := string(*filter.Status)
		status = &stat
	}

	return &svc.GetAllWithFilterRequest{
		Filter: &svc.SessionFilter{
			ID:           filter.ID,
			MovieID:      filter.MovieID,
			CinemaHallID: filter.CinemaHallID,
			StartTimeGTE: startTimeGTE,
			StartTimeLTE: startTimeLTE,
			Status:       status,
			IsDeleted:    filter.IsDeleted,
			HasAvailable: filter.HasAvailable,
		},
	}
}

func ToUpdateSeatAvailabilityRequest(id string, bookedSeatsCount int) *svc.UpdateSeatAvailabilityRequest {
	return &svc.UpdateSeatAvailabilityRequest{
		ID:               id,
		BookedSeatsCount: int32(bookedSeatsCount),
	}
}

func ToChangeStatusRequest(id string, status model.SessionStatus) *svc.ChangeStatusRequest {
	return &svc.ChangeStatusRequest{
		ID:     id,
		Status: string(status),
	}
}
