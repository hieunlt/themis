package modules

import (
	"github.com/hieunlt/themis/internal/handlers"
	"github.com/matzefriedrich/parsley/pkg/features"
	"github.com/matzefriedrich/parsley/pkg/registration"
	"github.com/matzefriedrich/parsley/pkg/types"
)

// RegisterHandlers registers the handlers with the service registry.
// It registers the handlers as transient services, meaning that a new instance of each handler will be created each time it is requested.
func RegisterHandlers(registry types.ServiceRegistry) error {
	err := features.RegisterList[handlers.Handler](registry)
	if err != nil {
		return err
	}
	return registration.RegisterTransient(registry,
		handlers.NewReviewHandler,
		handlers.NewPresetHandler,
	)
}
