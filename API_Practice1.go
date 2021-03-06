package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// API to use
// http://ron-swanson-quotes.herokuapp.com/v2/quotes

// Reference
// https://stackoverflow.com/questions/17156371/how-to-get-json-response-in-golang

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJSON(url string, target *swansonQuote) error {
	resp, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	res := []string{}
	json.NewDecoder(resp.Body).Decode(&res)
	target.Quote = res[0]
	return nil
}

type swansonQuote struct {
	Quote string `json:"quote"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		// resp, err := http.Get("http://ron-swanson-quotes.herokuapp.com/v2/quotes")

		// if err != nil {
		// 	log.Fatal("Request could not be completed")
		// }
		// return c.String(http.StatusOK, resp)
		testQuote := swansonQuote{}
		getJSON("http://ron-swanson-quotes.herokuapp.com/v2/quotes", &testQuote)
		return c.JSON(http.StatusOK, testQuote)

	})
	e.Logger.Fatal(e.Start(":1323"))
}
