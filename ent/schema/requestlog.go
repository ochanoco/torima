package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type RequestLog struct {
	ent.Schema
}

// Fields of the User.
func (RequestLog) Fields() []ent.Field {
	return []ent.Field{
		field.Time("time"),
		field.String("headers"),
		field.Bytes("body").Optional(),
		field.String("flag").Default(""),
	}
}
