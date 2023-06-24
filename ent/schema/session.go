package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.JSON("expires_at", map[string]time.Time{}),
		field.String("username").Optional(),
		field.String("subject").Optional(),
		field.JSON("extra", map[string]any{}).Optional(),
		field.Any("session").Optional(),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("access_token", AccessTokens.Type),
		edge.To("authorize_code", AuthorizeCodes.Type),
		edge.To("refresh_token", RefreshTokens.Type),
		edge.To("id_session", IDSessions.Type),
		edge.To("pkce", PKCES.Type),
	}
}
