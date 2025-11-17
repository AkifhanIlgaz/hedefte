package models

import "time"

type ExamPaginationQuery struct {
	UserId      string    `json:"-"`
	Page        int       `form:"page"`
	RowsPerPage int       `form:"rowsPerPage"`
	Start       time.Time `form:"start"`
	End         time.Time `form:"end"`
}
