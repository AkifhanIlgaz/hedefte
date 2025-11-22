package analysis

import (
	"math"
	"time"
)

type TYTAnalysis struct {
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

func (a TYTAnalysis) ApplyAnalysisToGeneralChartData(chartData *TytGeneralChartData) {
	exam := GeneralChartExam{
		TotalNet: a.TotalNet,
		Date:     a.Date,
		Name:     a.Name,
	}

	chartData.MaxNet = math.Max(chartData.MaxNet, exam.TotalNet)
	chartData.AverageNet = calculateAverage(chartData.AverageNet, a.TotalNet, float64(chartData.ExamCount))

	chartData.Exams = append(chartData.Exams, exam)

	a.ApplyLessonAnalysisToTytChartData(a.Türkçe, &chartData.Türkçe, chartData.ExamCount)
	a.ApplyLessonAnalysisToTytChartData(a.Tarih, &chartData.Tarih, chartData.ExamCount)
	a.ApplyLessonAnalysisToTytChartData(a.Coğrafya, &chartData.Coğrafya, chartData.ExamCount)
	a.ApplyLessonAnalysisToTytChartData(a.Felsefe, &chartData.Felsefe, chartData.ExamCount)
	a.ApplyLessonAnalysisToTytChartData(a.DinKültürü, &chartData.DinKültürü, chartData.ExamCount)
	a.ApplyLessonAnalysisToTytChartData(a.Matematik, &chartData.Matematik, chartData.ExamCount)
	a.ApplyLessonAnalysisToTytChartData(a.Fizik, &chartData.Fizik, chartData.ExamCount)
	a.ApplyLessonAnalysisToTytChartData(a.Kimya, &chartData.Kimya, chartData.ExamCount)
	a.ApplyLessonAnalysisToTytChartData(a.Biyoloji, &chartData.Biyoloji, chartData.ExamCount)

	chartData.ExamCount++
}

func (a TYTAnalysis) ApplyLessonAnalysisToTytChartData(lessonAnalysis LessonAnalysis, chartData *LessonChartData, examCount int) {
	chartData.MaxNet = math.Max(chartData.MaxNet, lessonAnalysis.Net)
	chartData.AverageTime = (chartData.AverageTime*(examCount) + lessonAnalysis.Time) / (examCount + 1)
	chartData.AverageNet = (chartData.AverageNet*float64(examCount) + lessonAnalysis.Net) / float64(examCount+1)
	chartData.Exams = append(chartData.Exams, GeneralChartExam{
		Date:     a.Date,
		Name:     a.Name,
		TotalNet: lessonAnalysis.Net,
	})
	for _, topicMistake := range lessonAnalysis.TopicMistakes {
		chartData.TopicMistakes[topicMistake.TopicName]++
	}
}
