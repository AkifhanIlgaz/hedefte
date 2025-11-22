package analysis

import "time"

type TytGeneralChartData struct {
	ExamCount  int                `json:"examCount"`
	MaxNet     float64            `json:"maxNet"`
	AverageNet float64            `json:"averageNet"`
	Exams      []GeneralChartExam `json:"exams"`
	Türkçe     LessonChartData    `json:"Türkçe" bson:"turkce"`
	Tarih      LessonChartData    `json:"Tarih" bson:"tarih"`
	Coğrafya   LessonChartData    `json:"Coğrafya" bson:"cografya"`
	Felsefe    LessonChartData    `json:"Felsefe" bson:"felsefe"`
	DinKültürü LessonChartData    `json:"Din Kültürü" bson:"din_kulturu"`
	Matematik  LessonChartData    `json:"Matematik" bson:"matematik"`
	Fizik      LessonChartData    `json:"Fizik" bson:"fizik"`
	Kimya      LessonChartData    `json:"Kimya" bson:"kimya"`
	Biyoloji   LessonChartData    `json:"Biyoloji" bson:"biyoloji"`
}

type GeneralChartExam struct {
	Date     time.Time `json:"date"`
	Name     string    `json:"name"`
	TotalNet float64   `json:"totalNet,omitempty"`
}

func NewTytGeneralChartData() TytGeneralChartData {
	return TytGeneralChartData{
		Türkçe:     NewLessonChartData(),
		Tarih:      NewLessonChartData(),
		Coğrafya:   NewLessonChartData(),
		Felsefe:    NewLessonChartData(),
		DinKültürü: NewLessonChartData(),
		Matematik:  NewLessonChartData(),
		Fizik:      NewLessonChartData(),
		Kimya:      NewLessonChartData(),
		Biyoloji:   NewLessonChartData(),
	}
}

type AytGeneralChartData struct {
	ExamCount  int                `json:"examCount"`
	MaxNet     float64            `json:"maxNet"`
	AverageNet float64            `json:"averageNet"`
	Exams      []GeneralChartExam `json:"exams"`
	Edebiyat   LessonChartData    `json:"Edebiyat" bson:"edebiyat,omitempty"`
	Tarih      LessonChartData    `json:"Tarih" bson:"tarih,omitempty"`
	Coğrafya   LessonChartData    `json:"Coğrafya" bson:"cografya,omitempty"`
	Matematik  LessonChartData    `json:"Matematik" bson:"matematik,omitempty"`
	Fizik      LessonChartData    `json:"Fizik" bson:"fizik,omitempty"`
	Kimya      LessonChartData    `json:"Kimya" bson:"kimya,omitempty"`
	Biyoloji   LessonChartData    `json:"Biyoloji" bson:"biyoloji,omitempty"`
}

func NewAytGeneralChartData() AytGeneralChartData {
	return AytGeneralChartData{
		Edebiyat:  NewLessonChartData(),
		Tarih:     NewLessonChartData(),
		Coğrafya:  NewLessonChartData(),
		Matematik: NewLessonChartData(),
		Fizik:     NewLessonChartData(),
		Kimya:     NewLessonChartData(),
		Biyoloji:  NewLessonChartData(),
	}
}

type LessonChartData struct {
	AverageNet    float64            `json:"averageNet"`
	AverageTime   int                `json:"averageTime"`
	MaxNet        float64            `json:"maxNet"`
	Exams         []GeneralChartExam `json:"exams,omitempty"`
	TopicMistakes map[string]int     `json:"topicMistakes"`
}

func NewLessonChartData() LessonChartData {
	return LessonChartData{
		AverageNet:    0,
		AverageTime:   0,
		MaxNet:        0,
		Exams:         []GeneralChartExam{},
		TopicMistakes: map[string]int{},
	}
}
