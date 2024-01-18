package db

import (
	"elite-backend/internal/models"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (l *LeagueDB) FixtureExist(homeTeam, awayTeam string) bool {
	var fixture *models.Fixtures
	filter := bson.M{"homeTeam": homeTeam, "awayTeam": awayTeam}

	err := FixturesCollection.FindOne(l.MongoCtx, filter).Decode(&fixture)
	// log.Panic("filter", filter, err, "fixture", fixture)
	return err == nil
}

func (l *LeagueDB) IsFixtureFound(req string) (*models.Fixtures, bool) {
	var itemFixture *models.Fixtures
	requestItem := strings.TrimSpace(req)

	id, _ := primitive.ObjectIDFromHex(requestItem)
	filter := bson.M{"_id": id}
	err := FixturesCollection.FindOne(l.MongoCtx, filter).Decode(&itemFixture)
	if err != nil {
		return nil, false
	}
	return itemFixture, true
}
