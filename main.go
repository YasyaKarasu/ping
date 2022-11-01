package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func main() {
	e := echo.New()
	e.HideBanner = true

	e.GET("/api/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	logrus.Info("Echo framework initialized")

	e.Logger.Fatal(e.Start(":10001"))
}
