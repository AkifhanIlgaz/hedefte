package analysis

import "time"

type AYTAnalysis struct {
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

func (a AYTAnalysis) GetDate() time.Time {
	return a.Date
}

func (a AYTAnalysis) GetName() string {
	return a.Name
}

func (a AYTAnalysis) GetTotalNet() float64 {
	return a.TotalNet
}

func (a AYTAnalysis) ApplyAnalysisToGeneralChartData(chartData *GeneralChartData) {
	applyAnalysisToGeneralChartData(a, chartData)
}

func (a AYTAnalysis) ApplyAllLessonsToChartData(chartData *AytAllLessonsChartData) {
	applyAnalysisToLessonChartData(a, a.Edebiyat, &chartData.Edebiyat, chartData.ExamCount)
	applyAnalysisToLessonChartData(a, a.Tarih, &chartData.Tarih, chartData.ExamCount)
	applyAnalysisToLessonChartData(a, a.Coğrafya, &chartData.Coğrafya, chartData.ExamCount)
	applyAnalysisToLessonChartData(a, a.Matematik, &chartData.Matematik, chartData.ExamCount)
	applyAnalysisToLessonChartData(a, a.Fizik, &chartData.Fizik, chartData.ExamCount)
	applyAnalysisToLessonChartData(a, a.Kimya, &chartData.Kimya, chartData.ExamCount)
	applyAnalysisToLessonChartData(a, a.Biyoloji, &chartData.Biyoloji, chartData.ExamCount)

	chartData.ExamCount++
}
