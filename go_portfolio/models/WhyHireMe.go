package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type WhyHireMe struct {
	ID          bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Icon        string        `json:"icon,omitempty" bson:"icon,omitempty"`
	Title       string        `json:"title,omitempty" bson:"title,omitempty"`
	Description string        `json:"description,omitempty" bson:"description,omitempty"`
	Color       string        `json:"color,omitempty" bson:"color,omitempty"`
}
