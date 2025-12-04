package tyt_models

import (
	"math"
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Exam struct {
	Id         bson.ObjectID         `bson:"_id,omitempty"`
	UserId     string                `json:"userId" bson:"user_id"`
	Date       time.Time             `json:"date" bson:"date"`
	Name       string                `json:"name" bson:"name"`
	TotalNet   float64               `json:"totalNet" bson:"total_net"`
	Türkçe     models.LessonAnalysis `json:"Türkçe" bson:"turkce"`
	Tarih      models.LessonAnalysis `json:"Tarih" bson:"tarih"`
	Coğrafya   models.LessonAnalysis `json:"Coğrafya" bson:"cografya"`
	Felsefe    models.LessonAnalysis `json:"Felsefe" bson:"felsefe"`
	DinKültürü models.LessonAnalysis `json:"Din Kültürü" bson:"din_kulturu"`
	Matematik  models.LessonAnalysis `json:"Matematik" bson:"matematik"`
	Fizik      models.LessonAnalysis `json:"Fizik" bson:"fizik"`
	Kimya      models.LessonAnalysis `json:"Kimya" bson:"kimya"`
	Biyoloji   models.LessonAnalysis `json:"Biyoloji" bson:"biyoloji"`
}

func (e Exam) ApplyExamToGeneralChartData(chartData *GeneralChartData) {
	exam := GeneralChartExam{
		TotalNet: e.TotalNet,
		Date:     e.Date,
		Name:     e.Name,
	}

	chartData.MaxNet = math.Max(chartData.MaxNet, exam.TotalNet)
	chartData.AverageNet = calculateAverage(chartData.AverageNet, e.TotalNet, float64(chartData.ExamCount))
	chartData.Exams = append(chartData.Exams, exam)

	ApplyLessonAnalysisToGeneralChartData(e.Türkçe, &chartData.Türkçe, chartData.ExamCount)
	ApplyLessonAnalysisToGeneralChartData(e.Tarih, &chartData.Tarih, chartData.ExamCount)
	ApplyLessonAnalysisToGeneralChartData(e.Coğrafya, &chartData.Coğrafya, chartData.ExamCount)
	ApplyLessonAnalysisToGeneralChartData(e.Felsefe, &chartData.Felsefe, chartData.ExamCount)
	ApplyLessonAnalysisToGeneralChartData(e.DinKültürü, &chartData.DinKültürü, chartData.ExamCount)
	ApplyLessonAnalysisToGeneralChartData(e.Matematik, &chartData.Matematik, chartData.ExamCount)
	ApplyLessonAnalysisToGeneralChartData(e.Fizik, &chartData.Fizik, chartData.ExamCount)
	ApplyLessonAnalysisToGeneralChartData(e.Kimya, &chartData.Kimya, chartData.ExamCount)
	ApplyLessonAnalysisToGeneralChartData(e.Biyoloji, &chartData.Biyoloji, chartData.ExamCount)

	chartData.ExamCount++
}

func ApplyLessonAnalysisToGeneralChartData(lessonAnalysis models.LessonAnalysis, lessonData *models.LessonGeneralChartData, examCount int) {
	lessonData.MaxNet = math.Max(lessonData.MaxNet, lessonAnalysis.Net)
	lessonData.AverageTime = calculateAverage(lessonData.AverageTime, lessonAnalysis.Time, examCount)
	lessonData.AverageNet = calculateAverage(lessonData.AverageNet, lessonAnalysis.Net, float64(examCount))

}

func calculateAverage[T int | float64](oldAverage, newValue, itemCount T) T {
	return (oldAverage*(itemCount) + newValue) / (itemCount + 1)
}
