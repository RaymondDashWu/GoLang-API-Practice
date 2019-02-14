// API to use
// http://api.icndb.com/jokes/random

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type norrisJoke struct {
	Value struct {
		Joke string `json:"joke"`
	} `json:"value"`
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		resp, err := http.Get("http://api.icndb.com/jokes/random")
		readResp, err := ioutil.ReadAll(resp.Body)
		joke := norrisJoke{}

		fmt.Println("resp:", resp)
		fmt.Println("readResp:", readResp)
		fmt.Println("joke:", joke)

		if err != nil {
			log.Fatal(err)
			// Route => handler
		}

		return json.Unmarshal(readResp, &joke)
	})
	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
