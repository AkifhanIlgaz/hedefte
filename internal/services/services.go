package services

import (
	"github.com/AkifhanIlgaz/hedefte/internal/services/ayt"
	"github.com/AkifhanIlgaz/hedefte/internal/services/study_material"
	"github.com/AkifhanIlgaz/hedefte/internal/services/tyt"
)

type TYTAnalysisService = tyt.AnalysisService
type TYTLessonService = tyt.LessonService
type TYTTopicService = tyt.TopicService

type AYTAnalysisService = ayt.AnalysisService
type AYTLessonService = ayt.LessonService
type AYTTopicService = ayt.TopicService

var (
	NewTYTAnalysisService = tyt.NewAnalysisService
	NewTYTLessonService   = tyt.NewLessonService
	NewTYTTopicService    = tyt.NewTopicService

	NewAYTAnalysisService = ayt.NewAnalysisService
	NewAYTLessonService   = ayt.NewLessonService
	NewAYTTopicService    = ayt.NewTopicService
)

type StudyMaterialService = study_material.StudyMaterialService

var (
	NewStudyMaterialService = study_material.NewStudyMaterialService
)
