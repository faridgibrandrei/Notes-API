package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Hello Go")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Notes RestApi using Go")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
