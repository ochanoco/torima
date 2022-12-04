package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// ServiceProvider holds the schema definition for the ServiceProvider entity.
type ServiceProvider struct {
	ent.Schema
}

// Fields of the ServiceProvider.
func (ServiceProvider) Fields() []ent.Field {
	return []ent.Field{
		field.String("host"),
		field.String("destination"),
	}
}

// Edges of the ServiceProvider.
func (ServiceProvider) Edges() []ent.Edge {
	return nil
}
