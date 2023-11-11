// Package main
// https://www.zenrows.com/blog/web-scraping-golang#visit-target-html-page
package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	log.Println("Welcome to Hao Go Lang Web Scraper")

	urlToVisit := "https://scrapeme.live/shop/"
	timeBetweenRequest := time.Second * 5

	// Implement a wait group
	wg := &sync.WaitGroup{}

	wg.Add(1)

	go startScraping(urlToVisit, timeBetweenRequest, wg)

	wg.Wait()
}
