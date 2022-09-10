package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
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

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("launches", Launch.Type),
	}
}

// Mixin of the Order.
func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		TimeMixin{},
	}
}

// Indexes of the Order.
func (Order) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at", "id"),
	}
}
