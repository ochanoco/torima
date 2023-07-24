package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ServiceLog struct {
	ent.Schema
}

// Fields of the User.
func (ServiceLog) Fields() []ent.Field {
	return []ent.Field{
		field.Time("time"),
		field.String("headers"),
		field.Bytes("body").Optional(),
	}
}

// Edges of the User.
func (ServiceLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("hashchains", HashChain.Type),
	}
}
