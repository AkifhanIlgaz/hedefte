package analysis

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type TopicMistake struct {
	Id        bson.ObjectID `json:"id" bson:"_id,omitempty"`
	ExamId    bson.ObjectID `json:"examId" bson:"exam_id" `
	UserId    string        `json:"userId" bson:"user_id" `
	ExamType  string        `json:"examType" bson:"exam_type" `
	Lesson    string        `json:"lesson" bson:"lesson" `
	TopicName string        `json:"topicName" bson:"topic_name" `
	IsSolved  bool          `json:"isSolved" bson:"is_solved" `
	ImageUrl  string        `json:"imageUrl" bson:"image_url" `
}
