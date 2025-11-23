package models

import "github.com/AkifhanIlgaz/hedefte/internal/models/analysis"

type Analysis = analysis.Analysis

type TytAnalysis = analysis.TYTAnalysis
type AddTytAnalysis = analysis.AddTYTAnalysis

type AytAnalysis = analysis.AYTAnalysis
type AddAytAnalysis = analysis.AddAYTAnalysis

type LessonAnalysis = analysis.LessonAnalysis
type TopicMistake = analysis.TopicMistake

// Chart Data types
type GeneralChartData = analysis.GeneralChartData

type GeneralChartExam = analysis.GeneralChartExam
type LessonGeneralChartData = analysis.LessonGeneralChartData

// Constructor functions
var NewGeneralChartData = analysis.NewGeneralChartData
