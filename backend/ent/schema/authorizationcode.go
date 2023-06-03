package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AuthorizationCode holds the schema definition for the AuthorizationCode entity.
type AuthorizationCode struct {
	ent.Schema
}

// Fields of the AuthorizationCode.
func (AuthorizationCode) Fields() []ent.Field {
	return []ent.Field{
		field.String("token"),
	}
}

// Edges of the AuthorizationCode.
func (AuthorizationCode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", ServiceProvider.Type).
			Ref("authorization_codes").
			Unique(),
	}
}
