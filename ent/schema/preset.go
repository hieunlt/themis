package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Preset holds the schema definition for the Preset entity.
type Preset struct {
	ent.Schema
}

// Fields of the Preset.
func (Preset) Fields() []ent.Field {
	return []ent.Field{
		field.String("display").NotEmpty(),
		field.Bool("is_active").Default(true),
		field.Bool("is_positive").Immutable(),
		field.Time("created_at").Nillable().Default(time.Now),
		field.Time("updated_at").Nillable().Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Preset.
func (Preset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("reviews", Review.Type).
			Ref("presets"),
	}
}
