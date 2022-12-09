package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group schema.
type Group struct {
	ent.Schema
}

// Fields of the group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}

// Mixin of the User.
func (Group) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SoftDeleteMixin{},
	}
}