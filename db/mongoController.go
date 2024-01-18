package db

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// var uri = os.Getenv("TRAVAS_DB_URI")

type LeagueDB struct {
	DB       *mongo.Client
	MongoCtx context.Context
}

var (
	TeamsCollection    *mongo.Collection
	PlayersCollection  *mongo.Collection
	NewsCollection     *mongo.Collection
	FixturesCollection *mongo.Collection
)

func SetConnection(uri string) (*mongo.Client, error) {
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)

	ctx, cancelCtx := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancelCtx()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPIOptions))
	if err != nil {
		log.Panicln(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalln(err)
	}

	return client, nil
}

func OpenDbConnection() (*LeagueDB, error) {
	uri := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")

	count := 0
	for {
		client, err := SetConnection(uri)

		if err != nil {
			log.Println("MongoDB not yet ready to connect ...")
			count++
			return nil, err
		} else {

			dbConn := &LeagueDB{
				DB: client,
			}
			log.Println("Connected to MongoDB ...")
			teamsCollection := os.Getenv("TEAMS_COLLECTION")
			playersCollection := os.Getenv("PLAYERS_COLLECTION")
			fixturesCollection := os.Getenv("FIXTURES_COLLECTION")

			TeamsCollection = client.Database(dbName).Collection(teamsCollection)
			PlayersCollection = client.Database(dbName).Collection(playersCollection)
			FixturesCollection = client.Database(dbName).Collection(fixturesCollection)

			return dbConn, nil
		}

		log.Println("Trying to reconnect MongoDB database ...")
		time.Sleep(5 * time.Second)
		continue
	}
}
