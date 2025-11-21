package models

import "time"

type GeneralChartData struct {
	ExamCount  int                `json:"examCount"`
	MaxNet     float64            `json:"maxNet"`
	MinNet     float64            `json:"minNet"`
	AverageNet float64            `json:"averageNet"`
	Exams      []GeneralChartExam `json:"exams"`
}

type GeneralChartExam struct {
	Date     time.Time `json:"date"`
	Name     string    `json:"name"`
	TotalNet float64   `json:"totalNet"`
}
