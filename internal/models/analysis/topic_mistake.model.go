package analysis

type TopicMistake struct {
	TopicName    string `json:"topicName" bson:"topic_name" `
	MistakeCount int    `json:"mistakeCount" bson:"mistake_count" `
}
