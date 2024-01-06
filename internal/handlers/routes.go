package handlers

import (
	// "recipes-fiber-basic-api/handlers"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Routes(server *Server) {
	app := server.App
	// method GET
	// app.Get("/", handlers.ReadData)
	// // Method GET by ID

	app.Post("/teams", func(c *fiber.Ctx) error {

		fmt.Println("CALLED UPLOAD REST /")
		// team := new(models.Team)

		// return c.Status(200).JSON()
		return nil
	})

	// app.Get("/:id", handlers.ReadDataById)
	// // Method POST
	// app.Post("/", handlers.InsertData)
	// // method DELETE
	// app.Delete("/:id", handlers.DeleteData)
	// // Method PATCH
	// app.Patch("/:id", handlers.PatchData)
}
