package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Travel struct
type Travel struct {
	ObjectID primitive.ObjectID `json:"id" bson:"_id`
	Name     string             `json:"name" bson:"name"`
	Contact  string             `json:"contact" bson:"contact"`
}
