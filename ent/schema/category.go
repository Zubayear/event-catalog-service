package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Category holds the schema definition for the Category entity.
type Category struct {
	ent.Schema
}

// Fields of the Category.
func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Enum("Name").
			Values("CONCERTS", "MUSICALS", "PLAYS"),
	}
}

// Edges of the Category.
func (Category) Edges() []ent.Edge {
	// each category has many event
	return []ent.Edge{
		edge.To("Events", Event.Type),
	}
}
