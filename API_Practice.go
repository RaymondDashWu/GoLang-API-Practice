package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

// API to use
// http://ron-swanson-quotes.herokuapp.com/v2/quotes

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		resp, err := http.Get("http://ron-swanson-quotes.herokuapp.com/v2/quotes")

		if err != nil {
			log.Fatal("Request could not be completed")
		}
		// return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
