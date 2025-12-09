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

func (e Exam) ApplyExamToGeneralChartData(chartData *models.GeneralChartData) {
	exam := models.GeneralChartExam{
		TotalNet: e.TotalNet,
		Date:     e.Date,
		Name:     e.Name,
	}

	chartData.MaxNet = math.Max(chartData.MaxNet, exam.TotalNet)
	chartData.AverageNet = calculateAverage(chartData.AverageNet, e.TotalNet, float64(chartData.ExamCount))
	chartData.Exams = append(chartData.Exams, exam)

	ApplyLessonAnalysisToGeneralChartData("Türkçe", e.Türkçe, chartData)
	ApplyLessonAnalysisToGeneralChartData("Tarih", e.Tarih, chartData)
	ApplyLessonAnalysisToGeneralChartData("Coğrafya", e.Coğrafya, chartData)
	ApplyLessonAnalysisToGeneralChartData("Felsefe", e.Felsefe, chartData)
	ApplyLessonAnalysisToGeneralChartData("Din Kültürü", e.DinKültürü, chartData)
	ApplyLessonAnalysisToGeneralChartData("Matematik", e.Matematik, chartData)
	ApplyLessonAnalysisToGeneralChartData("Fizik", e.Fizik, chartData)
	ApplyLessonAnalysisToGeneralChartData("Kimya", e.Kimya, chartData)
	ApplyLessonAnalysisToGeneralChartData("Biyoloji", e.Biyoloji, chartData)

	chartData.ExamCount++
}

func ApplyLessonAnalysisToGeneralChartData(lessonName string, lessonAnalysis models.LessonAnalysis, chartData *models.GeneralChartData) {
	lessonData := chartData.Lessons[lessonName]

	lessonData.MaxNet = math.Max(lessonData.MaxNet, lessonAnalysis.Net)
	lessonData.AverageTime = (lessonData.AverageTime*(chartData.ExamCount) + lessonAnalysis.Time) / (chartData.ExamCount + 1)
	lessonData.AverageNet = (lessonData.AverageNet*float64(chartData.ExamCount) + lessonAnalysis.Net) / float64(chartData.ExamCount+1)

	chartData.Lessons[lessonName] = lessonData
}

func calculateAverage[T int | float64](oldAverage, newValue, itemCount T) T {
	return (oldAverage*(itemCount) + newValue) / (itemCount + 1)
}
