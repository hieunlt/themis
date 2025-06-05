// Package main is the entry point for the Themis service.
// Themis provides a RESTful API for managing reviews and presets.
package main

import (
	"context"

	"github.com/hieunlt/themis/internal"
	"github.com/hieunlt/themis/internal/modules"
	"github.com/matzefriedrich/parsley/pkg/bootstrap"
)

// main initializes and starts the Themis service using the Parsley application framework.
// It configures the necessary dependencies including:
// - Fiber web framework
// - Database client
// - Review module
// - Preset module
// The application will panic if initialization fails.
func main() {
	ctx := context.Background()

	err := bootstrap.RunParsleyApplication(
		ctx,
		internal.NewApp,
		modules.ConfigureFiber,
		modules.ConfigureDBClient,
		modules.ConfigureReview,
		modules.ConfigurePreset,
	)
	if err != nil {
		panic(err)
	}
}
