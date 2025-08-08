package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type ExamAnalysisDB struct {
	Id                  bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Uid                 string        `bson:"uid" json:"uid"`
	TotalNet            float64       `bson:"totalNet" json:"totalNet"`
	ExamAnalysisRequest `bson:",inline"`
}

type ExamAnalysisRequest struct {
	Date     time.Time `json:"date" bson:"date" validate:"required"`
	ExamType ExamType  `json:"examType" bson:"-" validate:"required,oneof=TYT AYT"`
	Name     string    `json:"name" bson:"name" validate:"required,min=1,max=100"`
	Subjects []Subject `json:"subjects" bson:"subjects" validate:"required,min=1,dive"`
}

type Subject struct {
	Name          string         `json:"name" bson:"name" validate:"required,min=1,max=50"`
	Correct       int            `json:"correct" bson:"correct" validate:"min=0"`
	Wrong         int            `json:"wrong" bson:"wrong" validate:"min=0"`
	Empty         int            `json:"empty" bson:"empty" validate:"min=0"`
	TopicMistakes []TopicMistake `json:"topicMistakes" bson:"topicMistakes" validate:"dive"`
}

type TopicMistake struct {
	Topic    string `json:"topic" bson:"topic" validate:"required,min=1,max=100"`
	Mistakes int    `json:"mistakes" bson:"mistakes" validate:"min=0"`
}
