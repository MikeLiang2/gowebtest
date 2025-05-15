// echo_server.go
package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func startEchoServer() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		msg := c.QueryParam("msg")
		if msg == "" {
			msg = "Nothing received (did you forget ?msg=... ?)"
		}
		fmt.Println("Received:", msg)
		return c.String(http.StatusOK, fmt.Sprintf("Hello Client, I got: %s", msg))
	})

	e.Start(":8080")
}
