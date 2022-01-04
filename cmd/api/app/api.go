package app

import (
	"net"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	"github.com/neko-neko/echo-logrus/v2/log"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/app/url"
	"github.com/sirupsen/logrus"
)

func App() {
	startServer()
}

func startServer() {
	e := echo.New()
	// Logger
	log.Logger().SetOutput(os.Stdout)
	log.Logger().SetLevel(echoLog.INFO)
	log.Logger().SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})
	e.Logger = log.Logger()
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())

	e.POST("/api/create", url.Create)

	port := os.Getenv("PORT")
	if port == "" {
		e.Logger.Fatal("Port is not provided")
	}
	e.Logger.Fatal(e.Start(net.JoinHostPort("", port)))
}
