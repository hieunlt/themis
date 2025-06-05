package main

import (
	"context"

	"github.com/hieunlt/themis/internal"
	"github.com/hieunlt/themis/internal/modules"
	"github.com/matzefriedrich/parsley/pkg/bootstrap"
)

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
