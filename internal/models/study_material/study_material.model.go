package study_material

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type StudyMaterial struct {
	Id            bson.ObjectID `bson:"_id,omitempty"`
	UserId        string        `bson:"userId" json:"userId"`
	LessonId      bson.ObjectID `bson:"lessonId" json:"lessonId"`
	Name          string        `bson:"name" json:"name"`
	LastStudiedAt time.Time     `bson:"lastStudiedAt" json:"lastStudiedAt"`
}
