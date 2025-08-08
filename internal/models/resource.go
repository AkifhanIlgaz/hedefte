package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Resource struct {
	Id      bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Uid     string        `bson:"uid" json:"uid" validate:"required"`
	Name    string        `bson:"name" json:"name" validate:"required,min=1,max=100"`
	Subject string        `bson:"subject" json:"subject" validate:"required,min=1,max=50"`
}

type AddResourceRequest struct {
	ExamType string `json:"examType" bson:"-" validate:"required,oneof=TYT AYT"`
	Name     string `json:"name" bson:"name" validate:"required,min=1,max=100"`
	Subject  string `json:"subject" bson:"subject" validate:"required,min=1,max=50"`
}
