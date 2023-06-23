package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// IDSessions holds the schema definition for the IDSessions entity.
type IDSessions struct {
	ent.Schema
}

// Fields of the IDSessions.
func (IDSessions) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
	}
}

// Edges of the IDSessions.
func (IDSessions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("request_id", Request.Type).
			Unique().
			Ref("id_session"),
	}
}
