package analysis

import "go.mongodb.org/mongo-driver/v2/bson"

type Topic struct {
	Id       bson.ObjectID `bson:"_id,omitempty"`
	Name     string        `bson:"name"`
	LessonID bson.ObjectID `bson:"lessonId"`
}
