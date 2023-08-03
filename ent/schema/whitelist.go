package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// WhiteList holds the schema definition for the WhiteList entity.
type WhiteList struct {
	ent.Schema
}

// Fields of the WhiteList.
func (WhiteList) Fields() []ent.Field {
	return []ent.Field{
		field.String("path"),
	}
}

// Edges of the WhiteList.
func (WhiteList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", ServiceProvider.Type).
			Ref("whitelists").
			Unique(),
	}
}
