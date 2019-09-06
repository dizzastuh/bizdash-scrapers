package claws

import (
    "path/filepath"
    "fmt"
    "os"

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
    }

    for i:= 0; i < len(seriesList); i++ {
    // for i:= 0; i < 1; i++ {
        consumeSeries(seriesList[i], client)
    }
}

func consumeSeries(series string, client *FredClient) {
    fmt.Printf("Consuming %s\n", series)
    params := make(map[string]interface{})
    params["series_id"] = series

    fmt.Println(params)

    srs, err := client.GetSeriesObservations(params)

    if err != nil {
        fmt.Printf("Error retrieving series %s\n", series)
        fmt.Println(err)
    }
    
    for i:= 0; i < len(srs.Observations); i++ {
        obs := srs.Observations[i]
        fmt.Println(obs.Value)
    }
}

func getClient() (*FredClient, error) {
    logpath, _ := filepath.Abs("../log/fred.log")
    fredConfig := FredConfig{ APIKey: os.Getenv("FRED_API_KEY"), FileType: FileTypeJSON, LogFile: logpath, }
    return CreateFredClient(fredConfig)
}
