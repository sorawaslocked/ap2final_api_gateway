package usecase

import (
	"context"
	"github.com/sorawaslocked/ap2final_api_gateway/internal/model"
	"time"
)

type SessionUseCase struct {
	presenter SessionPresenter
}

func NewSession(presenter SessionPresenter) *SessionUseCase {
	return &SessionUseCase{
		presenter: presenter,
	}
}

func (uc *SessionUseCase) Create(ctx context.Context, session model.Session) (model.Session, error) {
	return uc.presenter.Create(ctx, session)
}

func (uc *SessionUseCase) GetByID(ctx context.Context, id string) (model.Session, error) {
	return uc.presenter.GetByID(ctx, id)
}

func (uc *SessionUseCase) GetAll(ctx context.Context) ([]model.Session, error) {
	return uc.presenter.GetAll(ctx)
}

func (uc *SessionUseCase) GetAllWithFilter(ctx context.Context, filter model.SessionFilter) ([]model.Session, error) {
	return uc.presenter.GetAllWithFilter(ctx, filter)
}

func (uc *SessionUseCase) GetByMovieID(ctx context.Context, movieID string) ([]model.Session, error) {
	return uc.presenter.GetByMovieID(ctx, movieID)
}

func (uc *SessionUseCase) GetByCinemaHallID(ctx context.Context, cinemaHallID string) ([]model.Session, error) {
	return uc.presenter.GetByCinemaHallID(ctx, cinemaHallID)
}

func (uc *SessionUseCase) GetByTimeRange(ctx context.Context, startTime, endTime time.Time) ([]model.Session, error) {
	return uc.presenter.GetByTimeRange(ctx, startTime, endTime)
}

func (uc *SessionUseCase) GetAvailableSessions(ctx context.Context) ([]model.Session, error) {
	return uc.presenter.GetAvailableSessions(ctx)
}

func (uc *SessionUseCase) UpdateByID(ctx context.Context, id string, update model.SessionUpdateData) (model.Session, error) {
	return uc.presenter.UpdateByID(ctx, id, update)
}

func (uc *SessionUseCase) UpdateSeatAvailability(ctx context.Context, id string, bookedSeatsCount int) (model.Session, error) {
	return uc.presenter.UpdateSeatAvailability(ctx, id, bookedSeatsCount)
}

func (uc *SessionUseCase) ChangeStatus(ctx context.Context, id string, status model.SessionStatus) (model.Session, error) {
	return uc.presenter.ChangeStatus(ctx, id, status)
}

func (uc *SessionUseCase) DeleteByID(ctx context.Context, id string) (model.Session, error) {
	return uc.presenter.DeleteByID(ctx, id)
}
