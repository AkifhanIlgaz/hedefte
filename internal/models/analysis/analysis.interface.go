package analysis

import (
	"time"
)

type Analysis interface {
	GetDate() time.Time
	GetName() string
	GetTotalNet() float64
	ApplyAnalysisToGeneralChartData(chartData *GeneralChartData)
}
