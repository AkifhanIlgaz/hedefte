package models

import (
	"time"
)

type ExamAnalysisRequest struct {
	Date     time.Time `json:"tarih" validate:"required"`
	ExamType ExamType  `json:"examType" validate:"required,oneof=TYT AYT"`
	Name     string    `json:"name" validate:"required,min=1,max=100"`
	Notes    string    `json:"notes" validate:"max=500"`
	Subjects []Subject `json:"subjects" validate:"required,min=1,dive"`
}

type Subject struct {
	Name          string         `json:"name" validate:"required,min=1,max=50"`
	Correct       int            `json:"correct" validate:"min=0"`
	Wrong         int            `json:"wrong" validate:"min=0"`
	Empty         int            `json:"empty" validate:"min=0"`
	TopicMistakes []TopicMistake `json:"topicMistakes" validate:"dive"`
}

type TopicMistake struct {
	Topic    string `json:"topic" validate:"required,min=1,max=100"`
	Mistakes int    `json:"mistakes" validate:"min=0"`
}
