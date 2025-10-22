package analysis

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Exam struct {
	ID             bson.ObjectID    `bson:"_id,omitempty"`
	UserId         string           `bson:"userId"`
	Name           string           `bson:"name"`
	Date           time.Time        `bson:"date"`
	TotalNet       float64          `bson:"totalNet"`
	LessonAnalysis []LessonAnalysis `bson:"lessonAnalysis"`
}

type LessonAnalysis struct {
	LessonName    string          `bson:"lessonName"`
	Correct       int             `bson:"correct"`
	Wrong         int             `bson:"wrong"`
	Empty         int             `bson:"empty"`
	Net           float64         `bson:"net"`
	Time          int             `bson:"time"` // in minutes
	TopicAnalysis []TopicAnalysis `bson:"topicAnalysis"`
}

type TopicAnalysis struct {
	TopicName string `bson:"topicName"`
	Mistakes  int    `bson:"mistakes"`
}
