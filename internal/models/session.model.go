package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Session struct {
	Id          bson.ObjectID `json:"id" bson:"_id,omitempty"`
	UserId      string        `json:"userId" bson:"user_id"`
	Exam        ExamType      `json:"exam" bson:"exam"`
	Type        string        `json:"type" bson:"type"`
	Lesson      string        `json:"lesson" bson:"lesson"`
	Topic       string        `json:"topic" bson:"topic"`
	Goal        string        `json:"goal" bson:"goal"`
	Date        time.Time     `json:"date" bson:"date"`
	Notes       string        `json:"notes" bson:"notes"`
	Duration    time.Duration `json:"duration" bson:"duration"`
	IsCompleted bool          `json:"isCompleted" bson:"is_completed"`
}

type AddSessionRequest struct {
	UserId      string        `json:"-"`
	Exam        ExamType      `json:"exam"`
	Type        string        `json:"type"`
	Lesson      string        `json:"lesson"`
	Topic       string        `json:"topic"`
	Goal        string        `json:"goal"`
	Date        time.Time     `json:"date"`
	Duration    time.Duration `json:"duration"`
	IsCompleted bool          `json:"isCompleted"`
}

type UpdateSessionRequest struct {
	Id          bson.ObjectID `json:"id"`
	UserId      string        `json:"-"`
	Exam        ExamType      `json:"exam"`
	Type        string        `json:"type"`
	Lesson      string        `json:"lesson"`
	Topic       string        `json:"topic"`
	Goal        string        `json:"goal"`
	Date        time.Time     `json:"date"`
	Duration    time.Duration `json:"duration"`
	Notes       string        `json:"notes" bson:"notes"`
	IsCompleted bool          `json:"isCompleted"`
}

type CompleteSessionRequest struct {
	Id       bson.ObjectID `json:"id"`
	UserId   string        `json:"-"`
	Duration time.Duration `json:"duration"`
	Notes    string        `json:"notes" bson:"notes"`
}

func (req AddSessionRequest) ToSession() Session {
	return Session{
		UserId:      req.UserId,
		Exam:        req.Exam,
		Type:        req.Type,
		Lesson:      req.Lesson,
		Topic:       req.Topic,
		Goal:        req.Goal,
		Date:        req.Date,
		Duration:    req.Duration,
		IsCompleted: req.IsCompleted,
	}
}

func (req UpdateSessionRequest) ToSession() Session {
	return Session{
		Id:          req.Id,
		UserId:      req.UserId,
		Exam:        req.Exam,
		Type:        req.Type,
		Lesson:      req.Lesson,
		Topic:       req.Topic,
		Goal:        req.Goal,
		Date:        req.Date,
		Duration:    req.Duration,
		Notes:       req.Notes,
		IsCompleted: req.IsCompleted,
	}
}
