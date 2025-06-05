// Package internal contains the core application logic and components for the Themis service.
// It provides the main application structure and configuration for the HTTP server.
package internal

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/hieunlt/themis/internal/handlers"
	"github.com/matzefriedrich/parsley/pkg/bootstrap"
	"github.com/spf13/viper"
)

// parsleyApplication implements the bootstrap.Application interface and manages the
// Fiber HTTP server instance. It provides the main entry point for running the service.
type parsleyApplication struct {
	app *fiber.App
}

// Run starts the HTTP server on the configured port. It implements the bootstrap.Application
// interface and blocks until the server is shutdown or encounters an error.
func (a *parsleyApplication) Run(_ context.Context) error {
	port := viper.GetInt("SERVER_PORT")
	return a.app.Listen(fmt.Sprintf(":%d", port))
}

// NewApp creates a new parsleyApplication instance with the provided Fiber app and handlers.
// It configures the base API routes and health check middleware.
// The handlers are registered under the /api/v1 path prefix.
func NewApp(app *fiber.App, handlers []handlers.Handler) bootstrap.Application {
	app.Use(healthcheck.New())
	api := app.Group("/api")
	v1 := api.Group("/v1")
	for _, handler := range handlers {
		handler.Register(v1)
	}
	return &parsleyApplication{app}
}
