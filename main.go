package main

import (
	"fmt"
    "io/ioutil"
    "log"
    "net/http"
    "os"
)

func main() {
	fmt.Println("Hello World")

	response, err := http.Get("https://api.kanye.rest/")
	if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(responseData) + " Zack and John")
}