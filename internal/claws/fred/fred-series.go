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

func ConsumeAllSeries(client *FredClient) {    
    seriesPath, _ := filepath.Abs("./res/fred-series.json")

    seriesFile := utils.OpenFile(seriesPath)
    defer seriesFile.Close()

    bytes, _ := ioutil.ReadAll(seriesFile)

    var series model.FSeriesList
    json.Unmarshal(bytes, &series)

    for i:= 0; i < len(series.List); i++ {
        consumeSeries(&series.List[i], client)
    }
}

func consumeSeries(series *model.FSeries, client *FredClient) {
    fmt.Printf("Consuming %s\n", series.Name)
    params := make(map[string]interface{})
    params["series_id"] = series.Name

    srs, err := client.GetSeriesObservations(params)

    if err != nil {
        fmt.Printf("Error retrieving series %s\n", series.Name)
        fmt.Println(err)
    }

    model.InsertFredObs(srs, series)
}