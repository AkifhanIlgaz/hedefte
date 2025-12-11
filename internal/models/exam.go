package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

// MongoDB schema
type Exam struct {
	Id       bson.ObjectID `bson:"_id,omitempty"`
	UserId   string        `bson:"user_id"`
	ExamType string        `bson:"exam_type"`
	Date     time.Time     `bson:"date"`
	Name     string        `bson:"name"`
	Result   float64       `bson:"result"`
	Lessons  []Lesson      `bson:"lessons"`
}

func (e Exam) ToExamAnalyticsUpsert() UpsertExamAnalytics {
	return UpsertExamAnalytics{
		UserId:   e.UserId,
		ExamType: e.ExamType,
		Date:     e.Date,
		Name:     e.Name,
		Result:   e.Result,
	}
}

func (e Exam) Collection() string {
	return "exams"
}

type Lesson struct {
	Name    string  `bson:"name"`
	Correct int     `bson:"correct"`
	Wrong   int     `bson:"wrong"`
	Empty   int     `bson:"empty"`
	Time    int     `bson:"time"`
	Result  float64 `bson:"result"`
}

func (l Lesson) ToLessonAnalyticsUpsert(userId, examType string, date time.Time) UpsertLessonAnalytics {
	return UpsertLessonAnalytics{
		UserId:   userId,
		ExamType: examType,
		Date:     date,
		Result:   l.Result,
		Lesson:   l.Name,
		Time:     l.Time,
	}
}

type TopicMistake struct {
	Id       bson.ObjectID `bson:"_id,omitempty"`
	Date     time.Time     `bson:"date"`
	ExamId   bson.ObjectID `bson:"exam_id"`
	UserId   string        `bson:"user_id"`
	ExamType string        `bson:"exam_type"`
	Lesson   string        `bson:"lesson"`
	Topic    string        `bson:"topic"`
	IsSolved bool          `bson:"is_solved"`
	ImageUrl string        `bson:"image_url"`
}

func (t TopicMistake) Collection() string {
	return "topic_mistakes"
}

type ExamAnalytics struct {
	Id            bson.ObjectID  `bson:"_id,omitempty"`
	UserId        string         `bson:"user_id"`
	ExamType      string         `bson:"exam_type"`
	ExamCount     int            `bson:"exam_count"`
	MaxResult     float64        `bson:"max_result"`
	AverageResult float64        `bson:"average_result"`
	ResultSeries  []ResultSeries `bson:"result_series"`
}

func (e ExamAnalytics) Collection() string {
	return "analytics"
}

type UpsertExamAnalytics struct {
	Date     time.Time `bson:"date"`
	Name     string    `bson:"name"`
	UserId   string    `bson:"user_id"`
	ExamType string    `bson:"exam_type"`
	Result   float64   `bson:"result"`
}

type DeleteExamAnalytics struct {
	Date     time.Time `bson:"date"`
	Name     string    `bson:"name"`
	UserId   string    `bson:"user_id"`
	ExamType string    `bson:"exam_type"`
	Result   float64   `bson:"result"`
}

type LessonAnalytics struct {
	Id            bson.ObjectID  `bson:"_id,omitempty"`
	UserId        string         `bson:"user_id"`
	ExamType      string         `bson:"exam_type"`
	ExamCount     int            `bson:"exam_count"`
	MaxResult     float64        `bson:"max_result"`
	AverageResult float64        `bson:"average_result"`
	ResultSeries  []ResultSeries `bson:"result_series"`
	Lesson        string         `bson:"lesson"`
	AverageTime   int            `bson:"average_time"`
}

type UpsertLessonAnalytics struct {
	Date     time.Time `bson:"date"`
	Name     string    `bson:"name"`
	UserId   string    `bson:"user_id"`
	ExamType string    `bson:"exam_type"`
	Lesson   string    `bson:"lesson"`
	Time     int       `bson:"time"`
	Result   float64   `bson:"result"`
}

type DeleteLessonAnalytics struct {
	Date     time.Time `bson:"date"`
	Name     string    `bson:"name"`
	UserId   string    `bson:"user_id"`
	ExamType string    `bson:"exam_type"`
	Lesson   string    `bson:"lesson"`
	Time     int       `bson:"time"`
	Result   float64   `bson:"result"`
}

type ResultSeries struct {
	Date   time.Time `bson:"date"`
	Name   string    `bson:"name"`
	Result float64   `bson:"result"`
	Time   int       `bson:"time,omitempty"`
}
