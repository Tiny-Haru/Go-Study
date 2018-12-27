package main

import (
	"github.com/labstack/echo"
	"fmt"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		fmt.Println("Success")

		return c.String(200, "Hello, World!")
	})

	e.Start(":8080")
}