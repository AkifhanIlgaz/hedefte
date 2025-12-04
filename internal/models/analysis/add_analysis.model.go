package analysis

import (
	"time"
)

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

func (a AddAYTAnalysis) ToAytAnalysis() AYTAnalysis {
	return AYTAnalysis{
		UserId:    a.UserID,
		Date:      a.Date,
		Name:      a.Name,
		TotalNet:  a.TotalNet,
		Edebiyat:  a.Edebiyat,
		Tarih:     a.Tarih,
		Coğrafya:  a.Coğrafya,
		Matematik: a.Matematik,
		Fizik:     a.Fizik,
		Kimya:     a.Kimya,
		Biyoloji:  a.Biyoloji,
	}
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
