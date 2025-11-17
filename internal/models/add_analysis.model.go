package models

import (
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/constants"
)

type AddAnalysisRequest interface {
	CollectionName() string
}

type AddTYTAnalysis struct {
	UserID     string         `json:"-" bson:"userId" `
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

type AddAYTAnalysis struct {
	UserID    string         `json:"-" bson:"userId" `
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
