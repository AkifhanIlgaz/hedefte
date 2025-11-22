package models

import "github.com/AkifhanIlgaz/hedefte/internal/models/analysis"

type TytAnalysis = analysis.TYTAnalysis
type AddTytAnalysis = analysis.AddTYTAnalysis

type AytAnalysis = analysis.AYTAnalysis
type AddAytAnalysis = analysis.AddAYTAnalysis

type LessonAnalysis = analysis.LessonAnalysis
type TopicMistake = analysis.TopicMistake

// Chart Data types
type TytGeneralChartData = analysis.TytGeneralChartData
type AytGeneralChartData = analysis.AytGeneralChartData

type GeneralChartExam = analysis.GeneralChartExam
type LessonChartData = analysis.LessonChartData

// Constructor functions
var NewLessonChartData = analysis.NewLessonChartData
var NewTytGeneralChartData = analysis.NewTytGeneralChartData
var NewAytGeneralChartData = analysis.NewAytGeneralChartData
