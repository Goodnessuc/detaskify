package main

import (
	"detaskify/internal/db"
	"detaskify/internal/transport/http"
	"detaskify/internal/users"
	"fmt"
	"log"
)

func Run() error {
	fmt.Println("starting up the application...")
	database, err := db.NewDatabase()
	err = database.MigrateDatabase()
	if err != nil {
		log.Printf("Error Migrating to the database %v", err)
		return err
	}
	userService := users.NewUserService(database)

	followerService := users.NewFollowerService(database)
	handler := http.NewHandler(userService, followerService, logger)
	if err := handler.Serve(); err != nil {
		log.Println("failed to gracefully serve our application")
		return err
	}

	return nil

}

func main() {
	if err := Run(); err != nil {
		fmt.Println("error starting up the application")
	}
}
