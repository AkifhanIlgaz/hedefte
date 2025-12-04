package services

// func GetLessonChartData(req models.ChartDataQuery) (models.LessonSpecificChartData, error) {
// 	exams, err := s.repo.FindExamsOfLesson(req.ExamType, req.UserId, req.Lesson, req.GetStart(), req.GetEnd())
// 	if err != nil {
// 		s.logger.Error("failed to get exams of tyt lesson", zap.Error(err))
// 		return models.LessonSpecificChartData{}, fmt.Errorf(`failed to get exams of tyt lesson: %w`, err)
// 	}

// 	chartData := models.NewLessonSpecificChartData()
// 	for _, exam := range exams {
// 		exam.ApplyToLessonSpecificChartData(&chartData)
// 	}

// 	return chartData, nil
// }
