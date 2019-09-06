package main

import (
    "log"
    "fmt"

    "github.com/joho/godotenv"
    "github.com/dizzastuh/bizdash-scrapers/internal/claws"
)

// step one -- load environment variables
func init() {
    if err := godotenv.Load(); err != nil {
        log.Print("Dude where's your env?")
    }
}

func main() {

    // Scrape all the sources
    fmt.Println("Scraping Freddy")
    claws.ScrapeFreddy()
    
    // fmt.Println("Scraping Census")
    // claws.ScrapeCensus()

}
