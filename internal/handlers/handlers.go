package handlers

import (
	"github.com/AkifhanIlgaz/hedefte/internal/handlers/ayt"
	"github.com/AkifhanIlgaz/hedefte/internal/handlers/study_material"
	"github.com/AkifhanIlgaz/hedefte/internal/handlers/tyt"
)

type TYTAnalysisHandler = tyt.AnalysisHandler
type TYTLessonHandler = tyt.LessonHandler
type TYTTopicHandler = tyt.TopicHandler

type AYTAnalysisHandler = ayt.AnalysisHandler
type AYTLessonHandler = ayt.LessonHandler
type AYTTopicHandler = ayt.TopicHandler

var (
	NewTYTAnalysisHandler = tyt.NewAnalysisHandler
	NewTYTLessonHandler   = tyt.NewLessonHandler
	NewTYTTopicHandler    = tyt.NewTopicHandler

	NewAYTAnalysisHandler = ayt.NewAnalysisHandler
	NewAYTLessonHandler   = ayt.NewLessonHandler
	NewAYTTopicHandler    = ayt.NewTopicHandler
)

type StudyMaterialHandler = study_material.StudyMaterialHandler

var (
	NewStudyMaterialHandler = study_material.NewStudyMaterialHandler
)
