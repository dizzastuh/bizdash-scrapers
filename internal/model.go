package internal

// time formatting must be done with  Mon Jan 2 15:04:05 MST 2006

import (
    "fmt"
    "os"
    "net/http"
    "io/ioutil"
    "time"
)

type Timestamp time.Time

func Insert() {
    host := os.Getenv("DB_HOST")

    req, _ := http.NewRequest("POST", host, nil)
    req.Header.Add("Content-Encoding", "gzip")

    res, _ := http.DefaultClient.Do(req)
}

func Fetch() {
    host := os.Getenv("DB_HOST")

    req, _ := http.NewRequest("POST", host, nil)
    req.Header.Add("Accept-Encoding", "gzip")

    res, _ := http.DefaultClient.Do(req)
}