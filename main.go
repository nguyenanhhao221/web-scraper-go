// Package main
package main

import (
	"log"
	"time"

	"web-scraper-go/db"
	"web-scraper-go/db/database"
	"web-scraper-go/handler"
	"web-scraper-go/scraper"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
)

// type apiConfig struct {
// 	DB *database.Queries
// }

func main() {
	sqlConnection, err := db.Init()
	if err != nil {
		log.Fatalf("Failed to ping the database, did you forget to run Docker? Error: %v", err)
	}

	// database is the package generate by sqlc which contain our queries to the actual database.
	// So basically here The queries object is responsible for executing SQL queries against the database using the underlying connection.
	queries := database.New(sqlConnection)

	// This will get pass to our handler so that the handler have access to the database
	// apiCfg := apiConfig{
	// 	DB: queries,
	// }

	log.Println("Welcome to Hao's Go Lang Web Scraper")

	app := fiber.New(fiber.Config{
		AppName: "Hao's Go Lang Web Scraper",
	})
	app.Use(cors.New())

	// Create a new route group '/api'
	api := app.Group("/api", func(c *fiber.Ctx) error {
		return c.Next()
	})

	// Create a new route for API v1
	v1Router := api.Group("/v1")

	v1Router.Get("/healthz", handler.HealthCheck)

	urlToVisit := "https://www.premierleague.com/fixtures"
	timeBetweenRequest := time.Second * 5

	go scraper.StartScraping(queries, urlToVisit, timeBetweenRequest)

	log.Fatal(app.Listen(":3000"))
}
