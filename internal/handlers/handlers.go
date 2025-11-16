package handlers

import (
	"github.com/AkifhanIlgaz/hedefte/internal/handlers/ayt"
	"github.com/AkifhanIlgaz/hedefte/internal/handlers/tyt"
)

type TYTAnalysisHandler = tyt.AnalysisHandler

type AYTAnalysisHandler = ayt.AnalysisHandler

var (
	NewTYTAnalysisHandler = tyt.NewAnalysisHandler

	NewAYTAnalysisHandler = ayt.NewAnalysisHandler
)
