package modules

import (
	"github.com/hieunlt/themis/internal/services"
	"github.com/matzefriedrich/parsley/pkg/types"
)

// ConfigurePreset configures the Preset service and registers it with the service registry.
// It registers the Preset service as a transient service, meaning that a new instance of the service will be created each time it is requested.
func ConfigurePreset(registry types.ServiceRegistry) error {
	return registry.Register(services.NewPresetService, types.LifetimeTransient)
}
