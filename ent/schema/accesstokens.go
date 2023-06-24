package schema

import (
	"net/url"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"golang.org/x/text/language"
)

// AccessTokens holds the schema definition for the AccessTokens entity.
type AccessTokens struct {
	ent.Schema
}

// Fields of the AccessTokens.
func (AccessTokens) Fields() []ent.Field {
	return request_session_fields()
}

// Edges of the AccessTokens.
func (AccessTokens) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("client_id", Clients.Type).
			Unique().
			Ref("access_token"),
		edge.From("session_id", Session.Type).
			Unique().
			Ref("access_token"),
	}
}

func request_session_fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("request_id"),
		field.Time("requestedAt"),
		field.Strings("scopes"),
		field.Strings("granted_scopes"),
		field.Strings("requested_audience"),
		field.Strings("granted_audience"),
		field.JSON("form", url.Values{}),
		field.JSON("lang", language.Tag{}).Optional(),
	}
}
