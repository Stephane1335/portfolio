package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Citation struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Citation string             `json:"citation,omitempty" bson:"citation,omitempty"`
	Author   string             `json:"author,omitempty" bson:"author,omitempty"`
}
