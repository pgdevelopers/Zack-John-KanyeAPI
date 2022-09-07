package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Quote string `json:"quote"`
}

func getQuote(c *gin.Context) {
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
	responseObject.Quote += " - Kanye - Zack and John"
	c.IndentedJSON(http.StatusOK, responseObject)
}

func main() {
	router := gin.Default()
	router.GET("/quote", getQuote)
	router.Run("localhost:8080")
}
