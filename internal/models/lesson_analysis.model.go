package models

type LessonAnalysis struct {
	Correct       int            `json:"correct" bson:"correct" binding:"min=0"`
	Wrong         int            `json:"wrong" bson:"wrong" binding:"min=0"`
	Empty         int            `json:"empty" bson:"empty" binding:"min=0"`
	Time          int            `json:"time" bson:"time" `
	Net           float64        `json:"net" bson:"net"`
	TopicMistakes []TopicMistake `json:"topicMistakes" bson:"topic_mistakes"`
}

func (req *LessonAnalysis) CalculateNet() {
	req.Net = float64(req.Correct) - (float64(req.Wrong) * 0.25)
}
