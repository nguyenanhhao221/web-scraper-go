package main

import (
	"log"
	"sync"
	"time"

	"github.com/gocolly/colly"
)

func startScraping(
	urlToVisit string,
	timeBetweenRequest time.Duration,
	wg *sync.WaitGroup,
) {
	log.Printf("Scrapping on go routines on every %s duration", timeBetweenRequest)
	// Similar to interval
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		go scrape(urlToVisit, wg)
	}
}

func scrape(urlToVisit string, wg *sync.WaitGroup) {
	defer wg.Done()
	// First, Colly's main entity is the Collector. A Collector allows you to perform HTTP requests. Also, it gives you access to the web scraping callbacks offered by the Colly interface.
	c := colly.NewCollector()

	// defining a data structure to store the scraped data
	type PokemonProduct struct {
		url, image, name, price string
	}
	// initializing the slice of structs that will contain the scraped data
	var pokemonProducts []PokemonProduct

	// iterating over the list of HTML product elements
	c.OnHTML("li.product", func(e *colly.HTMLElement) {
		// initializing a new PokemonProduct instance
		pokemonProduct := PokemonProduct{}

		// scraping the data of interest
		pokemonProduct.url = e.ChildAttr("a", "href")
		pokemonProduct.image = e.ChildAttr("img", "src")
		pokemonProduct.name = e.ChildText("h2")
		pokemonProduct.price = e.ChildText(".price")

		// adding the product instance with scraped data to the list of products
		pokemonProducts = append(pokemonProducts, pokemonProduct)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})
	err := c.Visit(urlToVisit)
	if err != nil {
		log.Fatalf("Error: failed to call website %v, with error: %v", urlToVisit, err)
	}

	log.Printf("Pokemons: %v", pokemonProducts)
}
