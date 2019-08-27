// Runs all the scrapers
package main

import (
    "io"
    "log"
    "net/http"
    "os"

    "./claws"
)

// step one -- load environment variables
func init() {
    if err := godotenv.Load(); err != nil {
        log.Print("Where's your env?")
    }
}

func main() {

    // Scrape all the sources
    freddy.ScrapeFreddy()
    census.ScrapeCensus()

}
