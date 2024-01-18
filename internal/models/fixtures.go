package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Fixtures struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	HomeTeam  string             `json:"homeTeam" bson:"homeTeam"`
	AwayTeam  string             `json:"awayTeam" bson:"awayTeam"`
	CreatedAt string             `json:"createdAt" bson:"createdAt"`
	UpdatedAt string             `json:"updatedAt" bson:"updatedAt"`
	DeletedAt string             `json:"deletedAt" bson:"deletedAt"`

	MatchTime string `json:"matchTime" bson:"matchTime"`
}
