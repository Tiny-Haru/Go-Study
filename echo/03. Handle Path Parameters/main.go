package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/:name", func(c echo.Context) error {
		name := c.Param("name")
		// (c *context) Param(name string) string 함수를 통해 path parameter를 받아올 수 있다
		// 요청 URI에 Path parameter 자리가 비어있는 경우, 프레임워크에서 404 반환.

		return c.String(http.StatusOK, fmt.Sprintf("Hello %s.", name))
	})

	e.Start(":8080")
}