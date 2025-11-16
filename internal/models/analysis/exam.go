package analysis

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Exam struct {
	ID             bson.ObjectID    `bson:"_id,omitempty" json:"id"`
	UserId         string           `bson:"userId" json:"userId"`
	Name           string           `bson:"name" json:"name"`
	Date           time.Time        `bson:"date" json:"date"`
	TotalNet       float64          `bson:"totalNet" json:"totalNet"`
	LessonAnalysis []LessonAnalysis `bson:"lessonAnalysis" json:"lessonAnalysis"`
}

type LessonAnalysis struct {
	LessonName    string          `bson:"lessonName" json:"lessonName"`
	Correct       int             `bson:"correct" json:"correct"`
	Wrong         int             `bson:"wrong" json:"wrong"`
	Empty         int             `bson:"empty" json:"empty"`
	TotalNet      float64         `bson:"totalNet" json:"totalNet"`
	Time          int             `bson:"time" json:"time"` // in minutes
	TopicAnalysis []TopicAnalysis `bson:"topicAnalysis" json:"topicAnalysis"`
}

type TopicAnalysis struct {
	TopicName string `bson:"topicName" json:"topicName"`
	Mistakes  int    `bson:"mistakes" json:"mistakes"`
}
