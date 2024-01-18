package db

import (
	"elite-backend/internal/models"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (l *LeagueDB) PlayerExist(fname, lname, phone, email string) bool {
	var player *models.Player
	filter := bson.M{"firstName": fname, "lastName": lname, "phone": phone, "email": email}

	err := PlayersCollection.FindOne(l.MongoCtx, filter).Decode(&player)
	return err == nil
}

func (l *LeagueDB) IsPlayerFound(req string) (*models.Player, bool) {
	var itemPlayer *models.Player
	requestItem := strings.TrimSpace(req)

	id, _ := primitive.ObjectIDFromHex(requestItem)
	filter := bson.M{"_id": id}
	err := PlayersCollection.FindOne(l.MongoCtx, filter).Decode(&itemPlayer)
	if err != nil {
		return nil, false
	}
	return itemPlayer, true
}
