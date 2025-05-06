package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Timeline struct {
	ID             bson.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Type           string        `json:"type,omitempty" bson:"type,omitempty"`
	Degree         string        `json:"degree,omitempty" bson:"degree,omitempty"`
	Classification string        `json:"classification,omitempty" bson:"classification,omitempty"`
	Institution    string        `json:"institution,omitempty" bson:"institution,omitempty"`
	StartYear      int           `json:"startyear,omitempty" bson:"startyear,omitempty"`
	EndYear        int           `json:"endyear,omitempty" bson:"endyear,omitempty"`
	Description    string        `json:"description,omitempty" bson:"description,omitempty"`
}
