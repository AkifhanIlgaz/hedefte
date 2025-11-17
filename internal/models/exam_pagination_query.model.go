package models

import (
	"time"
)

type ExamPaginationQuery struct {
	UserId       string `json:"-"`
	Page         int    `form:"page"`
	RowsPerPage  int    `form:"rowsPerPage"`
	TimeInterval int    `form:"timeInterval"`
}

func (p ExamPaginationQuery) GetStart() time.Time {
	switch p.TimeInterval {
	case 1:
		return time.Now().AddDate(0, -1, 0)
	case 3:
		return time.Now().AddDate(0, -3, 0)
	case 6:
		return time.Now().AddDate(0, -6, 0)
	default:
		return time.Now().AddDate(-2, 0, 0)
	}
}

func (p ExamPaginationQuery) GetEnd() time.Time {
	return time.Now()
}
