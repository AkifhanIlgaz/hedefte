package services

import (
	"github.com/AkifhanIlgaz/hedefte/internal/services/ayt"
	"github.com/AkifhanIlgaz/hedefte/internal/services/tyt"
)

type TYTAnalysisService = tyt.AnalysisService

type AYTAnalysisService = ayt.AnalysisService

var (
	NewTYTAnalysisService = tyt.NewAnalysisService
	NewAYTAnalysisService = ayt.NewAnalysisService
)
