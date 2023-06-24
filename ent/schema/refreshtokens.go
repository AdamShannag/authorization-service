package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RefreshTokens holds the schema definition for the RefreshTokens entity.
type RefreshTokens struct {
	ent.Schema
}

// Fields of the RefreshTokens.
func (RefreshTokens) Fields() []ent.Field {
	f := request_session_fields()
	f = append(f, field.Bool("active"))
	return f
}

// Edges of the RefreshTokens.
func (RefreshTokens) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("client_id", Clients.Type).
			Unique().
			Ref("refresh_token"),
		edge.From("session_id", Session.Type).
			Unique().
			Ref("refresh_token"),
	}
}
