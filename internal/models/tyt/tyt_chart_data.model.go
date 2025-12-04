package tyt_models

import (
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
)

type GeneralChartData struct {
	ExamCount  int                           `json:"examCount"`
	MaxNet     float64                       `json:"maxNet"`
	AverageNet float64                       `json:"averageNet"`
	Exams      []GeneralChartExam            `json:"exams"`
	Türkçe     models.LessonGeneralChartData `json:"Türkçe"`
	Tarih      models.LessonGeneralChartData `json:"Tarih"`
	Coğrafya   models.LessonGeneralChartData `json:"Coğrafya"`
	Felsefe    models.LessonGeneralChartData `json:"Felsefe"`
	DinKültürü models.LessonGeneralChartData `json:"Din Kültürü"`
	Matematik  models.LessonGeneralChartData `json:"Matematik"`
	Fizik      models.LessonGeneralChartData `json:"Fizik"`
	Kimya      models.LessonGeneralChartData `json:"Kimya"`
	Biyoloji   models.LessonGeneralChartData `json:"Biyoloji"`
}

func NewGeneralChartData() GeneralChartData {
	return GeneralChartData{
		ExamCount:  0,
		MaxNet:     0,
		AverageNet: 0,
		Exams:      []GeneralChartExam{},
	}
}

type GeneralChartExam struct {
	Date     time.Time `json:"date"`
	Name     string    `json:"name"`
	TotalNet float64   `json:"totalNet,omitempty"`
}
