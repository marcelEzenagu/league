package handlers

func Routes(s *Server) {
	app := s.App
	// method GET
	// app.Get("/", handlers.ReadData)
	// // Method GET by ID
	// teams
	app.Get("/teams", s.ListTeams)
	app.Get("/teams/new", s.NewTeam)
	app.Post("/teams", s.AddTeams)
	app.Get("/teams/:id", s.ViewTeams)
	app.Delete("/teams/:id", s.DeleteTeams)

	// players
	app.Post("/players", s.AddPlayer)
	app.Get("/players", s.ListPlayers)
	app.Get("/players/:id", s.ViewPlayer)
	app.Delete("/players/:id", s.DeletePlayer)

	// routes
	app.Post("/fixtures", s.AddFixtures)
	app.Get("/fixtures", s.ListFixtures)
	app.Delete("/fixtures/:id", s.DeleteFixtures)
	app.Get("/fixtures/:id", s.ViewFixtures)
	// app.Get("/:id", handlers.ReadDataById)
	// // Method POST
	// app.Post("/", handlers.InsertData)
	// // method DELETE
	// app.Delete("/:id", handlers.DeleteData)
	// // Method PATCH
	// app.Patch("/:id", handlers.PatchData)
}
