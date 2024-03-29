package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type News struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Title     string             `json:"title" bson:"title"`
	Body      string             `json:"body" bson:"body"`
	Position  string             `json:"position"  bson:"position"`
	CreatedAt string             `json:"createdAt" bson:"createdAt"`
	UpdatedAt string             `json:"updatedAt" bson:"updatedAt"`
	DeletedAt string             `json:"deletedAt" bson:"deletedAt"`
}
