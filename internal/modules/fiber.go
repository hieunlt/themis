package modules

import (
	"github.com/gofiber/fiber/v2"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/types"
)

var _ types.ModuleFunc = ConfigureFiber

// ConfigureFiber configures the Fiber web framework and registers it with the service registry.
// It registers the Fiber configuration, the Fiber app itself, and the handlers.
func ConfigureFiber(registry types.ServiceRegistry) error {
	err := registration.RegisterInstance(registry, fiber.Config{
		AppName:           "Themis",
		EnablePrintRoutes: true,
	})
	if err != nil {
		return err
	}

	err = registry.Register(newFiber, types.LifetimeSingleton)
	if err != nil {
		return err
	}

	return registry.RegisterModule(RegisterHandlers)
}

func newFiber(config fiber.Config) *fiber.App {
	return fiber.New(config)
}
