package models

import (
	"github.com/AkifhanIlgaz/hedefte/internal/models/analysis"
	"github.com/AkifhanIlgaz/hedefte/internal/models/study_material"
)

type Exam = analysis.Exam
type Lesson = analysis.Lesson
type LessonAnalysis = analysis.LessonAnalysis
type Topic = analysis.Topic
type TopicAnalysis = analysis.TopicAnalysis

type StudyMaterial = study_material.StudyMaterial
type AddStudyMaterialRequest = study_material.AddStudyMaterialRequest
type DeleteStudyMaterialRequest = study_material.DeleteStudyMaterialRequest
type GetStudyMaterialsRequest = study_material.GetStudyMaterialsRequest
