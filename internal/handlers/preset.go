// Package handlers provides HTTP request handlers for the Themis service API.
// It implements the routing and request/response handling logic for all endpoints.
package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hieunlt/themis/internal/services"
)

// presetHandler implements the Handler interface for preset-related endpoints.
// It manages the creation and management of preset templates used in reviews.
type presetHandler struct {
	preset services.Preset
}

// Register sets up the routing for preset-related endpoints under the /preset path.
// It configures the following routes:
// - POST /preset: Create a new preset
func (h *presetHandler) Register(router fiber.Router) {
	endpoint := router.Group("/preset")
	endpoint.Post("/", h.HandleCreatePreset)
}

// HandleCreatePreset processes POST requests to create new presets.
// It expects a JSON body with 'display' and 'is_positive' fields.
// Returns the created preset or an error if creation fails.
func (h *presetHandler) HandleCreatePreset(c *fiber.Ctx) error {
	type input struct {
		Display    string `json:"display"`
		IsPositive bool   `json:"is_positive"`
	}

	var body input

	err := parseRequestBody(c, &body)
	if err != nil {
		return err
	}

	preset, err := h.preset.Create(body.Display, body.IsPositive)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(preset)
}

// NewPresetHandler creates a new presetHandler instance with the given Preset service.
// It implements the Handler interface for preset-related functionality.
func NewPresetHandler(preset services.Preset) Handler {
	return &presetHandler{preset}
}
