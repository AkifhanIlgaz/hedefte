package analysis

import (
	"math"
	"time"
)

type LessonAnalysis struct {
	Correct       int            `json:"correct" bson:"correct" binding:"min=0"`
	Wrong         int            `json:"wrong" bson:"wrong" binding:"min=0"`
	Empty         int            `json:"empty" bson:"empty" binding:"min=0"`
	Time          int            `json:"time" bson:"time" `
	Net           float64        `json:"net" bson:"net"`
	TopicMistakes []TopicMistake `json:"topicMistakes" bson:"topic_mistakes"`
}

func (req *LessonAnalysis) CalculateNet() {
	req.Net = float64(req.Correct) - (float64(req.Wrong) * 0.25)
}

type LessonSpecificAnalysis struct {
	Date time.Time `bson:"date"`
	Name string    `bson:"name"`
	LessonAnalysis
}

func (a LessonSpecificAnalysis) ApplyToLessonSpecificChartData(chartData *LessonSpecificChartData) {
	chartData.MaxNet = math.Max(chartData.MaxNet, a.Net)
	chartData.AverageTime = (chartData.AverageTime*(chartData.ExamCount) + a.Time) / (chartData.ExamCount + 1)
	chartData.AverageNet = (chartData.AverageNet*float64(chartData.ExamCount) + a.Net) / float64(chartData.ExamCount+1)

	chartData.Exams = append(chartData.Exams, GeneralChartExam{
		Date:     a.Date,
		Name:     a.Name,
		TotalNet: a.LessonAnalysis.Net,
	})

	chartData.ExamCount++
}
