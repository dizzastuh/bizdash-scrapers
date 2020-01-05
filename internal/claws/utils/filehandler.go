package utils

import (
    "bufio"
    "log"
    "os"
    "encoding/json"
    "io/ioutil"
)

func ToStringArray(path string) ([]string, error) {
    file := OpenFile(path)
    defer file.Close()

    var result []string
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        result = append(result, scanner.Text())
    }

    return result, scanner.Err()
}

func ParseJSONFile(path string) (map[string]interface{}) {
    file := OpenFile(path)
    defer file.Close()

    bytes, _ := ioutil.ReadAll(file)

    var result map[string]interface{}
    json.Unmarshal([]byte(bytes), &result)

    return result
}

func OpenFile(path string) (*os.File) {
    file, err := os.Open(path)

    if err != nil {
        log.Fatal(err)
    }

    return file
}
