// Package handlers provides HTTP request handlers and routing configuration for the Themis API.
// It defines common interfaces and utilities used across different handler implementations.
package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// Handler defines the interface that must be implemented by route handler services.
// A handler is responsible for registering its routes with a given Fiber router.
type Handler interface {
	// Register takes a Fiber router and registers the handler's routes with it.
	Register(router fiber.Router)
}
