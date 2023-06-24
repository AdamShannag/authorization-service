package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// IDSessions holds the schema definition for the IDSessions entity.
type IDSessions struct {
	ent.Schema
}

// Fields of the IDSessions.
func (IDSessions) Fields() []ent.Field {
	return request_session_fields()
}

// Edges of the IDSessions.
func (IDSessions) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("client_id", Clients.Type).
			Unique().
			Ref("id_session"),
		edge.From("session_id", Session.Type).
			Unique().
			Ref("id_session"),
	}
}
