package models

type LessonAnalysis struct {
	Correct       int            `json:"correct" bson:"correct" binding:"min=0"`
	Wrong         int            `json:"wrong" bson:"wrong" binding:"min=0"`
	Empty         int            `json:"empty" bson:"empty" binding:"min=0"`
	Time          int            `json:"time" bson:"time" `
	TopicMistakes []TopicMistake `json:"topicMistakes" bson:"topic_mistakes"`
}
