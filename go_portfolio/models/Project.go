package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Project struct {
	ID          bson.ObjectID      `json:"id,omitempty" bson:"_id,omitempty"`
	Categories  []bson.ArrayWriter `json:"categories,omitempty" bson:"categories,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	ImageUrl    string             `json:"imageurl,omitempty" bson:"imageurl,omitempty"`
	Link        string             `json:"link,omitempty" bson:"link,omitempty"`
}
