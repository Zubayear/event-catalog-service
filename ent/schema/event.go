package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.String("Name"),
		field.Float32("Price").
			SchemaType(map[string]string{
				dialect.MySQL: "decimal(6,2)",
			}),
		field.String("Artist"),
		field.Int64("Date"),
		field.String("Description"),
		field.String("ImageUrl"),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Category.Type).
			Ref("Events").
			Unique(),
	}
}
