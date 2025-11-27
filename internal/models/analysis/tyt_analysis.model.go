package analysis

import (
	"math"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Analysis interface {
	ApplyAnalysisToGeneralChartData(chartData *GeneralChartData)
}

type TYTAnalysis struct {
	Id         bson.ObjectID  `bson:"_id,omitempty"`
	UserId     string         `json:"userId" bson:"user_id"`
	Date       time.Time      `json:"date" bson:"date"`
	Name       string         `json:"name" bson:"name"`
	TotalNet   float64        `json:"totalNet" bson:"total_net"`
	Türkçe     LessonAnalysis `json:"Türkçe" bson:"turkce"`
	Tarih      LessonAnalysis `json:"Tarih" bson:"tarih"`
	Coğrafya   LessonAnalysis `json:"Coğrafya" bson:"cografya"`
	Felsefe    LessonAnalysis `json:"Felsefe" bson:"felsefe"`
	DinKültürü LessonAnalysis `json:"Din Kültürü" bson:"din_kulturu"`
	Matematik  LessonAnalysis `json:"Matematik" bson:"matematik"`
	Fizik      LessonAnalysis `json:"Fizik" bson:"fizik"`
	Kimya      LessonAnalysis `json:"Kimya" bson:"kimya"`
	Biyoloji   LessonAnalysis `json:"Biyoloji" bson:"biyoloji"`
}

func (a TYTAnalysis) ApplyAnalysisToGeneralChartData(chartData *GeneralChartData) {
	exam := GeneralChartExam{
		TotalNet: a.TotalNet,
		Date:     a.Date,
		Name:     a.Name,
	}

	chartData.MaxNet = math.Max(chartData.MaxNet, exam.TotalNet)
	chartData.AverageNet = calculateAverage(chartData.AverageNet, a.TotalNet, float64(chartData.ExamCount))
	chartData.Exams = append(chartData.Exams, exam)

	ApplyLessonAnalysisToGeneralChartData("Türkçe", a.Türkçe, chartData)
	ApplyLessonAnalysisToGeneralChartData("Tarih", a.Tarih, chartData)
	ApplyLessonAnalysisToGeneralChartData("Coğrafya", a.Coğrafya, chartData)
	ApplyLessonAnalysisToGeneralChartData("Felsefe", a.Felsefe, chartData)
	ApplyLessonAnalysisToGeneralChartData("Din Kültürü", a.DinKültürü, chartData)
	ApplyLessonAnalysisToGeneralChartData("Matematik", a.Matematik, chartData)
	ApplyLessonAnalysisToGeneralChartData("Fizik", a.Fizik, chartData)
	ApplyLessonAnalysisToGeneralChartData("Kimya", a.Kimya, chartData)
	ApplyLessonAnalysisToGeneralChartData("Biyoloji", a.Biyoloji, chartData)

	chartData.ExamCount++
}

func ApplyLessonAnalysisToGeneralChartData(lessonName string, lessonAnalysis LessonAnalysis, chartData *GeneralChartData) {
	lessonData := chartData.Lessons[lessonName]

	lessonData.MaxNet = math.Max(lessonData.MaxNet, lessonAnalysis.Net)
	lessonData.AverageTime = (lessonData.AverageTime*(chartData.ExamCount) + lessonAnalysis.Time) / (chartData.ExamCount + 1)
	lessonData.AverageNet = (lessonData.AverageNet*float64(chartData.ExamCount) + lessonAnalysis.Net) / float64(chartData.ExamCount+1)

	chartData.Lessons[lessonName] = lessonData
	// lessonData.Exams = append(lessonData.Exams, GeneralChartExam{
	// 	Date:     a.Date,
	// 	Name:     a.Name,
	// 	TotalNet: lessonAnalysis.Net,
	// })
	// for _, topicMistake := range lessonAnalysis.TopicMistakes {
	// 	lessonData.TopicMistakes[topicMistake.TopicName]++
	// }
}
