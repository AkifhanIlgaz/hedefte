package models

import (
	"time"
)

type ExamAnalysisRequest struct {
	Date     time.Time `json:"date" validate:"required"`
	ExamType ExamType  `json:"examType"  validate:"required,oneof=TYT AYT"`
	TotalNet float64   `json:"totalNet" validate:"required,gte=0"`
	Name     string    `json:"name" validate:"required,min=1,max=100"`
	Subjects []Subject `json:"subjects" validate:"required,min=1,dive"`
}

type SubjectRequest struct {
	Id           uint                  `json:"id" validate:"required"`
	Correct      int                   `json:"correct" validate:"min=0"`
	Wrong        int                   `json:"wrong" validate:"min=0"`
	Empty        int                   `json:"empty" validate:"min=0"`
	Total        int                   `json:"total" validate:"min=1"`
	TopicMistake []TopicMistakeRequest `json:"topicMistakes" validate:"dive"`
}

type TopicMistakeRequest struct {
	Id       uint `json:"id" validate:"required"`
	Mistakes int  `json:"mistakes" validate:"min=0"`
}
