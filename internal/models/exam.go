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

func (e Exam) ToExamResponse() ExamResponse {
	lessons := make([]LessonResponse, len(e.Lessons))
	for i, lesson := range e.Lessons {
		lessons[i] = lesson.ToLessonResponse()
	}

	return ExamResponse{
		Id:      e.Id.Hex(),
		Date:    e.Date,
		Name:    e.Name,
		Result:  e.Result,
		Lessons: lessons,
	}
}

func (e Exam) ToUpsertExamAnalytics() UpsertExamAnalytics {
	return UpsertExamAnalytics{
		ExamId:   e.Id,
		UserId:   e.UserId,
		ExamType: e.ExamType,
		Date:     e.Date,
		Name:     e.Name,
		Result:   e.Result,
	}
}

func (e Exam) ToDeleteExamAnalytics() DeleteExamAnalytics {
	return DeleteExamAnalytics{
		ExamId:   e.Id,
		UserId:   e.UserId,
		ExamType: e.ExamType,
		Result:   e.Result,
	}
}

func (e Exam) ToDeleteLessonAnalytics() []DeleteLessonAnalytics {
	deleteLessonAnalytics := make([]DeleteLessonAnalytics, len(e.Lessons))
	for i, lesson := range e.Lessons {
		deleteLessonAnalytics[i] = lesson.ToDeleteLessonAnalytics(e.UserId, e.ExamType, e.Id, e.Date)
	}
	return deleteLessonAnalytics
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

func (l Lesson) ToLessonResponse() LessonResponse {
	return LessonResponse{
		Name:    l.Name,
		Correct: l.Correct,
		Wrong:   l.Wrong,
		Empty:   l.Empty,
		Time:    l.Time,
		Result:  l.Result,
	}
}

func (l Lesson) ToUpsertLessonAnalytics(userId, examType string, examName string, examId bson.ObjectID, date time.Time) UpsertLessonAnalytics {
	return UpsertLessonAnalytics{
		ExamId:   examId,
		UserId:   userId,
		ExamType: examType,
		Date:     date,
		Name:     examName,
		Result:   l.Result,
		Lesson:   l.Name,
		Time:     l.Time,
	}
}

func (l Lesson) ToDeleteLessonAnalytics(userId, examType string, examId bson.ObjectID, date time.Time) DeleteLessonAnalytics {
	return DeleteLessonAnalytics{
		ExamId:   examId,
		UserId:   userId,
		ExamType: examType,
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
	ExamId   bson.ObjectID `bson:"exam_id"`
	Date     time.Time     `bson:"date"`
	Name     string        `bson:"name"`
	UserId   string        `bson:"user_id"`
	ExamType string        `bson:"exam_type"`
	Result   float64       `bson:"result"`
}

type DeleteExamAnalytics struct {
	ExamId   bson.ObjectID `bson:"exam_id"`
	UserId   string        `bson:"user_id"`
	ExamType string        `bson:"exam_type"`
	Result   float64       `bson:"result"`
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
	ExamId   bson.ObjectID `bson:"exam_id"`
	Date     time.Time     `bson:"date"`
	Name     string        `bson:"name"`
	UserId   string        `bson:"user_id"`
	ExamType string        `bson:"exam_type"`
	Lesson   string        `bson:"lesson"`
	Time     int           `bson:"time"`
	Result   float64       `bson:"result"`
}

type DeleteLessonAnalytics struct {
	ExamId   bson.ObjectID `bson:"exam_id"`
	UserId   string        `bson:"user_id"`
	ExamType string        `bson:"exam_type"`
	Lesson   string        `bson:"lesson"`
	Time     int           `bson:"time"`
	Result   float64       `bson:"result"`
}

type ResultSeries struct {
	ExamId bson.ObjectID `bson:"exam_id"`
	Date   time.Time     `bson:"date"`
	Name   string        `bson:"name"`
	Result float64       `bson:"result"`
	Time   int           `bson:"time,omitempty"`
}
