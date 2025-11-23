package analysis

import (
	"math"
	"time"
)

type AYTAnalysis struct {
	Date      time.Time      `json:"date" bson:"date"`
	Name      string         `json:"name" bson:"name"`
	TotalNet  float64        `json:"totalNet" bson:"total_net"`
	Edebiyat  LessonAnalysis `json:"Edebiyat" bson:"Edebiyat,omitempty"`
	Tarih     LessonAnalysis `json:"Tarih" bson:"Tarih,omitempty"`
	Coğrafya  LessonAnalysis `json:"Coğrafya" bson:"Coğrafya,omitempty"`
	Matematik LessonAnalysis `json:"Matematik" bson:"Matematik,omitempty"`
	Fizik     LessonAnalysis `json:"Fizik" bson:"Fizik,omitempty"`
	Kimya     LessonAnalysis `json:"Kimya" bson:"Kimya,omitempty"`
	Biyoloji  LessonAnalysis `json:"Biyoloji" bson:"Biyoloji,omitempty"`
}

func (a AYTAnalysis) ApplyAnalysisToGeneralChartData(chartData *GeneralChartData) {
	exam := GeneralChartExam{
		TotalNet: a.TotalNet,
		Date:     a.Date,
		Name:     a.Name,
	}

	chartData.MaxNet = math.Max(chartData.MaxNet, exam.TotalNet)
	chartData.AverageNet = calculateAverage(chartData.AverageNet, a.TotalNet, float64(chartData.ExamCount))

	chartData.Exams = append(chartData.Exams, exam)

	ApplyLessonAnalysisToGeneralChartData("Edebiyat", a.Edebiyat, chartData)
	ApplyLessonAnalysisToGeneralChartData("Tarih", a.Tarih, chartData)
	ApplyLessonAnalysisToGeneralChartData("Coğrafya", a.Coğrafya, chartData)
	ApplyLessonAnalysisToGeneralChartData("Matematik", a.Matematik, chartData)
	ApplyLessonAnalysisToGeneralChartData("Fizik", a.Fizik, chartData)
	ApplyLessonAnalysisToGeneralChartData("Kimya", a.Kimya, chartData)
	ApplyLessonAnalysisToGeneralChartData("Biyoloji", a.Biyoloji, chartData)

	chartData.ExamCount++
}
