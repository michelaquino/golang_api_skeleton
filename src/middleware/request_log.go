package middleware

import (
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// RequestLogger is a middleware to log the request data
func RequestLogger(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		request := c.Request()
		response := c.Response()

		start := time.Now()
		defer func() {
			slog.LogAttrs(
				c.Request().Context(),
				slog.LevelInfo,
				"request received",
				slog.Group("request",
					"method", request.Method,
					"request_uri", request.RequestURI,
					"status", strconv.Itoa(response.Status),
					"time", fmt.Sprintf("%d%s", time.Since(start).Milliseconds(), "ms"),
				),
			)
		}()

		if err := next(c); err != nil {
			c.Error(err)
		}

		return nil
	}
}
