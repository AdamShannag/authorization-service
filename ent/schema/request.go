package schema

import (
	"net/url"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"golang.org/x/text/language"
)

// Request holds the schema definition for the Request entity.
type Request struct {
	ent.Schema
}

// Fields of the Request.
func (Request) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.Time("requestedAt"),
		field.Strings("scopes"),
		field.Strings("granted_scopes"),
		field.Strings("requested_audience"),
		field.Strings("granted_audience"),
		field.JSON("form", url.Values{}),
		field.JSON("lang", language.Tag{}).Optional(),
	}
}

// Edges of the Request.
func (Request) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("client_id", Clients.Type).
			Unique().
			Ref("requests"),
		edge.From("session_id", Session.Type).
			Unique().
			Ref("requests"),
		edge.To("refresh_token", RefreshTokens.Type).
			Unique(),
		edge.To("authorize_code", AuthorizeCodes.Type).
			Unique(),
		edge.To("access_token", AccessTokens.Type).
			Unique(),
		edge.To("id_session", IDSessions.Type).
			Unique(),
		edge.To("pkce", PKCES.Type).
			Unique(),
	}
}
