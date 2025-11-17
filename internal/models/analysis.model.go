package models

import "time"

type TYTAnalysis struct {
	Date       time.Time      `json:"date" bson:"date"`
	Name       string         `json:"name" bson:"name"`
	TotalNet   float64        `json:"totalNet" bson:"total_net"`
	Türkçe     LessonAnalysis `json:"Türkçe" bson:"turkce"`
	Tarih      LessonAnalysis `json:"Tarih" bson:"tarih"`
	Coğrafya   LessonAnalysis `json:"Coğrafya" bson:"cografya"`
	Felsefe    LessonAnalysis `json:"Felsefe" bson:"felsefe"`
	DinKültürü LessonAnalysis `json:"Din Kültürü" bson:"din_kulturu"`
	Matematik  LessonAnalysis `json:"Matematik" bson:"matematik"`
	Fizik      LessonAnalysis `json:"Fizik" bson:"fizik"`
	Kimya      LessonAnalysis `json:"Kimya" bson:"kimya"`
	Biyoloji   LessonAnalysis `json:"Biyoloji" bson:"biyoloji"`
}

type AYTAnalysis struct {
	Date      time.Time      `json:"date" bson:"date"`
	Name      string         `json:"name" bson:"name"`
	TotalNet  float64        `json:"totalNet" bson:"total_net"`
	Edebiyat  LessonAnalysis `json:"Edebiyat" bson:"edebiyat,omitempty"`
	Tarih     LessonAnalysis `json:"Tarih" bson:"tarih,omitempty"`
	Coğrafya  LessonAnalysis `json:"Coğrafya" bson:"cografya,omitempty"`
	Matematik LessonAnalysis `json:"Matematik" bson:"matematik,omitempty"`
	Fizik     LessonAnalysis `json:"Fizik" bson:"fizik,omitempty"`
	Kimya     LessonAnalysis `json:"Kimya" bson:"kimya,omitempty"`
	Biyoloji  LessonAnalysis `json:"Biyoloji" bson:"biyoloji,omitempty"`
}
