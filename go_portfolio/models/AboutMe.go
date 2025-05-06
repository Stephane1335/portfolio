package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type AbouMe struct {
	ID          bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Quote       string        `json:"quote,omitempty" bson:"quote,omitempty"`
	Description string        `json:"description,omitempty" bson:"description,omitempty"`
	Position    string        `json:"position,omitempty" bson:"position,omitempty"`
	Name        string        `json:"name,omitempty" bson:"name,omitempty"`
	Signature   string        `json:"signature,omitempty" bson:"signature,omitempty"`
}
