package internal

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/hieunlt/themis/internal/route_handlers"
	"github.com/matzefriedrich/parsley/pkg/bootstrap"
	"github.com/spf13/viper"
)

type parsleyApplication struct {
	app *fiber.App
}

func (a *parsleyApplication) Run(_ context.Context) error {
	port := viper.GetInt("SERVER_PORT")
	return a.app.Listen(fmt.Sprintf(":%d", port))
}

var _ bootstrap.Application = &parsleyApplication{}

func NewApp(app *fiber.App, routeHandlers []route_handlers.RouteHandler) bootstrap.Application {
	app.Use(healthcheck.New())
	api := app.Group("/api")
	v1 := api.Group("/v1")
	for _, routeHandler := range routeHandlers {
		routeHandler.Register(v1)
	}
	return &parsleyApplication{app}
}
