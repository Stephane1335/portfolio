package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Skill struct {
	ID          bson.ObjectID      `json:"id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Techno      []bson.ArrayWriter `json:"techno,omitempty" bson:"techno,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Level       string             `json:"int,omitempty" bson:"int,omitempty"`
	Color       string             `json:"color,omitempty" bson:"color,omitempty"`
}
