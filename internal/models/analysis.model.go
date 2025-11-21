package models

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

func (a TYTAnalysis) ApplyToGeneralChartData(chartData *GeneralChartData) {
	exam := GeneralChartExam{
		TotalNet: a.TotalNet,
		Date:     a.Date,
		Name:     a.Name,
	}

	chartData.MaxNet = math.Max(chartData.MaxNet, exam.TotalNet)
	chartData.AverageNet = (chartData.AverageNet*float64(chartData.ExamCount) + exam.TotalNet) / float64(chartData.ExamCount+1)
	chartData.ExamCount++
	chartData.Exams = append(chartData.Exams, exam)
}

func (a TYTAnalysis) ApplyToAllLessonsChartData(chartData *TytAllLessonsChartData) {
	chartData.MaxNet = math.Max(chartData.MaxNet, a.TotalNet)
	chartData.AverageNet = (chartData.AverageNet*float64(chartData.ExamCount) + a.TotalNet) / float64(chartData.ExamCount+1)
	defer func() {
		chartData.ExamCount++
	}()

	{
		if chartData.Türkçe.TopicMistakes == nil {
			chartData.Türkçe.TopicMistakes = make(map[string]int)
		}
		chartData.Türkçe.MaxNet = math.Max(chartData.Türkçe.MaxNet, a.Türkçe.Net)
		chartData.Türkçe.AverageTime = (chartData.Türkçe.AverageTime*(chartData.ExamCount) + a.Türkçe.Time) / (chartData.ExamCount + 1)
		chartData.Türkçe.AverageNet = (chartData.Türkçe.AverageNet*float64(chartData.ExamCount) + a.Türkçe.Net) / float64(chartData.ExamCount+1)
		chartData.Türkçe.Exams = append(chartData.Türkçe.Exams, GeneralChartExam{
			Date:     a.Date,
			Name:     a.Name,
			TotalNet: a.Türkçe.Net,
		})
		for _, topicMistake := range a.Türkçe.TopicMistakes {
			chartData.Türkçe.TopicMistakes[topicMistake.TopicName]++
		}
	}

	{
		if chartData.Tarih.TopicMistakes == nil {
			chartData.Tarih.TopicMistakes = make(map[string]int)
		}
		chartData.Tarih.MaxNet = math.Max(chartData.Tarih.MaxNet, a.Tarih.Net)
		chartData.Tarih.AverageTime = (chartData.Tarih.AverageTime*(chartData.ExamCount) + a.Tarih.Time) / (chartData.ExamCount + 1)
		chartData.Tarih.AverageNet = (chartData.Tarih.AverageNet*float64(chartData.ExamCount) + a.Tarih.Net) / float64(chartData.ExamCount+1)
		chartData.Tarih.Exams = append(chartData.Tarih.Exams, GeneralChartExam{
			Date:     a.Date,
			Name:     a.Name,
			TotalNet: a.Tarih.Net,
		})
		for _, topicMistake := range a.Tarih.TopicMistakes {
			chartData.Tarih.TopicMistakes[topicMistake.TopicName]++
		}
	}

	{
		if chartData.Coğrafya.TopicMistakes == nil {
			chartData.Coğrafya.TopicMistakes = make(map[string]int)
		}
		chartData.Coğrafya.MaxNet = math.Max(chartData.Coğrafya.MaxNet, a.Coğrafya.Net)
		chartData.Coğrafya.AverageTime = (chartData.Coğrafya.AverageTime*(chartData.ExamCount) + a.Coğrafya.Time) / (chartData.ExamCount + 1)
		chartData.Coğrafya.AverageNet = (chartData.Coğrafya.AverageNet*float64(chartData.ExamCount) + a.Coğrafya.Net) / float64(chartData.ExamCount+1)
		chartData.Coğrafya.Exams = append(chartData.Coğrafya.Exams, GeneralChartExam{
			Date:     a.Date,
			Name:     a.Name,
			TotalNet: a.Coğrafya.Net,
		})
		for _, topicMistake := range a.Coğrafya.TopicMistakes {
			chartData.Coğrafya.TopicMistakes[topicMistake.TopicName]++
		}
	}

	{
		if chartData.Felsefe.TopicMistakes == nil {
			chartData.Felsefe.TopicMistakes = make(map[string]int)
		}
		chartData.Felsefe.MaxNet = math.Max(chartData.Felsefe.MaxNet, a.Felsefe.Net)
		chartData.Felsefe.AverageTime = (chartData.Felsefe.AverageTime*(chartData.ExamCount) + a.Felsefe.Time) / (chartData.ExamCount + 1)
		chartData.Felsefe.AverageNet = (chartData.Felsefe.AverageNet*float64(chartData.ExamCount) + a.Felsefe.Net) / float64(chartData.ExamCount+1)
		chartData.Felsefe.Exams = append(chartData.Felsefe.Exams, GeneralChartExam{
			Date:     a.Date,
			Name:     a.Name,
			TotalNet: a.Felsefe.Net,
		})
		for _, topicMistake := range a.Felsefe.TopicMistakes {
			chartData.Felsefe.TopicMistakes[topicMistake.TopicName]++
		}
	}

	{
		if chartData.DinKültürü.TopicMistakes == nil {
			chartData.DinKültürü.TopicMistakes = make(map[string]int)
		}
		chartData.DinKültürü.MaxNet = math.Max(chartData.DinKültürü.MaxNet, a.DinKültürü.Net)
		chartData.DinKültürü.AverageTime = (chartData.DinKültürü.AverageTime*(chartData.ExamCount) + a.DinKültürü.Time) / (chartData.ExamCount + 1)
		chartData.DinKültürü.AverageNet = (chartData.DinKültürü.AverageNet*float64(chartData.ExamCount) + a.DinKültürü.Net) / float64(chartData.ExamCount+1)
		chartData.DinKültürü.Exams = append(chartData.DinKültürü.Exams, GeneralChartExam{
			Date:     a.Date,
			Name:     a.Name,
			TotalNet: a.DinKültürü.Net,
		})
		for _, topicMistake := range a.DinKültürü.TopicMistakes {
			chartData.DinKültürü.TopicMistakes[topicMistake.TopicName]++
		}
	}

	{
		if chartData.Matematik.TopicMistakes == nil {
			chartData.Matematik.TopicMistakes = make(map[string]int)
		}
		chartData.Matematik.MaxNet = math.Max(chartData.Matematik.MaxNet, a.Matematik.Net)
		chartData.Matematik.AverageTime = (chartData.Matematik.AverageTime*(chartData.ExamCount) + a.Matematik.Time) / (chartData.ExamCount + 1)
		chartData.Matematik.AverageNet = (chartData.Matematik.AverageNet*float64(chartData.ExamCount) + a.Matematik.Net) / float64(chartData.ExamCount+1)
		chartData.Matematik.Exams = append(chartData.Matematik.Exams, GeneralChartExam{
			Date:     a.Date,
			Name:     a.Name,
			TotalNet: a.Matematik.Net,
		})
		for _, topicMistake := range a.Matematik.TopicMistakes {
			chartData.Matematik.TopicMistakes[topicMistake.TopicName]++
		}
	}

	{
		if chartData.Fizik.TopicMistakes == nil {
			chartData.Fizik.TopicMistakes = make(map[string]int)
		}
		chartData.Fizik.MaxNet = math.Max(chartData.Fizik.MaxNet, a.Fizik.Net)
		chartData.Fizik.AverageTime = (chartData.Fizik.AverageTime*(chartData.ExamCount) + a.Fizik.Time) / (chartData.ExamCount + 1)
		chartData.Fizik.AverageNet = (chartData.Fizik.AverageNet*float64(chartData.ExamCount) + a.Fizik.Net) / float64(chartData.ExamCount+1)
		chartData.Fizik.Exams = append(chartData.Fizik.Exams, GeneralChartExam{
			Date:     a.Date,
			Name:     a.Name,
			TotalNet: a.Fizik.Net,
		})
		for _, topicMistake := range a.Fizik.TopicMistakes {
			chartData.Fizik.TopicMistakes[topicMistake.TopicName]++
		}
	}

	{
		if chartData.Kimya.TopicMistakes == nil {
			chartData.Kimya.TopicMistakes = make(map[string]int)
		}
		chartData.Kimya.MaxNet = math.Max(chartData.Kimya.MaxNet, a.Kimya.Net)
		chartData.Kimya.AverageTime = (chartData.Kimya.AverageTime*(chartData.ExamCount) + a.Kimya.Time) / (chartData.ExamCount + 1)
		chartData.Kimya.AverageNet = (chartData.Kimya.AverageNet*float64(chartData.ExamCount) + a.Kimya.Net) / float64(chartData.ExamCount+1)
		chartData.Kimya.Exams = append(chartData.Kimya.Exams, GeneralChartExam{
			Date:     a.Date,
			Name:     a.Name,
			TotalNet: a.Kimya.Net,
		})
		for _, topicMistake := range a.Kimya.TopicMistakes {
			chartData.Kimya.TopicMistakes[topicMistake.TopicName]++
		}
	}

	{
		if chartData.Biyoloji.TopicMistakes == nil {
			chartData.Biyoloji.TopicMistakes = make(map[string]int)
		}
		chartData.Biyoloji.MaxNet = math.Max(chartData.Biyoloji.MaxNet, a.Biyoloji.Net)
		chartData.Biyoloji.AverageTime = (chartData.Biyoloji.AverageTime*(chartData.ExamCount) + a.Biyoloji.Time) / (chartData.ExamCount + 1)
		chartData.Biyoloji.AverageNet = (chartData.Biyoloji.AverageNet*float64(chartData.ExamCount) + a.Biyoloji.Net) / float64(chartData.ExamCount+1)
		chartData.Biyoloji.Exams = append(chartData.Biyoloji.Exams, GeneralChartExam{
			Date:     a.Date,
			Name:     a.Name,
			TotalNet: a.Biyoloji.Net,
		})
		for _, topicMistake := range a.Biyoloji.TopicMistakes {
			chartData.Biyoloji.TopicMistakes[topicMistake.TopicName]++
		}
	}

}

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

func (a AYTAnalysis) ApplyToAllLessonsChartData(chartData *AytAllLessonsChartData) {
	chartData.MaxNet = math.Max(chartData.MaxNet, a.TotalNet)
	chartData.AverageNet = (chartData.AverageNet*float64(chartData.ExamCount) + a.TotalNet) / float64(chartData.ExamCount+1)
	chartData.ExamCount++

	{
		if chartData.Edebiyat.TopicMistakes == nil {
			chartData.Edebiyat.TopicMistakes = make(map[string]int)
		}
		chartData.Edebiyat.MaxNet = math.Max(chartData.Edebiyat.MaxNet, a.Edebiyat.Net)
		chartData.Edebiyat.AverageTime = (chartData.Edebiyat.AverageTime*(chartData.ExamCount) + a.Edebiyat.Time) / (chartData.ExamCount + 1)
		chartData.Edebiyat.AverageNet = (chartData.Edebiyat.AverageNet*float64(chartData.ExamCount) + a.Edebiyat.Net) / float64(chartData.ExamCount+1)
		chartData.Edebiyat.Exams = append(chartData.Edebiyat.Exams, GeneralChartExam{
			Date:     a.Date,
			Name:     a.Name,
			TotalNet: a.Edebiyat.Net,
		})
		for _, topicMistake := range a.Edebiyat.TopicMistakes {
			chartData.Edebiyat.TopicMistakes[topicMistake.TopicName]++
		}
	}

	{
		if chartData.Tarih.TopicMistakes == nil {
			chartData.Tarih.TopicMistakes = make(map[string]int)
		}
		chartData.Tarih.MaxNet = math.Max(chartData.Tarih.MaxNet, a.Tarih.Net)
		chartData.Tarih.AverageTime = (chartData.Tarih.AverageTime*(chartData.ExamCount) + a.Tarih.Time) / (chartData.ExamCount + 1)
		chartData.Tarih.AverageNet = (chartData.Tarih.AverageNet*float64(chartData.ExamCount) + a.Tarih.Net) / float64(chartData.ExamCount+1)
		chartData.Tarih.Exams = append(chartData.Tarih.Exams, GeneralChartExam{
			Date:     a.Date,
			Name:     a.Name,
			TotalNet: a.Tarih.Net,
		})
		for _, topicMistake := range a.Tarih.TopicMistakes {
			chartData.Tarih.TopicMistakes[topicMistake.TopicName]++
		}
	}

	{
		if chartData.Coğrafya.TopicMistakes == nil {
			chartData.Coğrafya.TopicMistakes = make(map[string]int)
		}
		chartData.Coğrafya.MaxNet = math.Max(chartData.Coğrafya.MaxNet, a.Coğrafya.Net)
		chartData.Coğrafya.AverageTime = (chartData.Coğrafya.AverageTime*(chartData.ExamCount) + a.Coğrafya.Time) / (chartData.ExamCount + 1)
		chartData.Coğrafya.AverageNet = (chartData.Coğrafya.AverageNet*float64(chartData.ExamCount) + a.Coğrafya.Net) / float64(chartData.ExamCount+1)
		chartData.Coğrafya.Exams = append(chartData.Coğrafya.Exams, GeneralChartExam{
			Date:     a.Date,
			Name:     a.Name,
			TotalNet: a.Coğrafya.Net,
		})
		for _, topicMistake := range a.Coğrafya.TopicMistakes {
			chartData.Coğrafya.TopicMistakes[topicMistake.TopicName]++
		}
	}

	{
		if chartData.Matematik.TopicMistakes == nil {
			chartData.Matematik.TopicMistakes = make(map[string]int)
		}
		chartData.Matematik.MaxNet = math.Max(chartData.Matematik.MaxNet, a.Matematik.Net)
		chartData.Matematik.AverageTime = (chartData.Matematik.AverageTime*(chartData.ExamCount) + a.Matematik.Time) / (chartData.ExamCount + 1)
		chartData.Matematik.AverageNet = (chartData.Matematik.AverageNet*float64(chartData.ExamCount) + a.Matematik.Net) / float64(chartData.ExamCount+1)
		chartData.Matematik.Exams = append(chartData.Matematik.Exams, GeneralChartExam{
			Date:     a.Date,
			Name:     a.Name,
			TotalNet: a.Matematik.Net,
		})
		for _, topicMistake := range a.Matematik.TopicMistakes {
			chartData.Matematik.TopicMistakes[topicMistake.TopicName]++
		}
	}

	{
		if chartData.Fizik.TopicMistakes == nil {
			chartData.Fizik.TopicMistakes = make(map[string]int)
		}
		chartData.Fizik.MaxNet = math.Max(chartData.Fizik.MaxNet, a.Fizik.Net)
		chartData.Fizik.AverageTime = (chartData.Fizik.AverageTime*(chartData.ExamCount) + a.Fizik.Time) / (chartData.ExamCount + 1)
		chartData.Fizik.AverageNet = (chartData.Fizik.AverageNet*float64(chartData.ExamCount) + a.Fizik.Net) / float64(chartData.ExamCount+1)
		chartData.Fizik.Exams = append(chartData.Fizik.Exams, GeneralChartExam{
			Date:     a.Date,
			Name:     a.Name,
			TotalNet: a.Fizik.Net,
		})
		for _, topicMistake := range a.Fizik.TopicMistakes {
			chartData.Fizik.TopicMistakes[topicMistake.TopicName]++
		}
	}

	{
		if chartData.Kimya.TopicMistakes == nil {
			chartData.Kimya.TopicMistakes = make(map[string]int)
		}
		chartData.Kimya.TopicMistakes = map[string]int{}
		chartData.Kimya.MaxNet = math.Max(chartData.Kimya.MaxNet, a.Kimya.Net)
		chartData.Kimya.AverageTime = (chartData.Kimya.AverageTime*(chartData.ExamCount) + a.Kimya.Time) / (chartData.ExamCount + 1)
		chartData.Kimya.AverageNet = (chartData.Kimya.AverageNet*float64(chartData.ExamCount) + a.Kimya.Net) / float64(chartData.ExamCount+1)
		chartData.Kimya.Exams = append(chartData.Kimya.Exams, GeneralChartExam{
			Date:     a.Date,
			Name:     a.Name,
			TotalNet: a.Kimya.Net,
		})
		for _, topicMistake := range a.Kimya.TopicMistakes {
			chartData.Kimya.TopicMistakes[topicMistake.TopicName]++
		}
	}

	{
		if chartData.Biyoloji.TopicMistakes == nil {
			chartData.Biyoloji.TopicMistakes = make(map[string]int)
		}
		chartData.Biyoloji.TopicMistakes = map[string]int{}
		chartData.Biyoloji.MaxNet = math.Max(chartData.Biyoloji.MaxNet, a.Biyoloji.Net)
		chartData.Biyoloji.AverageTime = (chartData.Biyoloji.AverageTime*(chartData.ExamCount) + a.Biyoloji.Time) / (chartData.ExamCount + 1)
		chartData.Biyoloji.AverageNet = (chartData.Biyoloji.AverageNet*float64(chartData.ExamCount) + a.Biyoloji.Net) / float64(chartData.ExamCount+1)
		chartData.Biyoloji.Exams = append(chartData.Biyoloji.Exams, GeneralChartExam{
			Date:     a.Date,
			Name:     a.Name,
			TotalNet: a.Biyoloji.Net,
		})
		for _, topicMistake := range a.Biyoloji.TopicMistakes {
			chartData.Biyoloji.TopicMistakes[topicMistake.TopicName]++
		}
	}
}

func (a AYTAnalysis) ApplyToGeneralChartData(chartData *GeneralChartData) {
	exam := GeneralChartExam{
		TotalNet: a.TotalNet,
		Date:     a.Date,
		Name:     a.Name,
	}

	chartData.MaxNet = math.Max(chartData.MaxNet, exam.TotalNet)
	chartData.AverageNet = (chartData.AverageNet*float64(chartData.ExamCount) + exam.TotalNet) / float64(chartData.ExamCount+1)
	chartData.ExamCount++
	chartData.Exams = append(chartData.Exams, exam)
}
