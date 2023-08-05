package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type CommunicationLog struct {
	ent.Schema
}

// Fields of the User.
func (CommunicationLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("type"),
		field.Time("time"),
		field.String("headers"),
		field.Bytes("body").Optional(),
	}
}
