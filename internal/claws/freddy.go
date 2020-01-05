package claws

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "strings"
    
    "github.com/dizzastuh/bizdash-scrapers/internal/claws/fred"
    // "github.com/dizzastuh/bizdash-scrapers/internal/claws/utils"
    . "github.com/nswekosk/fred_go_toolkit"
)

func ScrapeFreddy() {
    client, err := getClient()

    if err == nil {
        consumeData(client)
    } else {
        fmt.Println("Client could not be acquired")
        fmt.Println(err)
        log.Fatal(err)
    }
}

func consumeData(client *FredClient) {
    fred.ConsumeAllSeries(client)

    // TODO: import the rest
}

func getClient() (*FredClient, error) {
    logpath, _ := filepath.Abs("./log/fred.log")

    apiKey := os.Getenv("FRED_API_KEY")
    trimmed := strings.Trim(apiKey, "\r\n")

    fredConfig := FredConfig{ APIKey: trimmed, FileType: FileTypeJSON, LogFile: logpath, }
    return CreateFredClient(fredConfig)
}
