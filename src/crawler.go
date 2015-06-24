package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "log"
)

func main() {
    response, err := http.Get("http://google.com/")
    defer response.Body.Close()

    if err != nil {
        log.Fatalf("ERROR: %s", err)
    }

    body, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatalf("ERROR: %s", err)
    }

    fmt.Printf("%s", body)
}