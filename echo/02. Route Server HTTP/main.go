package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()

	e.DELETE("/", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	e.POST("/", func(c echo.Context) error {
		// echo 구조체의 메소드로 표준 HTTP 메소드 라우팅 지원
		return c.NoContent(http.StatusOK)
	})

	e.Start(":8080")
}