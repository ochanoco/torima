package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type HashChain struct {
	ent.Schema
}

// Fields of the User.
func (HashChain) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("hash"),
		field.Bytes("signature"),
	}
}

// Edges of the WhiteList.
func (HashChain) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("log", ServiceLog.Type).
			Ref("hashchains").
			Unique(),
	}
}
