package Models

import (
    // "go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID     	 int `json:"_id,omitempty" bson:"_id,omitempty"`
	Quantity   int           `json:"quantity" bson:"quantity,omitempty"`
	Title  string             `json:"title" bson:"title,omitempty"`
	Author string          `json:"author" bson:"author,omitempty"`
}
