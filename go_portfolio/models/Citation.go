package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Citation struct {
	ID       bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Citation string        `json:"citation,omitempty" bson:"citation,omitempty"`
	Author   string        `json:"author,omitempty" bson:"author,omitempty"`
}
