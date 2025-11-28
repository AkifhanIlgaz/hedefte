package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type SessionRepository interface {
	InsertSession(models.Session) (models.Session, error)
	UpdateSession(id bson.ObjectID, userId string, fieldsToUpdate map[string]interface{}) error
	DeleteSession(id bson.ObjectID, userId string) error
	FindSession(id bson.ObjectID, userId string) (models.Session, error)
	FindAllSessionsOfDay(userId string, date time.Time) ([]models.Session, error)
}

type sessionRepository struct {
	collection *mongo.Collection
}

func NewSessionRepository(db *mongo.Database) SessionRepository {
	return sessionRepository{
		collection: db.Collection("sessions"),
	}
}

func (r sessionRepository) InsertSession(session models.Session) (models.Session, error) {
	res, err := r.collection.InsertOne(context.Background(), session)
	if err != nil {
		return models.Session{}, fmt.Errorf("failed to insert session: %w", err)
	}

	objectID, ok := res.InsertedID.(bson.ObjectID)
	if !ok {
		return models.Session{}, fmt.Errorf("failed to convert inserted ID to ObjectID")
	}

	session.Id = objectID
	return session, nil
}

func (r sessionRepository) UpdateSession(id bson.ObjectID, userId string, fieldsToUpdate map[string]any) error {
	filter := bson.M{"_id": id, "user_id": userId}
	update := bson.M{"$set": fieldsToUpdate}

	_, err := r.collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update session: %w", err)
	}

	return nil
}

func (r sessionRepository) DeleteSession(id bson.ObjectID, userId string) error {
	filter := bson.M{"_id": id, "user_id": userId}

	_, err := r.collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete session: %w", err)
	}
	return nil
}

// FindSession finds a session by its ID and user ID
func (r sessionRepository) FindSession(id bson.ObjectID, userId string) (models.Session, error) {
	filter := bson.M{"_id": id, "user_id": userId}

	var session models.Session
	err := r.collection.FindOne(context.Background(), filter).Decode(&session)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Session{}, fmt.Errorf("session not found")
		}
		return models.Session{}, fmt.Errorf("failed to find session: %w", err)
	}
	return session, nil
}

// FindAllSessions finds all sessions for a given user ID
func (r sessionRepository) FindAllSessionsOfDay(userId string, date time.Time) ([]models.Session, error) {
	filter := bson.M{"user_id": userId, "date": date}

	cursor, err := r.collection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find sessions: %w", err)
	}
	defer cursor.Close(context.Background())

	var sessions []models.Session
	for cursor.Next(context.Background()) {
		var session models.Session
		if err := cursor.Decode(&session); err != nil {
			return nil, fmt.Errorf("failed to decode session: %w", err)
		}
		sessions = append(sessions, session)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return sessions, nil
}
