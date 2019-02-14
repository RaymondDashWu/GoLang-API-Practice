// API to use
// http://api.icndb.com/jokes/random

package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type norrisJoke struct {
	Value struct {
		Joke string `json:"joke"`
	} `json:"value"`
}

func main2() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	resp, _ := http.Get("http://api.icndb.com/jokes/random")

	// Route => handler
	e.GET("/", func(c echo.Context) error {
		fmt.Println(json.Marshal(&resp))
		return nil
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
