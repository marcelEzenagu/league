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

// Player
// list
func (s *Server) ListPlayers(c *fiber.Ctx) error {
	// team := new(models.Team)
	filter := bson.M{}
	cur, err := db.PlayersCollection.Find(c.Context(), filter)
	if err != nil {
		fmt.Println("error listing players", err)
		return c.Status(500).SendString(fmt.Sprintf("Internal error: %v\n", err))
	}
	players := []*models.Player{}
	if err = cur.All(context.TODO(), &players); err != nil {
		return c.Status(500).SendString(fmt.Sprintf("Internal error: %v\n", err))
	}
	fmt.Printf("playersList: %v", players[1])

	return c.Render("player.list", fiber.Map{
		"PageTitle": "Players List",
		"Players":   players,
	})
	// return c.Render("player.form", fiber.Map{})
}

func (s *Server) NewPlayers(c *fiber.Ctx) error {
	fmt.Println("CALLED UPLOAD REST /")
	// team := new(models.Team)

	return c.Render("player.form", fiber.Map{})
}
func (s *Server) AddPlayer(c *fiber.Ctx) error {
	// player := new(models.Player)
	// if err := c.BodyParser(player); err != nil {
	// 	fmt.Println("error = ", err)
	// 	return c.SendStatus(200)
	// }
	fname := strings.ToLower(c.FormValue("fName"))
	lname := strings.ToLower(c.FormValue("lName"))
	email := c.FormValue("email")
	phone := strings.ToLower(c.FormValue("phone"))
	position := strings.ToLower(c.FormValue("position"))
	team := strings.ToLower(c.FormValue("team"))

	if fname == "" || lname == "" || team == "" || email == "" || phone == "" {
		return c.Status(400).SendString("fName,lName,team,email,phone is required but empty.")
	}
	teamExist := s.LeagueDB.TeamExist(team)
	if !teamExist {
		return c.Status(400).SendString("a team with this name does not exist.")
	}
	playerExists := s.LeagueDB.PlayerExist(fname, lname, phone, email)
	if playerExists {
		return c.Status(400).SendString("a player with this details already exists in the team.")
	}

	player := &models.Player{
		FirstName: fname,
		LastName:  lname,
		Email:     email,
		Phone:     phone,
		Position:  position,
		Team:      team,
		CreatedAt: time.Now().String(),
	}

	_, err := db.PlayersCollection.InsertOne(c.Context(), player)
	if err != nil {
		fmt.Println("error creating item", err)
		return c.Status(500).SendString(fmt.Sprintf("Internal error: %v\n", err))
	}
	// s.LeagueDB
	return c.Status(201).SendString("player added successfully")
}
func (s *Server) EditPlayer(c *fiber.Ctx) error {

	return nil
}

func (s *Server) DeletePlayer(c *fiber.Ctx) error {

	playerID := strings.TrimSpace(c.Params("id"))

	requestData, _ := primitive.ObjectIDFromHex(playerID)
	filter := bson.M{"_id": requestData}

	_, isFound := s.LeagueDB.IsPlayerFound(playerID)
	if !isFound {
		return c.Status(400).SendString("a player with this id not found.")
	}

	_, err := db.PlayersCollection.DeleteOne(c.Context(), filter)
	if err != nil {
		fmt.Println("error deleting players", err)
		return c.Status(500).SendString(fmt.Sprintf("Internal error: %v\n", err))
	}
	return c.Status(201).SendString("player deleted successfully")
}
func (s *Server) ViewPlayer(c *fiber.Ctx) error {
	playerID := strings.TrimSpace(c.Params("id"))
	player, isFound := s.LeagueDB.IsPlayerFound(playerID)
	if !isFound {
		return c.Status(400).SendString("a player with this id not found.")
	}

	return c.Status(200).JSON(player)
}
