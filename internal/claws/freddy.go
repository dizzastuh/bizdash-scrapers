package claws

import (
    "fmt"
    "os"

    . "github.com/nswekosk/fred_go_toolkit"
)

func ScrapeFreddy() {
    seriesList := initSeries()
    client, err := getClient()

    if err == nil {
        // consumeTestSeries(client)
        consumeAllSeries(seriesList, client)
    } else {
        fmt.Println("Client could not be acquired")
        fmt.Println(err)
    }
}

func consumeAllSeries(seriesList []string, client *FredClient) {
    // for i:= 0; i < len(seriesList); i++ {
    for i:= 0; i < 1; i++ {
        consumeSeries(seriesList[i], client)
    }
}

func consumeSeries(series string, client *FredClient) {
    fmt.Printf("Consuming %s\n", series)
    params := make(map[string]interface{})
    params["series_id"] = series

    fmt.Println(params)

    srs, err := client.GetSeriesObservations(params)

    if err == nil {

        for i:= 0; i < len(srs.Observations); i++ {
            obs := srs.Observations[i]
            fmt.Println(obs.Value)
        }

    } else {
        fmt.Printf("Error retrieving series %s\n", series)
        fmt.Println(err)
    }
}

func getClient() (*FredClient, error) {
    fredConfig := FredConfig{ APIKey: os.Getenv("FRED_API_KEY"), FileType: FileTypeJSON, LogFile: "fred.log", }
    return CreateFredClient(fredConfig)
}

func initSeries() []string {
    seriesList := make([]string, 24)

    seriesList[1] = "T10Y2Y"
    seriesList[2] = "ACDGNO"
    seriesList[3] = "PERMIT"
    seriesList[4] = "UNRATE"
    seriesList[5] = "INDPRO"
    seriesList[6] = "USSLIND"
    seriesList[7] = "NFCI"
    seriesList[8] = "RRSFS"
    seriesList[9] = "BAMLH0A0HYM2"
    seriesList[10] = "PCE"
    seriesList[11] = "RRSFS"
    seriesList[12] = "RSAFS"
    seriesList[13] = "DSPI"
    seriesList[14] = "PAYEMS"
    seriesList[15] = "ICSA"
    seriesList[16] = "TOTALSA"
    seriesList[17] = "INDPRO"
    seriesList[18] = "NEWORDER"
    seriesList[19] = "DGORDER"
    seriesList[20] = "WHLSLRIMSA"
    seriesList[21] = "CPIAUCNS"
    seriesList[22] = "CPILFENS"
    seriesList[23] = "PCEPI"
    seriesList[23] = "WPUFD49207"

    return seriesList
}
