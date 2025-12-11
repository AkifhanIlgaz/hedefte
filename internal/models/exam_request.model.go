package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type AddExamRequest struct {
	UserId   string              `json:"-"`
	ExamType string              `json:"examType"`
	Date     time.Time           `json:"date"`
	Name     string              `json:"name"`
	Lessons  []ExamRequestLesson `json:"lessons"`
}

func (req AddExamRequest) ToExam() Exam {
	result := 0.0
	lessons := make([]Lesson, len(req.Lessons))

	for i, reqLesson := range req.Lessons {
		lessons[i] = reqLesson.ToLesson()
		result += lessons[i].Result
	}

	return Exam{
		UserId:   req.UserId,
		ExamType: req.ExamType,
		Date:     req.Date,
		Name:     req.Name,
		Lessons:  lessons,
		Result:   result,
	}
}

func (req AddExamRequest) ExtractTopicMistakeRequests() []ExamRequestTopicMistake {
	topicMistakes := []ExamRequestTopicMistake{}
	for _, lesson := range req.Lessons {
		for _, topicMistake := range lesson.TopicMistakes {
			topicMistakes = append(topicMistakes, topicMistake)
		}
	}

	return topicMistakes
}

type ExamRequestLesson struct {
	Name          string                    `json:"name"`
	Correct       int                       `json:"correct"`
	Wrong         int                       `json:"wrong"`
	Empty         int                       `json:"empty"`
	Time          int                       `json:"time"`
	TopicMistakes []ExamRequestTopicMistake `json:"topicMistakes"`
}

func (req ExamRequestLesson) ToLesson() Lesson {
	result := float64(req.Correct) - (float64(req.Wrong) * 0.25)

	return Lesson{
		Name:    req.Name,
		Correct: req.Correct,
		Wrong:   req.Wrong,
		Empty:   req.Empty,
		Time:    req.Time,
		Result:  result,
	}
}

type ExamRequestTopicMistake struct {
	Date     time.Time `json:"date"`
	ExamType string    `json:"examType"`
	ExamId   string    `json:"examId"`
	Lesson   string    `json:"lesson"`
	Topic    string    `json:"topic"`
	IsSolved bool      `json:"isSolved"`
	ImageUrl string    `json:"imageUrl"`
}

func (req ExamRequestTopicMistake) ToTopicMistake(userId string) TopicMistake {
	examId, err := bson.ObjectIDFromHex(req.ExamId)
	if err != nil {
		examId = bson.NilObjectID
	}

	return TopicMistake{
		Date:     req.Date,
		ExamId:   examId,
		UserId:   userId,
		ExamType: req.ExamType,
		Lesson:   req.Lesson,
		Topic:    req.Topic,
		IsSolved: req.IsSolved,
		ImageUrl: req.ImageUrl,
	}
}
