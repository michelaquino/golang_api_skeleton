package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/michelaquino/golang_api_skeleton/context"
	apiMiddleware "github.com/michelaquino/golang_api_skeleton/middleware"
	"github.com/michelaquino/golang_api_skeleton/models"
)

var userHandlerLog context.Logger

func init() {
	userHandlerLog = context.GetLogger()
}

// Healthcheck is a method that respond only WORKING
func Healthcheck(echoContext echo.Context) error {
	logger := context.GetLogger()
	requestLogData := echoContext.Get(apiMiddleware.RequestIDKey).(models.RequestLogData)

	logger.Info("Handlers", "Healthcheck", requestLogData.ID, requestLogData.OriginIP, "Verify Healthcheck", "success", "")
	return echoContext.String(http.StatusOK, "WORKING")
}
