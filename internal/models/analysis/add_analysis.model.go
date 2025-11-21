package analysis

import (
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/constants"
)

type AddTYTAnalysis struct {
	UserID     string         `json:"-" bson:"userId" `
	TotalNet   float64        `json:"-" bson:"total_net"`
	Date       time.Time      `json:"date" bson:"date" binding:"required" validate:"required"`
	Name       string         `json:"name" bson:"name" binding:"required" validate:"required"`
	Türkçe     LessonAnalysis `json:"Türkçe" bson:"turkce" binding:"required" validate:"required"`
	Tarih      LessonAnalysis `json:"Tarih" bson:"tarih" binding:"required" validate:"required"`
	Coğrafya   LessonAnalysis `json:"Coğrafya" bson:"cografya" binding:"required" validate:"required"`
	Felsefe    LessonAnalysis `json:"Felsefe" bson:"felsefe" binding:"required" validate:"required"`
	DinKültürü LessonAnalysis `json:"Din Kültürü" bson:"din_kulturu" binding:"required" validate:"required"`
	Matematik  LessonAnalysis `json:"Matematik" bson:"matematik" binding:"required" validate:"required"`
	Fizik      LessonAnalysis `json:"Fizik" bson:"fizik" binding:"required" validate:"required"`
	Kimya      LessonAnalysis `json:"Kimya" bson:"kimya" binding:"required" validate:"required"`
	Biyoloji   LessonAnalysis `json:"Biyoloji" bson:"biyoloji" binding:"required" validate:"required"`
}

func (AddTYTAnalysis) CollectionName() string {
	return constants.TytAnalysisCollection
}

func (req *AddTYTAnalysis) CalculateNet() {
	req.Türkçe.CalculateNet()
	req.Tarih.CalculateNet()
	req.Coğrafya.CalculateNet()
	req.Felsefe.CalculateNet()
	req.DinKültürü.CalculateNet()
	req.Matematik.CalculateNet()
	req.Fizik.CalculateNet()
	req.Kimya.CalculateNet()
	req.Biyoloji.CalculateNet()

	req.TotalNet = req.Türkçe.Net + req.Tarih.Net + req.Coğrafya.Net + req.Felsefe.Net + req.DinKültürü.Net + req.Matematik.Net + req.Fizik.Net + req.Kimya.Net + req.Biyoloji.Net
}

type AddAYTAnalysis struct {
	UserID    string         `json:"-" bson:"userId" `
	TotalNet  float64        `json:"-" bson:"total_net"`
	Date      time.Time      `json:"date" bson:"date" binding:"required"`
	Name      string         `json:"name" bson:"name" binding:"required"`
	Edebiyat  LessonAnalysis `json:"Edebiyat" bson:"edebiyat,omitempty"`
	Tarih     LessonAnalysis `json:"Tarih" bson:"tarih,omitempty"`
	Coğrafya  LessonAnalysis `json:"Coğrafya" bson:"cografya,omitempty"`
	Matematik LessonAnalysis `json:"Matematik" bson:"matematik,omitempty"`
	Fizik     LessonAnalysis `json:"Fizik" bson:"fizik,omitempty"`
	Kimya     LessonAnalysis `json:"Kimya" bson:"kimya,omitempty"`
	Biyoloji  LessonAnalysis `json:"Biyoloji" bson:"biyoloji,omitempty"`
}

func (AddAYTAnalysis) CollectionName() string {
	return constants.AytAnalysisCollection
}

func (req *AddAYTAnalysis) CalculateNet() {
	req.Edebiyat.CalculateNet()
	req.Tarih.CalculateNet()
	req.Coğrafya.CalculateNet()
	req.Matematik.CalculateNet()
	req.Fizik.CalculateNet()
	req.Kimya.CalculateNet()
	req.Biyoloji.CalculateNet()

	req.TotalNet = req.Edebiyat.Net + req.Tarih.Net + req.Coğrafya.Net + req.Matematik.Net + req.Fizik.Net + req.Kimya.Net + req.Biyoloji.Net
}
