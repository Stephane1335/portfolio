package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Skill struct {
	ID     bson.ObjectID      `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title,omitempty" bson:"title,omitempty"`
	Techno []bson.ArrayWriter `json:"techno,omitempty" bson:"techno,omitempty"`
	Level  int                `json:"int,omitempty" bson:"int,omitempty"`
}
