package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	apiMiddleware "github.com/michelaquino/golang_api_skeleton/src/middleware"
)

func configureMiddlewares(echoInstance *echo.Echo) {
	echoInstance.Pre(middleware.RemoveTrailingSlash())
	echoInstance.Use(apiMiddleware.AssignRequestID)
	echoInstance.Use(apiMiddleware.RequestLogger)
	echoInstance.Use(middleware.Recover())
}
