package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/michelaquino/golang_api_skeleton/context"
	apiMiddleware "github.com/michelaquino/golang_api_skeleton/middleware"
	"github.com/michelaquino/golang_api_skeleton/models"
)

var healthcheckHandlerLog context.Logger

func init() {
	healthcheckHandlerLog = context.GetLogger()
}

// Healthcheck is a method that respond only WORKING
func Healthcheck(echoContext echo.Context) error {
	requestLogData := echoContext.Get(apiMiddleware.RequestIDKey).(models.RequestLogData)

	healthcheckHandlerLog.Info("Handlers", "Healthcheck", requestLogData.ID, requestLogData.OriginIP, "Verify Healthcheck", "success", "")
	return echoContext.String(http.StatusOK, "WORKING")
}
