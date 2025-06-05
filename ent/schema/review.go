package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Review holds the schema definition for the Review entity.
type Review struct {
	ent.Schema
}

// Fields of the Review.
func (Review) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").Immutable().NotEmpty().Sensitive().Comment("user submitted the rating"),
		field.String("target_id").Immutable().NotEmpty().Comment("the entity being rated"),
		field.Uint8("rating").Min(1).Max(5),
		field.String("comment").Optional(),
		field.Time("created_at").Nillable().Default(time.Now),
		field.Time("updated_at").Nillable().Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Review.
func (Review) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("presets", Preset.Type),
	}
}
