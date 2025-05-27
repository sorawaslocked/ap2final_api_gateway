package grpc

import (
	"context"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/adapter/grpc/dto"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
	svc "github.com/sorawaslocked/ap2final_protos_gen/service/session"
	"time"
)

type Session struct {
	client svc.SessionServiceClient
}

func NewSession(client svc.SessionServiceClient) *Session {
	return &Session{
		client: client,
	}
}

func (c *Session) Create(ctx context.Context, session model.Session) (model.Session, error) {
	res, err := c.client.Create(ctx, dto.ToCreateSessionRequest(session))
	if err != nil {
		return model.Session{}, wrapError(err)
	}

	createdSession := dto.FromBaseSession(res.Session)

	return createdSession, nil
}

func (c *Session) GetByID(ctx context.Context, id string) (model.Session, error) {
	res, err := c.client.Get(ctx, dto.ToGetSessionRequest(id))
	if err != nil {
		return model.Session{}, wrapError(err)
	}

	session := dto.FromBaseSession(res.Session)

	return session, nil
}

func (c *Session) GetAll(ctx context.Context) ([]model.Session, error) {
	res, err := c.client.GetAll(ctx, &svc.GetAllRequest{})
	if err != nil {
		return []model.Session{}, wrapError(err)
	}

	var sessions []model.Session

	for _, session := range res.Sessions {
		sessions = append(sessions, dto.FromBaseSession(session))
	}

	return sessions, nil
}

func (c *Session) GetAllWithFilter(ctx context.Context, filter model.SessionFilter) ([]model.Session, error) {
	res, err := c.client.GetAllWithFilter(ctx, dto.ToSessionFilterRequest(filter))
	if err != nil {
		return []model.Session{}, wrapError(err)
	}

	var sessions []model.Session

	for _, session := range res.Sessions {
		sessions = append(sessions, dto.FromBaseSession(session))
	}

	return sessions, nil
}

func (c *Session) GetByMovieID(ctx context.Context, movieID string) ([]model.Session, error) {
	res, err := c.client.GetByMovieID(ctx, dto.ToGetSessionsByMovieIDRequest(movieID))
	if err != nil {
		return []model.Session{}, wrapError(err)
	}

	var sessions []model.Session

	for _, session := range res.Sessions {
		sessions = append(sessions, dto.FromBaseSession(session))
	}

	return sessions, nil
}

func (c *Session) GetByCinemaHallID(ctx context.Context, cinemaHallID string) ([]model.Session, error) {
	res, err := c.client.GetByCinemaHallID(ctx, dto.ToGetSessionsByCinemaHallIDRequest(cinemaHallID))
	if err != nil {
		return []model.Session{}, wrapError(err)
	}

	var sessions []model.Session

	for _, session := range res.Sessions {
		sessions = append(sessions, dto.FromBaseSession(session))
	}

	return sessions, nil
}

func (c *Session) GetByTimeRange(ctx context.Context, startTime, endTime time.Time) ([]model.Session, error) {
	res, err := c.client.GetByTimeRange(ctx, dto.ToGetSessionsByTimeRangeRequest(startTime, endTime))
	if err != nil {
		return []model.Session{}, wrapError(err)
	}

	var sessions []model.Session

	for _, session := range res.Sessions {
		sessions = append(sessions, dto.FromBaseSession(session))
	}

	return sessions, nil
}

func (c *Session) GetAvailableSessions(ctx context.Context) ([]model.Session, error) {
	res, err := c.client.GetAvailableSessions(ctx, &svc.GetAvailableSessionsRequest{})
	if err != nil {
		return []model.Session{}, wrapError(err)
	}

	var sessions []model.Session

	for _, session := range res.Sessions {
		sessions = append(sessions, dto.FromBaseSession(session))
	}

	return sessions, nil
}

func (c *Session) UpdateByID(ctx context.Context, id string, update model.SessionUpdateData) (model.Session, error) {
	res, err := c.client.Update(ctx, dto.ToUpdateSessionRequest(id, update))
	if err != nil {
		return model.Session{}, wrapError(err)
	}

	updatedSession := dto.FromBaseSession(res.Session)

	return updatedSession, nil
}

func (c *Session) UpdateSeatAvailability(ctx context.Context, id string, bookedSeatsCount int) (model.Session, error) {
	res, err := c.client.UpdateSeatAvailability(ctx, dto.ToUpdateSeatAvailabilityRequest(id, bookedSeatsCount))
	if err != nil {
		return model.Session{}, wrapError(err)
	}

	updatedSession := dto.FromBaseSession(res.Session)

	return updatedSession, nil
}

func (c *Session) ChangeStatus(ctx context.Context, id string, status model.SessionStatus) (model.Session, error) {
	res, err := c.client.ChangeStatus(ctx, dto.ToChangeStatusRequest(id, status))
	if err != nil {
		return model.Session{}, wrapError(err)
	}

	updatedSession := dto.FromBaseSession(res.Session)

	return updatedSession, nil
}

func (c *Session) DeleteByID(ctx context.Context, id string) (model.Session, error) {
	res, err := c.client.Delete(ctx, dto.ToDeleteSessionRequest(id))
	if err != nil {
		return model.Session{}, wrapError(err)
	}

	deletedSession := dto.FromBaseSession(res.Session)

	return deletedSession, nil
}
