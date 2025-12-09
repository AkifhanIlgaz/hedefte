package analysis

import "time"

type GeneralChartData struct {
	ExamCount  int                               `json:"examCount"`
	MaxNet     float64                           `json:"maxNet"`
	AverageNet float64                           `json:"averageNet"`
	Exams      []GeneralChartExam                `json:"exams"`
	Lessons    map[string]LessonGeneralChartData `json:"lessons"`
}

func NewGeneralChartData() GeneralChartData {
	return GeneralChartData{
		ExamCount:  0,
		MaxNet:     0,
		AverageNet: 0,
		Exams:      []GeneralChartExam{},
		Lessons:    map[string]LessonGeneralChartData{},
	}
}

type GeneralChartExam struct {
	Date     time.Time `json:"date"`
	Name     string    `json:"name"`
	TotalNet float64   `json:"totalNet,omitempty"`
}

type LessonGeneralChartData struct {
	AverageNet  float64 `json:"averageNet"`
	AverageTime int     `json:"averageTime"`
	MaxNet      float64 `json:"maxNet"`
}

type LessonSpecificChartData struct {
	ExamCount   int                `json:"examCount"`
	AverageNet  float64            `json:"averageNet"`
	AverageTime int                `json:"averageTime"`
	MaxNet      float64            `json:"maxNet"`
	Exams       []GeneralChartExam `json:"exams"`
}

func NewLessonSpecificChartData() LessonSpecificChartData {
	return LessonSpecificChartData{
		Exams: []GeneralChartExam{},
	}
}
