package utils

import (
    "bufio"
    "fmt"
    "os"
)

func ToStringArray(path string) ([]string, error) {
    file, err := os.Open(path)

    if err != nil {
        return nil, err
    }

    defer file.Close()

    var result []string
    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        fmt.Println(scanner.Text())
        result = append(result, scanner.Text())
    }

    return result, scanner.Err()
}
