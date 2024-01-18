package models

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Team struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name      string             `json:"name" bson:"name"`
	Logo      string             `json:"logo" bson:"logo"`
	Position  string             `json:"position" bson:"position"`
	Played    string             `json:"played" bson:"played"`
	Won       string             `json:"won" bson:"won"`
	Lost      string             `json:"lost" bson:"lost"`
	Drawn     string             `json:"drawn" bson:"drawn"`
	Players   []Player           `json:"players,omitempty"`
	CreatedAt string             `json:"createdAt" bson:"createdAt"`
	UpdatedAt string             `json:"updatedAt" bson:"updatedAt"`
	DeletedAt string             `json:"deletedAt" bson:"deletedAt"`
}

func (ms *Team) ToLower() {
	ms.Name = strings.ToLower(ms.Name)
}

type Player struct {
	Id             primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	FirstName      string             `json:"firstName" bson:"firstName"`
	LastName       string             `json:"lastName" bson:"lastName"`
	ProfilePicture string             `json:"profilePicture" bson:"profilePicture"`
	Email          string             `json:"email" bson:"email"`
	Position       string             `json:"position"  bson:"position"`
	Team           string             `json:"team"  bson:"team"`
	TeamID         string             `json:"teamID"  bson:"teamID"`
	Phone          string             `json:"phone"  bson:"phone"`
	CreatedAt      string             `json:"createdAt" bson:"createdAt"`
	UpdatedAt      string             `json:"updatedAt" bson:"updatedAt"`
	DeletedAt      string             `json:"deletedAt" bson:"deletedAt"`
}

func (ms *Player) ToLower() {
	ms.FirstName = strings.ToLower(ms.FirstName)
	ms.LastName = strings.ToLower(ms.LastName)
	ms.Email = strings.ToLower(ms.Email)
	ms.Position = strings.ToLower(ms.Position)
}
