package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type AddExamRequest struct {
	Date           time.Time               `json:"date" validate:"required"`
	Name           string                  `json:"name" validate:"required,min=1,max=100"`
	TotalNet       float64                 `json:"totalNet" validate:"min=1"`
	LessonAnalysis []LessonAnalysisRequest `json:"lessonAnalysis" validate:"required,min=1,dive"`
}

type LessonAnalysisRequest struct {
	LessonId      bson.ObjectID          `json:"lessonId" validate:"required"`
	Correct       int                    `json:"correct" validate:"min=0"`
	Wrong         int                    `json:"wrong" validate:"min=0"`
	Empty         int                    `json:"empty" validate:"min=0"`
	TotalNet      float64                `json:"totalNet" validate:"min=1"`
	Time          int                    `json:"time" validate:"min=0"` // in minutes
	TopicAnalysis []TopicAnalysisRequest `json:"topicAnalysis" validate:"dive"`
}

type TopicAnalysisRequest struct {
	TopicId  bson.ObjectID `json:"topicId" validate:"required"`
	Mistakes int           `json:"mistakes" validate:"min=0"`
}
