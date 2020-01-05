package fred

import (
    "encoding/json"
    "io/ioutil"
    "fmt"
    "path/filepath"
    
    model "github.com/dizzastuh/bizdash-db/model"
    utils "github.com/dizzastuh/bizdash-scrapers/internal/claws/utils"
    . "github.com/nswekosk/fred_go_toolkit"
)

type FSeriesList struct {
    List []FSeries `json:"series"`
}

type FSeries struct {
    Name        string `json:"name"`
    Description string `json:"description"`
}

func ConsumeAllSeries(client *FredClient) {    
    seriesPath, _ := filepath.Abs("./res/fred-series.json")

    seriesFile := utils.OpenFile(seriesPath)
    defer seriesFile.Close()

    bytes, _ := ioutil.ReadAll(seriesFile)

    var series FSeriesList
    json.Unmarshal(bytes, &series)

    for i:= 0; i < len(series.List); i++ {
        consumeSeries(&series.List[i], client)
    }
}

func consumeSeries(series *FSeries, client *FredClient) {
    fmt.Printf("Consuming %s\n", series)
    params := make(map[string]interface{})
    params["series_id"] = series.Name

    srs, err := client.GetSeriesObservations(params)

    if err != nil {
        fmt.Printf("Error retrieving series %s\n", series)
        fmt.Println(err)
    }

    fmt.Println(srs)
    model.InsertFredObs(srs, series)
}