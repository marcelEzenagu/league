package handlers

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"

	"elite-backend/db"
)

type Server struct {
	App      *fiber.App
	LeagueDB *db.LeagueDB
}

func Start(db *db.LeagueDB) error {

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// handling cors
	app.Use(cors.New(cors.Config{AllowOrigins: "*"}))

	// app.Post("/teams", func(c *fiber.Ctx) error {

	// 	fmt.Println("CALLED UPLOAD REST /")
	// 	team := new(models.Team)

	// 	return c.Status(200).JSON()
	// })

	server := &Server{
		App:      app,
		LeagueDB: db,
	}

	Routes(server)
	port := fmt.Sprintf(":%v", os.Getenv("PORT"))

	fmt.Println("PORT FOR Server:: ", port)
	err := app.Listen(port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return err
	}
	return nil

}
