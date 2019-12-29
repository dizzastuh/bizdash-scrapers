package main

import (
    "fmt"

    "github.com/dizzastuh/bizdash-scrapers/internal/claws"
)

func main() {

    // Scrape all the sources
    fmt.Println("Scraping Freddy")
    claws.ScrapeFreddy()
    
    // fmt.Println("Scraping Census")
    // claws.ScrapeCensus()

}
