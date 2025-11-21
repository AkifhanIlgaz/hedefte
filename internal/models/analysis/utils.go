package analysis

import "math"

func calculateAverage[T int | float64](oldAverage, newValue, itemCount T) T {
	return (oldAverage*(itemCount) + newValue) / (itemCount + 1)
}

func applyAnalysisToLessonChartData(analyis Analysis, lessonAnalysis LessonAnalysis, chartData *LessonChartData, examCount int) {
	chartData.MaxNet = math.Max(chartData.MaxNet, lessonAnalysis.Net)
	chartData.AverageTime = (chartData.AverageTime*(examCount) + lessonAnalysis.Time) / (examCount + 1)
	chartData.AverageNet = (chartData.AverageNet*float64(examCount) + lessonAnalysis.Net) / float64(examCount+1)
	chartData.Exams = append(chartData.Exams, GeneralChartExam{
		Date:     analyis.GetDate(),
		Name:     analyis.GetName(),
		TotalNet: lessonAnalysis.Net,
	})
	for _, topicMistake := range lessonAnalysis.TopicMistakes {
		chartData.TopicMistakes[topicMistake.TopicName]++
	}
}

func applyAnalysisToGeneralChartData(analysis Analysis, chartData *GeneralChartData) {
	exam := GeneralChartExam{
		TotalNet: analysis.GetTotalNet(),
		Date:     analysis.GetDate(),
		Name:     analysis.GetName(),
	}

	chartData.MaxNet = math.Max(chartData.MaxNet, exam.TotalNet)
	chartData.AverageNet = calculateAverage(chartData.AverageNet, analysis.GetTotalNet(), float64(chartData.ExamCount))
	chartData.ExamCount++
	chartData.Exams = append(chartData.Exams, exam)
}
