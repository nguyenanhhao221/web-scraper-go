// Package main
// https://www.zenrows.com/blog/web-scraping-golang#visit-target-html-page
package main

import (
	"log"
	"time"

	"web-scraper/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
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

	urlToVisit := "https://scrapeme.live/shop/"
	timeBetweenRequest := time.Minute * 1

	go startScraping(urlToVisit, timeBetweenRequest)

	log.Fatal(app.Listen(":3000"))
}
