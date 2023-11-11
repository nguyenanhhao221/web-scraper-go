package scraper

import (
	"context"
	"log"
	"sync"
	"time"

	"web-scraper-go/db/database"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

func StartScraping(
	queries *database.Queries,
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
		go scrape(queries, urlToVisit, wg)
		wg.Wait()
	}
}

func scrape(
	queries *database.Queries,
	urlToVisit string, wg *sync.WaitGroup,
) {
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

	for _, node := range matchListNodes {

		_, err := queries.CreateMatch(context.Background(), database.CreateMatchParams{
			ID:       node.AttributeValue("data-comp-match-item"),
			Hometeam: node.AttributeValue("data-home"),
			Awayteam: node.AttributeValue("data-away"),
			Stadium:  node.AttributeValue("data-venue"),
			Status:   node.AttributeValue("data-comp-match-item-status"),
			Datetime: node.AttributeValue("data-comp-match-item-ko"),
		})
		if err != nil {
			log.Fatalf("Fail to perform extract from node: %v with error: %v", node, err)
		}
	}
}
