package analysis

import (
	"math"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type AYTAnalysis struct {
	Id        bson.ObjectID  `bson:"_id,omitempty"`
	UserId    string         `json:"userId" bson:"user_id"`
	Date      time.Time      `json:"date" bson:"date"`
	Name      string         `json:"name" bson:"name"`
	TotalNet  float64        `json:"totalNet" bson:"total_net"`
	Edebiyat  LessonAnalysis `json:"Edebiyat" bson:"edebiyat,omitempty"`
	Tarih     LessonAnalysis `json:"Tarih" bson:"tarih,omitempty"`
	Coğrafya  LessonAnalysis `json:"Coğrafya" bson:"cografya,omitempty"`
	Matematik LessonAnalysis `json:"Matematik" bson:"matematik,omitempty"`
	Fizik     LessonAnalysis `json:"Fizik" bson:"fizik,omitempty"`
	Kimya     LessonAnalysis `json:"Kimya" bson:"kimya,omitempty"`
	Biyoloji  LessonAnalysis `json:"Biyoloji" bson:"biyoloji,omitempty"`
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
