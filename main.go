package main

import (
	"context"
	"flag"
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/martinborgman/websay/pkg/say"
)

func main() {
	var bgColor string
	flag.StringVar(&bgColor, "bg-color", "white", "background color")

	var fgColor string
	flag.StringVar(&fgColor, "fg-color", "black", "foreground color")

	var message string
	flag.StringVar(&message, "message", "Hello", "message")

	var sayType string
	flag.StringVar(&sayType, "type", "default", "type")

	flag.Parse()

	e := echo.New()

	//e.Use(middleware.Logger())

	e.GET("/", func(c *echo.Context) error {
		content, err := say.Say(message, bgColor, fgColor, sayType)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.HTML(http.StatusOK, content)
	})
	sc := echo.StartConfig{
		HideBanner: true,
		HidePort:   true,
	}
	if err := sc.Start(context.Background(), e); err != nil {
		e.Logger.Error("Failed to start server", "error", err)
	}
}
