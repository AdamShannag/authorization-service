package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AuthorizeCodes holds the schema definition for the AuthorizeCodes entity.
type AuthorizeCodes struct {
	ent.Schema
}

// Fields of the AuthorizeCodes.
func (AuthorizeCodes) Fields() []ent.Field {
	f := request_session_fields()
	f = append(f, field.Bool("active"))
	return f
}

// Edges of the AuthorizeCodes.
func (AuthorizeCodes) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("client_id", Clients.Type).
			Unique().
			Ref("authorize_code"),
		edge.From("session_id", Session.Type).
			Unique().
			Ref("authorize_code"),
	}
}
