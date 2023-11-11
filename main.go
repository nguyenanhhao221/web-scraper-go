// Package main
// https://www.zenrows.com/blog/web-scraping-golang#visit-target-html-page
package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	log.Println("Welcome to Hao Go Lang Web Scraper")

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	urlToVisit := "https://scrapeme.live/shop/"
	timeBetweenRequest := time.Second * 5

	go startScraping(urlToVisit, timeBetweenRequest)

	err := app.Listen(":3000")
	if err != nil {
		log.Fatalf("Fail to start fiber server, %v", err)
	}
}
