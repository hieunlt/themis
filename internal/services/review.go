package services

import (
	"context"

	"github.com/hieunlt/themis/ent"
)

// Review defines the interface for review operations,
// providing a contract for interacting with reviews.
type Review interface {
	Create(userID string, targetID string, rating uint8, comment string, presetIDs []int) (*ent.Review, error)
}

// review implements the Review interface, providing concrete
// implementations for review-related operations.
type review struct {
	Client *ent.Client
}

// Create creates a new review.
// It takes the user ID, target ID, rating, comment, and a list of preset IDs as input.
// It returns the created review or an error.
func (r *review) Create(userID string, targetID string, rating uint8, comment string, presetIDs []int) (*ent.Review, error) {
	return r.Client.Review.Create().
		SetUserID(userID).
		SetTargetID(targetID).
		SetComment(comment).
		SetRating(rating).
		AddPresetIDs(presetIDs...).
		Save(context.Background())
}

// NewReviewService creates a new review service,
// injecting the ent.Client for database interaction.
func NewReviewService(client *ent.Client) Review {
	return &review{client}
}
