package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Kind holds the schema definition for the Kind entity.
type Kind struct {
	ent.Schema
}

// Fields of the Kind.
func (Kind) Fields() []ent.Field {
	return []ent.Field{
		field.Text("name").
			Unique().
			NotEmpty(),
		field.Text("description").
			Optional().
			Nillable(),
		field.Enum("status").
			Values("published", "draft").
			Default("draft"),
	}
}

// Edges of the Kind.
func (Kind) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("launches", Launch.Type),
	}
}

// Mixin of the Kind.
func (Kind) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		TimeMixin{},
	}
}

// Indexes of the Kind.
func (Kind) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at", "id"),
	}
}
