package app

import (
	"net"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoLog "github.com/labstack/gommon/log"
	"github.com/neko-neko/echo-logrus/v2/log"
	echoDelivery "github.com/simonnik/GB_Backend1_CW_GO/internal/app/echo/delivery"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/app/link/delivery"
	linkRepo "github.com/simonnik/GB_Backend1_CW_GO/internal/app/link/repository/inmemory"
	linkUsecase "github.com/simonnik/GB_Backend1_CW_GO/internal/app/link/usecase"
	"github.com/simonnik/GB_Backend1_CW_GO/internal/app/middlewares"
	"github.com/sirupsen/logrus"
)

func App() {
	e := echo.New()
	// changed the error handler
	e.HTTPErrorHandler = echoDelivery.HTTPErrorHandler
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

	secretKey := os.Getenv("APP_SECRET")
	if secretKey == "" {
		e.Logger.Fatal("APP_SECRET is not provided")
	}
	authMiddleware := middlewares.JWTAuthMiddleware(secretKey)

	repository := linkRepo.New()
	linksUsecase := linkUsecase.New(repository)
	linksDelivery := delivery.New(linksUsecase)

	e.POST("/api/create", linksDelivery.Create, authMiddleware)

	port := os.Getenv("PORT")
	if port == "" {
		e.Logger.Fatal("Port is not provided")
	}
	e.Logger.Fatal(e.Start(net.JoinHostPort("", port)))
}
