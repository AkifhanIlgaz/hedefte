package analysis

import "time"

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

func (a TYTAnalysis) GetDate() time.Time {
	return a.Date
}

func (a TYTAnalysis) GetName() string {
	return a.Name
}

func (a TYTAnalysis) GetTotalNet() float64 {
	return a.TotalNet
}

func (a TYTAnalysis) ApplyAnalysisToGeneralChartData(chartData *GeneralChartData) {
	applyAnalysisToGeneralChartData(a, chartData)
}

func (a TYTAnalysis) ApplyAllLessonsToChartData(chartData *TytAllLessonsChartData) {
	applyAnalysisToLessonChartData(a, a.Türkçe, &chartData.Türkçe, chartData.ExamCount)
	applyAnalysisToLessonChartData(a, a.Tarih, &chartData.Tarih, chartData.ExamCount)
	applyAnalysisToLessonChartData(a, a.Coğrafya, &chartData.Coğrafya, chartData.ExamCount)
	applyAnalysisToLessonChartData(a, a.Felsefe, &chartData.Felsefe, chartData.ExamCount)
	applyAnalysisToLessonChartData(a, a.DinKültürü, &chartData.DinKültürü, chartData.ExamCount)
	applyAnalysisToLessonChartData(a, a.Matematik, &chartData.Matematik, chartData.ExamCount)
	applyAnalysisToLessonChartData(a, a.Fizik, &chartData.Fizik, chartData.ExamCount)
	applyAnalysisToLessonChartData(a, a.Kimya, &chartData.Kimya, chartData.ExamCount)
	applyAnalysisToLessonChartData(a, a.Biyoloji, &chartData.Biyoloji, chartData.ExamCount)

	chartData.ExamCount++
}
