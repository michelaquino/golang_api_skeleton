package server

import (
	"context"
	"log/slog"
	"net/http"
	"net/url"

	"github.com/labstack/echo/v4"
	apiMiddleware "github.com/michelaquino/golang_api_skeleton/src/middleware"
	newrelic "github.com/newrelic/go-agent"
	"github.com/spf13/viper"
)

// configureNewRelic is the method that enables New Relic.
func configureNewRelic(ctx context.Context, echoInstance *echo.Echo) {
	newRelicEnable := viper.GetBool("new_relic.is.enabled")
	if !newRelicEnable {
		return
	}

	newRelicApp, err := createNewRelicApp(ctx)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
	}

	echoInstance.Use(apiMiddleware.NewRelicMiddleware(newRelicApp))
}

// createNewRelicApp is the method that creates New Relic configuration.
func createNewRelicApp(ctx context.Context) (newrelic.Application, error) {
	licenseKeyEnvVar := viper.GetString("new_relic.licence.key")

	config := newrelic.NewConfig("My Awesome API", licenseKeyEnvVar)
	proxyURL, err := url.Parse(viper.GetString("new_relic.proxy.url"))
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
	}

	config.Transport = &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	newRelicApp, err := newrelic.NewApplication(config)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return nil, err
	}

	return newRelicApp, nil
}
