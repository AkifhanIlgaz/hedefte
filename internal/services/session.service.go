package services

import (
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"github.com/AkifhanIlgaz/hedefte/internal/repositories"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

type SessionService struct {
	repo   repositories.SessionRepository
	logger *zap.Logger
}

func NewSessionService(repo repositories.SessionRepository, logger *zap.Logger) SessionService {
	return SessionService{
		repo:   repo,
		logger: logger,
	}
}

func (s SessionService) AddSession(req models.AddSessionRequest) (models.Session, error) {
	session := req.ToSession()
	insertedSession, err := s.repo.InsertSession(session)
	if err != nil {
		s.logger.Error("failed to add session", zap.Error(err))
		return models.Session{}, fmt.Errorf(`failed to add session: %w`, err)
	}
	return insertedSession, nil
}

func (s SessionService) UpdateSession(req models.UpdateSessionRequest) (models.Session, error) {
	session := req.ToSession()
	updatedSession, err := s.repo.UpdateSession(session)
	if err != nil {
		s.logger.Error("failed to update session", zap.Error(err))
		return models.Session{}, fmt.Errorf(`failed to update session: %w`, err)
	}
	return updatedSession, nil
}

func (s SessionService) DeleteSession(id string, userId string) error {
	sessionId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid session ID: %w", err)
	}

	err = s.repo.DeleteSession(sessionId, userId)
	if err != nil {
		s.logger.Error("failed to delete session", zap.Error(err))
		return fmt.Errorf(`failed to delete session: %w`, err)
	}
	return nil
}

func (s SessionService) GetSessionsOfDay(userId string, day time.Time) ([]models.Session, error) {
	sessions, err := s.repo.FindAllSessionsOfDay(userId, day)
	if err != nil {
		s.logger.Error("failed to retrieve sessions", zap.Error(err))
		return nil, fmt.Errorf("failed to retrieve sessions: %w", err)
	}

	return sessions, nil
}

func (s SessionService) ToggleCompletion(id string, userId string, isCompleted bool) (models.Session, error) {
	sessionId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		return models.Session{}, fmt.Errorf("invalid session ID: %w", err)
	}
	session := models.Session{
		Id:          sessionId,
		UserId:      userId,
		IsCompleted: isCompleted,
	}
	updatedSession, err := s.repo.UpdateSession(session)
	if err != nil {
		s.logger.Error("failed to update session", zap.Error(err))
		return models.Session{}, fmt.Errorf(`failed to update session: %w`, err)
	}
	return updatedSession, nil
}
