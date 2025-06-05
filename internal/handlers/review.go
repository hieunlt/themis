// Package handlers provides HTTP request handlers for the Themis service API.
// It handles incoming HTTP requests, validates inputs, and coordinates with services.
package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hieunlt/themis/internal/services"
)

// reviewHandler implements the Handler interface for review-related endpoints.
// It manages the creation and retrieval of reviews in the system.
type reviewHandler struct {
	// review is the service that handles review-related business logic.
	review services.Review
}

// Register sets up the routing for review-related endpoints under the /review path.
// It configures the following routes:
// - POST /review: Create a new review
func (h *reviewHandler) Register(router fiber.Router) {
	// endpoint is a sub-router for review-related endpoints.
	endpoint := router.Group("/review")
	// POST /review: handles the creation of new reviews.
	endpoint.Post("/", h.HandleCreateRatingRequest)
}

// HandleCreateRatingRequest processes POST requests to create new reviews.
// It expects a JSON body with the following fields:
// - user_id: ID of the user submitting the review
// - target_id: ID of the entity being reviewed
// - rating: numeric rating (1-5)
// - comment: optional text comment
// - preset_ids: array of preset template IDs to associate
// Returns the created review or an error if validation/creation fails.
// HandleCreateRatingRequest processes POST requests to create new reviews.
func (h *reviewHandler) HandleCreateRatingRequest(c *fiber.Ctx) error {
	// input defines the expected JSON request body for creating a review.
	type input struct {
		UserID    string `json:"user_id"`
		TargetID  string `json:"target_id"`
		Rating    uint8  `json:"rating"`
		Comment   string `json:"comment"`
		PresetIDs []int  `json:"preset_ids"`
	}

	var body input

	// Parse the request body into the input struct.
	err := parseRequestBody(c, &body)
	if err != nil {
		return err
	}
	review, err := h.review.Create(body.UserID, body.TargetID, body.Rating, body.Comment, body.PresetIDs)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	// Return the created review as a JSON response.
	return c.JSON(review)
}

// NewReviewHandler creates a new reviewHandler instance with the given Review service.
// It implements the Handler interface for review-related functionality.
func NewReviewHandler(review services.Review) Handler {
	// Returns a pointer to a new reviewHandler instance.
	return &reviewHandler{review}
}
