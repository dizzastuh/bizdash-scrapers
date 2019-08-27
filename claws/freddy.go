// Scrapes from FRED
package claws

import (
    "fmt"
    "os"
    "github.com/nswekosk/fred_go_toolkit"
)

func ScrapeFreddy() {
    seriesList := initSeries()
    client, err := getClient()

    if err != nil {
        consumeAllSeries(seriesList, client)
    } else {
        fmt.Println(err)
    }
}

func consumeAllSeries(seriesList []string, client *FredClient) {
    for i:= 0; i < len(seriesList); i++ {
        consumeSeries(seriesList[i], client)
    }
}

func consumeSeries(series string, client *FredClient) {
    params := make(map[string]interface{})
    params["series"] = series

    fc, err := client.GetSeries(params)

    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("WE GOT A RESPONSE FROM %s", series)
    }
}

func getClient() (*FredClient, error) {
    fredConfig := FredConfig{ APIKey: os.Getenv("FRED_API_KEY"), FileType: FileTypeJSON, LogFile: "fredlog.log", }
    return CreateFredClient(fredConfig)
}

func initSeries() []string {
    seriesList := make([]string, 11)

    seriesList[0] = "PCE"
    seriesList[1] = "RRSFS"
    seriesList[2] = "RSAFS"
    seriesList[3] = "DSPI"
    seriesList[4] = "PAYEMS"
    seriesList[5] = "ICSA"
    seriesList[6] = "TOTALSA"
    seriesList[7] = "INDPRO"
    seriesList[8] = "NEWORDER"
    seriesList[9] = "DGORDER"
    seriesList[10] = "WHLSLRIMSA"

    return seriesList
}
