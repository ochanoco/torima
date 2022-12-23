package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// ServiceProvider holds the schema definition for the ServiceProvider entity.
type ServiceProvider struct {
	ent.Schema
}

// Fields of the ServiceProvider.
func (ServiceProvider) Fields() []ent.Field {
	return []ent.Field{
		field.String("host"),
		field.String("destination_ip"),
	}
}

// Edges of the Project.
func (ServiceProvider) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("whitelists", WhiteList.Type),
		edge.To("authorization_codes", AuthorizationCode.Type),
	}
}

func (ServiceProvider) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("host").Unique(),
	}
}
