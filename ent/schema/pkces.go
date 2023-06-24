package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// PKCES holds the schema definition for the PKCES entity.
type PKCES struct {
	ent.Schema
}

// Fields of the PKCES.
func (PKCES) Fields() []ent.Field {
	return request_session_fields()
}

// Edges of the PKCES.
func (PKCES) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("client_id", Clients.Type).
			Unique().
			Ref("pkce"),
		edge.From("session_id", Session.Type).
			Unique().
			Ref("pkce"),
	}
}
