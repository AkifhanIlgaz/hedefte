package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/constants"
	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TopicMistakeRepository interface {
	Insert(topicMistake models.TopicMistake) error
	InsertBulk(topicMistakes []models.TopicMistake) error
	FindAllByLesson(userId string, examType string, lesson string, start, end time.Time) ([]models.TopicMistake, error)
	FindAllByExamType(userId string, examType string) ([]models.TopicMistake, error)
	FindAllByUser(userId string) ([]models.TopicMistake, error)
	FindAllByExamId(userId string, examId bson.ObjectID) ([]models.TopicMistake, error)

	Delete(id bson.ObjectID, userId string) error
}

type topicMistakeRepository struct {
	topicMistakesCollection *mongo.Collection
}

func NewTopicMistakeRepository(db *mongo.Database) TopicMistakeRepository {
	return topicMistakeRepository{
		topicMistakesCollection: db.Collection(constants.TopicMistakesCollection),
	}
}

func (r topicMistakeRepository) Insert(topicMistake models.TopicMistake) error {
	_, err := r.topicMistakesCollection.InsertOne(context.Background(), topicMistake)
	return err
}

func (r topicMistakeRepository) InsertBulk(topicMistakes []models.TopicMistake) error {
	_, err := r.topicMistakesCollection.InsertMany(context.Background(), topicMistakes)
	return err
}

func (r topicMistakeRepository) Delete(id bson.ObjectID, userId string) error {
	_, err := r.topicMistakesCollection.DeleteOne(context.Background(), bson.M{"_id": id, "user_id": userId})
	return err
}

func (r topicMistakeRepository) FindAllByLesson(userId string, examType string, lesson string, start, end time.Time) ([]models.TopicMistake, error) {
	lesson = strings.ToUpper(lesson[:1]) + lesson[1:]

	filter := bson.M{"user_id": userId, "exam_type": examType, "lesson": lesson}

	fmt.Println(filter)

	cursor, err := r.topicMistakesCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var topicMistakes []models.TopicMistake
	if err := cursor.All(context.Background(), &topicMistakes); err != nil {
		return nil, err
	}
	fmt.Println(topicMistakes)

	return topicMistakes, nil
}

func (r topicMistakeRepository) FindAllByExamType(userId string, exam string) ([]models.TopicMistake, error) {
	cursor, err := r.topicMistakesCollection.Find(context.Background(), bson.M{"user_id": userId, "exam": exam})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var topicMistakes []models.TopicMistake
	if err := cursor.All(context.Background(), &topicMistakes); err != nil {
		return nil, err
	}

	return topicMistakes, nil
}

func (r topicMistakeRepository) FindAllByUser(userId string) ([]models.TopicMistake, error) {
	cursor, err := r.topicMistakesCollection.Find(context.Background(), bson.M{"user_id": userId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var topicMistakes []models.TopicMistake
	if err := cursor.All(context.Background(), &topicMistakes); err != nil {
		return nil, err
	}

	return topicMistakes, nil
}

func (r topicMistakeRepository) FindAllByExamId(userId string, examId bson.ObjectID) ([]models.TopicMistake, error) {
	cursor, err := r.topicMistakesCollection.Find(context.Background(), bson.M{"user_id": userId, "exam_id": examId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var topicMistakes []models.TopicMistake
	if err := cursor.All(context.Background(), &topicMistakes); err != nil {
		return nil, err
	}

	return topicMistakes, nil
}
