package modules

import (
	"github.com/hieunlt/themis/internal/services"
	"github.com/matzefriedrich/parsley/pkg/types"
)

// ConfigureReview configures the review service and registers it with the service registry.
// It registers the ReviewService as a transient service.
func ConfigureReview(registry types.ServiceRegistry) error {
	return registry.Register(services.NewReviewService, types.LifetimeTransient)
}
