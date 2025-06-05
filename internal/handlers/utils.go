// Package handlers provides HTTP request handlers for the Themis service API.
// It defines utils for request handling
package handlers

import "github.com/gofiber/fiber/v2"

// parseRequestBody attempts to parse the request body into the provided struct.
// It returns an error if parsing fails.
func parseRequestBody(c *fiber.Ctx, body any) error {
	if err := c.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return nil
}
