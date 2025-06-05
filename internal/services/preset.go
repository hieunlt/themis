package services

import (
	"context"

	"github.com/hieunlt/themis/ent"
)

// Preset defines the interface for preset operations,
// providing a contract for interacting with presets.
type Preset interface {
	Create(display string, isPositive bool) (*ent.Preset, error)
}

// preset implements the Preset interface, providing concrete
// implementations for preset-related operations.
type preset struct {
	Client *ent.Client
}

// Create creates a new preset.
// It takes the display string and a boolean indicating if it's positive as input.
// It returns the created preset or an error.
func (p *preset) Create(display string, isPositive bool) (*ent.Preset, error) {
	return p.Client.Preset.Create().
		SetDisplay(display).
		SetIsPositive(isPositive).
		Save(context.Background())
}

// NewPresetService creates a new preset service,
// injecting the ent.Client for database interaction.
func NewPresetService(client *ent.Client) Preset {
	return &preset{Client: client}
}
