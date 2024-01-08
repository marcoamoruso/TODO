package main

import (
	"TODO/Config"
	"TODO/Models"
	"TODO/Routes"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// Initialize Echo instance
	e := echo.New()

	// Connect to the database
	Config.DatabaseInit()
	defer Config.GetDB().DB()

	// Perform migrations using AutoMigrate
	db := Config.GetDB()
	err := db.AutoMigrate(&Models.TODO{})
	if err != nil {
		panic(err)
	}

	// Set up Routes
	Routes.SetupRoutes(e)

	// Start the server
	serverPort := os.Getenv("SERVER_PORT")
	e.Logger.Fatal(e.Start(":" + serverPort))
}
