package tyt_models

import (
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
)

type AddExamRequest struct {
	UserID     string                `json:"-" bson:"userId" `
	TotalNet   float64               `json:"-" bson:"total_net"`
	Date       time.Time             `json:"date" bson:"date" binding:"required" validate:"required"`
	Name       string                `json:"name" bson:"name" binding:"required" validate:"required"`
	Türkçe     models.LessonAnalysis `json:"Türkçe" bson:"turkce" binding:"required" validate:"required"`
	Tarih      models.LessonAnalysis `json:"Tarih" bson:"tarih" binding:"required" validate:"required"`
	Coğrafya   models.LessonAnalysis `json:"Coğrafya" bson:"cografya" binding:"required" validate:"required"`
	Felsefe    models.LessonAnalysis `json:"Felsefe" bson:"felsefe" binding:"required" validate:"required"`
	DinKültürü models.LessonAnalysis `json:"Din Kültürü" bson:"din_kulturu" binding:"required" validate:"required"`
	Matematik  models.LessonAnalysis `json:"Matematik" bson:"matematik" binding:"required" validate:"required"`
	Fizik      models.LessonAnalysis `json:"Fizik" bson:"fizik" binding:"required" validate:"required"`
	Kimya      models.LessonAnalysis `json:"Kimya" bson:"kimya" binding:"required" validate:"required"`
	Biyoloji   models.LessonAnalysis `json:"Biyoloji" bson:"biyoloji" binding:"required" validate:"required"`
}

func (a AddExamRequest) ToExam() Exam {
	a.CalculateNet()

	return Exam{
		UserId:     a.UserID,
		Date:       a.Date,
		Name:       a.Name,
		TotalNet:   a.TotalNet,
		Türkçe:     a.Türkçe,
		Tarih:      a.Tarih,
		Coğrafya:   a.Coğrafya,
		Felsefe:    a.Felsefe,
		DinKültürü: a.DinKültürü,
		Matematik:  a.Matematik,
		Fizik:      a.Fizik,
		Kimya:      a.Kimya,
		Biyoloji:   a.Biyoloji,
	}
}

func (req *AddExamRequest) CalculateNet() {
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
