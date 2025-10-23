package analysis

import "go.mongodb.org/mongo-driver/v2/bson"

type Lesson struct {
	Id             bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Name           string        `bson:"name" json:"name"`
	TotalQuestions int           `bson:"totalQuestions" json:"totalQuestions"`
}
