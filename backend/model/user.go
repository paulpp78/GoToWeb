package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
	Age  string             `json:"age" bson:"age"`
}
