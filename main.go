package main

import (
	"fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
	"encoding/json"
)

type Response struct {
    Quote    string    `json:"quote"`
}

func main() {
	response, err := http.Get("https://api.kanye.rest/")
	if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
        log.Fatal(err)
    }

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Quote + " - Kanye - Zack and John")
}