package handlers

// // getChartDataByType retrieves chart data based on chart type and exam type
// func getChartDataByType(req models.ChartDataQuery) (any, error) {
// 	switch req.ChartType {
// 	case models.ChartTypeGeneral:
// 		switch req.ExamType {
// 		case models.ExamTypeTYT:
// 			return h.analysisService.GetTytGeneralChartData(req)
// 		case models.ExamTypeAYT:
// 			return h.analysisService.GetAytGeneralChartData(req)
// 		default:
// 			return nil, errors.New("invalid chart type")
// 		}
// 	case models.ChartTypeLessonSpecific:
// 		return h.analysisService.GetLessonChartData(req)
// 	default:
// 		return nil, errors.New("invalid chart type")
// 	}
// }
