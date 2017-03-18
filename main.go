package main

import (
	"github.com/labstack/echo"
	"github.com/michelaquino/golang_api_skeleton/context"
	"github.com/michelaquino/golang_api_skeleton/handlers"
	apiMiddleware "github.com/michelaquino/golang_api_skeleton/middleware"
	"github.com/michelaquino/golang_api_skeleton/repository"
)

func init() {
	context.GetMongoSession()
}

func main() {
	logger := context.GetLogger()

	echoInstance := echo.New()
	echoInstance.Use(apiMiddleware.RequestLogDataMiddleware())

	configureHealthcheckRoute(echoInstance)
	configureUserRoutes(echoInstance)

	logger.Info("Main", "main", "", "", "start app", "success", "Started at port 8888!")
	echoInstance.Logger.Fatal(echoInstance.Start(":8888"))
}

func configureHealthcheckRoute(echoInstance *echo.Echo) {
	echoInstance.GET("/healthcheck", handlers.Healthcheck)
}

func configureUserRoutes(echoInstance *echo.Echo) {
	userRepository := new(repository.UserMongoRepository)
	userHandler := handlers.NewUserHandler(userRepository)

	userGroup := echoInstance.Group("/user")
	userGroup.POST("", userHandler.CreateUser)
}
