package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Page holds the schema definition for the Page entity.
type Page struct {
	ent.Schema
}

// Fields of the Page.
func (Page) Fields() []ent.Field {
	return []ent.Field{
		field.String("url"),
		field.Bool("skip"),
		field.Int("project_id"),
	}
}

// Edges of the Page.
func (Page) Edges() []ent.Edge {
	return nil
}
