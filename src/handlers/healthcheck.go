package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Healthcheck is a method that responds only WORKING.
func Healthcheck(echoContext echo.Context) error {
	return echoContext.String(http.StatusOK, "WORKING")
}
