package models

import (
	"github.com/AkifhanIlgaz/hedefte/internal/models/analysis"
)

type Analysis = analysis.Analysis

type TytAnalysis = analysis.TYTAnalysis

type AytAnalysis = analysis.AYTAnalysis
type AddAytAnalysis = analysis.AddAYTAnalysis

type LessonAnalysis = analysis.LessonAnalysis
type LessonSpecificAnalysis = analysis.LessonSpecificAnalysis
type TopicMistake = analysis.TopicMistake

// Chart Data types
type GeneralChartData = analysis.GeneralChartData

type GeneralChartExam = analysis.GeneralChartExam
type LessonGeneralChartData = analysis.LessonGeneralChartData

type LessonSpecificChartData = analysis.LessonSpecificChartData

// Constructor functions
var NewGeneralChartData = analysis.NewGeneralChartData
var NewLessonSpecificChartData = analysis.NewLessonSpecificChartData
