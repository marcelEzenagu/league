package handlers

import (
	"context"
	"elite-backend/db"
	"elite-backend/internal/models"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Fixtures
// list
func (s *Server) ListFixtures(c *fiber.Ctx) error {
	// fixture := new(models.fixture)
	filter := bson.M{}

	cur, err := db.FixturesCollection.Find(c.Context(), filter)
	if err != nil {
		fmt.Println("error listing Fixtures", err)
		return c.Status(500).SendString(fmt.Sprintf("Internal error: %v\n", err))
	}
	Fixtures := []*models.Fixtures{}
	if err = cur.All(context.TODO(), &Fixtures); err != nil {
		return c.Status(500).SendString(fmt.Sprintf("Internal error: %v\n", err))
	}
	fmt.Printf("Fixtures-LIST: %s", Fixtures[0])

	return c.Render("fixture.list", fiber.Map{
		"PageTitle": "fixture List",
		"Fixtures":  Fixtures,
	})
}

// TODO
func (s *Server) ListFixturesByDate(c *fiber.Ctx) error {
	// fixture := new(models.fixture)
	filter := bson.M{}

	cur, err := db.FixturesCollection.Find(c.Context(), filter)
	if err != nil {
		fmt.Println("error listing Fixtures", err)
		return c.Status(500).SendString(fmt.Sprintf("Internal error: %v\n", err))
	}
	Fixtures := []*models.Fixtures{}
	if err = cur.All(context.TODO(), &Fixtures); err != nil {
		return c.Status(500).SendString(fmt.Sprintf("Internal error: %v\n", err))
	}
	fmt.Printf("Fixtures-LIST: %s", Fixtures)

	return c.Render("fixturelist", fiber.Map{
		"PageTitle": "fixture List",
		"Fixtures":  Fixtures,
	})
}

func (s *Server) NewFixture(c *fiber.Ctx) error {
	return c.Render("fixture.form", fiber.Map{})
}

func (s *Server) EditFixtures(c *fiber.Ctx) error {

	return nil
}

func (s *Server) AddFixtures(c *fiber.Ctx) error {

	home := strings.ToLower(c.FormValue("homeTeam"))
	away := strings.ToLower(c.FormValue("awayTeam"))
	mTime := (c.FormValue("matchTime"))

	if home == "" || away == "" || mTime == "" {
		return c.Status(400).SendString("matchTime, homeTeam, awayTeam required but empty.")
	}
	fixtureExists := s.LeagueDB.FixtureExist(home, away)
	if fixtureExists {
		return c.Status(400).SendString("a fixture with this name already exists.")
	}

	fixture := &models.Fixtures{
		HomeTeam:  home,
		AwayTeam:  away,
		MatchTime: mTime,
		CreatedAt: time.Now().String(),
	}
	_, err := db.FixturesCollection.InsertOne(c.Context(), fixture)
	if err != nil {
		fmt.Println("error creating fixtures", err)
		return c.Status(500).SendString(fmt.Sprintf("Internal error: %v\n", err))
	}
	return c.Status(201).SendString("fixture added successfully")
}

func (s *Server) DeleteFixtures(c *fiber.Ctx) error {
	fixtureID := strings.TrimSpace(c.Params("id"))

	requestData, _ := primitive.ObjectIDFromHex(fixtureID)
	filter := bson.M{"_id": requestData}

	_, isFound := s.LeagueDB.IsFixtureFound(fixtureID)
	if !isFound {
		return c.Status(400).SendString("a fixture with this id not found.")
	}

	_, err := db.FixturesCollection.DeleteOne(c.Context(), filter)
	if err != nil {
		fmt.Println("error deleting Fixtures", err)
		return c.Status(500).SendString(fmt.Sprintf("Internal error: %v\n", err))
	}
	return c.Status(201).SendString("fixture deleted successfully")
}

func (s *Server) ViewFixtures(c *fiber.Ctx) error {

	fixtureID := strings.TrimSpace(c.Params("id"))
	fixture, isFound := s.LeagueDB.IsFixtureFound(fixtureID)
	if !isFound {
		return c.Status(400).SendString("a fixture with this id not found.")
	}

	return c.Status(200).JSON(fixture)
}
