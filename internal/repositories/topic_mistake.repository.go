package repositories

import (
	"context"

	"github.com/AkifhanIlgaz/hedefte/internal/constants"
	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TopicMistakeRepository interface {
	Insert(topicMistake models.TopicMistake) error
	InsertBulk(topicMistakes []models.TopicMistake) error
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
