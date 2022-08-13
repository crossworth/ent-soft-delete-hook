package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Other holds the schema definition for the Other entity.
type Other struct {
	ent.Schema
}

// Fields of the Other.
func (Other) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Other.
func (Other) Edges() []ent.Edge {
	return nil
}
