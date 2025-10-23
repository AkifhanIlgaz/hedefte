package analysis

import "go.mongodb.org/mongo-driver/v2/bson"

type Topic struct {
	Id       bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string        `bson:"name" json:"name"`
	LessonID bson.ObjectID `bson:"lessonId" json:"lessonId"`
}
