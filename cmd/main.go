package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/sse-notification/handler"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/events", handler.EventHandler)

	e.Start(":8005")
}