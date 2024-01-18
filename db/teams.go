package db

import (
	"elite-backend/internal/models"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (l *LeagueDB) TeamExist(name string) bool {
	var Team *models.Team
	teamName := strings.TrimSpace(strings.ToLower(name))
	filter := bson.M{"name": teamName}
	err := TeamsCollection.FindOne(l.MongoCtx, filter).Decode(&Team)
	return err == nil
}

func (l *LeagueDB) IsTeamFound(req string) (*models.Team, bool) {
	var itemTeam *models.Team
	requestItem := strings.TrimSpace(req)

	id, _ := primitive.ObjectIDFromHex(requestItem)
	filter := bson.M{"_id": id}
	err := TeamsCollection.FindOne(l.MongoCtx, filter).Decode(&itemTeam)
	if err != nil {
		return nil, false
	}
	return itemTeam, true
}
