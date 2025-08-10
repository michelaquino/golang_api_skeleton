package server

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// Start HTTP server
func Start() {
	echoInstance := echo.New()
	ctx := context.Background()

	// Configure New Relic
	configureNewRelic(ctx, echoInstance)

	// Middlewares
	configureMiddlewares(echoInstance)

	// Configure routes
	configureRoutes(echoInstance)

	port := viper.GetInt("api.host.port")
	slog.InfoContext(ctx, "start api", "port", port)

	echoInstance.Logger.Fatal(echoInstance.Start(fmt.Sprintf(":%d", port)))
}
