package main

import (
	"github.com/labstack/echo"
	"github.com/michelaquino/golang_api_skeleton/context"
	"github.com/michelaquino/golang_api_skeleton/handlers"
	apiMiddleware "github.com/michelaquino/golang_api_skeleton/middleware"
)

func init() {
	context.GetMongoSession()
}

func main() {
	logger := context.GetLogger()

	echoInstance := echo.New()
	echoInstance.Use(apiMiddleware.RequestLogDataMiddleware())

	echoInstance.GET("/healthcheck", handlers.Healthcheck)

	logger.Info("Main", "main", "", "", "start app", "success", "Started at port 8888!")
	echoInstance.Logger.Fatal(echoInstance.Start(":8888"))
}
