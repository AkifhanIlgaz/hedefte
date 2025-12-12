package models

import (
	"time"
)

type ExamResponse struct {
	Id      string           `json:"id"`
	Date    time.Time        `json:"date"`
	Name    string           `json:"name"`
	Result  float64          `json:"result"`
	Lessons []LessonResponse `json:"lessons"`
}

type LessonResponse struct {
	Name    string  `json:"name"`
	Correct int     `json:"correct"`
	Wrong   int     `json:"wrong"`
	Empty   int     `json:"empty"`
	Time    int     `json:"time"`
	Result  float64 `json:"result"`
}
