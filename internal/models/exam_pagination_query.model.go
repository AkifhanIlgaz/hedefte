package models

import "time"

type ExamPaginationQuery struct {
	UserId string    `json:"-"`
	Page   int       `json:"page" validate:"required,min=1"`
	Limit  int       `json:"limit" validate:"required,min=1,max=100"`
	Start  time.Time `json:"start" validate:"required"`
	End    time.Time `json:"end" validate:"required"`
}
