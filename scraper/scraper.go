package scraper

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"web-scraper-go/types"
)

func StartScraping(
	urlToVisit string,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scrapping on go routines on every %s duration", timeBetweenRequest)

	// Implement a wait group
	wg := &sync.WaitGroup{}

	// Similar to interval
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		wg.Add(1)
		go scrape(urlToVisit, wg)
		wg.Wait()
	}
}

func scrape(urlToVisit string, wg *sync.WaitGroup) {
	defer wg.Done()
	// initializing a chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	// navigate to the target web page and select the HTML elements of interest
	log.Printf("Scraping URL: %s", urlToVisit)

	var matchListNodes []*cdp.Node
	err := chromedp.Run(ctx,
		chromedp.Navigate(urlToVisit),
		chromedp.Nodes(".match-fixture", &matchListNodes, chromedp.ByQueryAll, chromedp.NodeVisible),
	)
	if err != nil {
		log.Fatalf("Fail to run chromedp, %v", err)
	}

	var matches []types.Match
	var match types.Match

	for _, node := range matchListNodes {
		match.Id = node.AttributeValue("data-comp-match-item")
		match.HomeTeam = node.AttributeValue("data-home")
		match.AwayTeam = node.AttributeValue("data-away")
		match.Stadium = node.AttributeValue("data-venue")
		match.Status = node.AttributeValue("data-comp-match-item-status")
		match.DateTime = node.AttributeValue("data-comp-match-item-ko")

		if err != nil {
			log.Fatalf("Fail to perform extract from node: %v with error: %v", node, err)
		}
		matches = append(matches, match)
	}

	log.Printf("matches %v", matches)
}
