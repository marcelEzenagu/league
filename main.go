package main

import (
	"fmt"
	"sync"

	"elite-backend/db"
	"elite-backend/internal/handlers"

	"github.com/joho/godotenv"
)

var (
	mutex     sync.RWMutex
	queueName string
)

type Server struct {
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("cannot load up the env file : %v \n", err)
	}

	// connect DB
	client, err := db.OpenDbConnection()
	if err != nil {
		fmt.Printf("error connecting to DB: %v \n", err)
	}
	// start server
	handlers.Start(client)
}
