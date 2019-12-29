package claws

import (
    "log"
    "fmt"
    "os"
    "path/filepath"
    "strings"
    
    model "github.com/dizzastuh/bizdash-db/model"
    "github.com/dizzastuh/bizdash-scrapers/internal/claws/utils"
    . "github.com/nswekosk/fred_go_toolkit"
)

func ScrapeFreddy() {
    client, err := getClient()

    if err == nil {
        consumeData(client)
    } else {
        fmt.Println("Client could not be acquired")
        fmt.Println(err)
    }
}

func consumeData(client *FredClient) {
    consumeAllSeries(client)

    // TODO: import the rest of the data from Fred
}

func consumeAllSeries(client *FredClient) {    
    seriespath, _ := filepath.Abs("../res/fred-series.txt")

    seriesList, err := utils.ToStringArray(seriespath)
    
    if err != nil {
        fmt.Println(err)
        log.Fatal(err)
    }

    for i:= 0; i < len(seriesList); i++ {
        consumeSeries(seriesList[i], client)
    }
}

func consumeSeries(series string, client *FredClient) {
    fmt.Printf("Consuming %s\n", series)
    params := make(map[string]interface{})
    params["series_id"] = series

    srs, err := client.GetSeriesObservations(params)

    if err != nil {
        fmt.Printf("Error retrieving series %s\n", series)
        fmt.Println(err)
    }

    model.InsertFredObs(srs, series)
}

func getClient() (*FredClient, error) {
    logpath, _ := filepath.Abs("./log/fred.log")

    apiKey := os.Getenv("FRED_API_KEY")
    trimmed := strings.Trim(apiKey, "\r\n")

    fredConfig := FredConfig{ APIKey: trimmed, FileType: FileTypeJSON, LogFile: logpath, }
    return CreateFredClient(fredConfig)
}
