package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Title       string             `json:"title,omitempty" bson:"title"`
	Description string             `json:"description,omitempty" bson:"description"`
	Completed   bool               `json:"completed,omitempty" bson:"completed"`
}
