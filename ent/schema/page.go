package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
	}
}

func (Page) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("url").Unique(),
	}
}

// func (Page) Edges() []ent.Edge {
// 	return []ent.Edge{
// 		edge.From("owner", Project.Type).
// 			Ref("project").
// 			Unique(),
// 	}
// }
