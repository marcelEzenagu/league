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

// teams
// list
func (s *Server) ListTeams(c *fiber.Ctx) error {
	// team := new(models.Team)
	filter := bson.M{}

	cur, err := db.TeamsCollection.Find(c.Context(), filter)
	if err != nil {
		fmt.Println("error listing teams", err)
		return c.Status(500).SendString(fmt.Sprintf("Internal error: %v\n", err))
	}
	teams := []*models.Team{}
	if err = cur.All(context.TODO(), &teams); err != nil {
		return c.Status(500).SendString(fmt.Sprintf("Internal error: %v\n", err))
	}

	fmt.Printf("TEAMS-LIST: %s", teams)

	return c.Render("teamlist", fiber.Map{
		"PageTitle": "Team List",
		"Teams":     teams,
	})
}

func (s *Server) NewTeam(c *fiber.Ctx) error {
	return c.Render("team.form", fiber.Map{})
}

func (s *Server) EditTeams(c *fiber.Ctx) error {

	return nil
}

func (s *Server) AddTeams(c *fiber.Ctx) error {

	logo, err := c.FormFile("logo")
	if err != nil {
		fmt.Println("error = ", err)
		return c.SendStatus(400)
	}
	name := strings.ToLower(c.FormValue("name"))

	fmt.Println("LGOG:: ", logo.Filename)

	if name == "" {
		return c.Status(400).SendString("name is required but empty.")
	}
	teamExists := s.LeagueDB.TeamExist(name)

	if teamExists {
		return c.Status(400).SendString("a team with this name already exists.")
	}

	team := &models.Team{
		Name: name,
		// Logo: logo,
		CreatedAt: time.Now().String(),
	}
	_, err = db.TeamsCollection.InsertOne(c.Context(), team)
	if err != nil {
		fmt.Println("error creating item", err)
		return c.Status(500).SendString(fmt.Sprintf("Internal error: %v\n", err))
	}
	// s.LeagueDB
	return c.Status(201).SendString("team added successfully")
}

func (s *Server) DeleteTeams(c *fiber.Ctx) error {
	teamID := strings.TrimSpace(c.Params("id"))

	requestData, _ := primitive.ObjectIDFromHex(teamID)
	filter := bson.M{"_id": requestData}
	_, isFound := s.LeagueDB.IsTeamFound(teamID)
	if !isFound {
		return c.Status(400).SendString("a team with this id not found.")
	}

	res, err := db.TeamsCollection.DeleteOne(c.Context(), filter)
	if err != nil {
		fmt.Println("error deleting teams", err)
		return c.Status(500).SendString(fmt.Sprintf("Internal error: %v\n", err))
	}

	fmt.Println("result:: ", res)
	return c.Status(201).SendString("team deleted successfully")

}

func (s *Server) ViewTeams(c *fiber.Ctx) error {
	teamID := strings.TrimSpace(c.Params("id"))
	team, isFound := s.LeagueDB.IsTeamFound(teamID)
	if !isFound {
		return c.Status(400).SendString("a team with this id not found.")
	}

	return c.Status(200).JSON(team)
}
