package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Launch holds the schema definition for the Launch entity.
type Launch struct {
	ent.Schema
}

// Fields of the Launch.
func (Launch) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("status").
			Values("todo", "started", "completed").
			Default("todo"),
		field.Text("reason").
			NotEmpty(),
		field.Text("description").
			Optional().
			Nillable(),
		field.Strings("files").
			Optional(),
		field.String("applicant").
			NotEmpty(),
	}
}

// Edges of the Launch.
func (Launch) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("launches").
			Unique().
			Required(),
		edge.From("kind", Kind.Type).
			Ref("launches").
			Unique().
			Required(),
	}
}

// Mixin of the Launch.
func (Launch) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		TimeMixin{},
	}
}

// Indexes of the Launch.
func (Launch) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at", "id"),
	}
}
