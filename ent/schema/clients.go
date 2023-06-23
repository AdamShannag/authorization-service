package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Clients holds the schema definition for the Clients entity.
type Clients struct {
	ent.Schema
}

// Fields of the Clients.
func (Clients) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.Bytes("client_secret"),
		field.JSON("rotated_secrets", [][]byte{}),
		field.Strings("redirect_uris"),
		field.Strings("grant_types"),
		field.Strings("response_types"),
		field.Strings("scopes"),
		field.Strings("audience").Optional(),
		field.Bool("public").Optional(),
	}
}

// Edges of the Clients.
func (Clients) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("requests", Request.Type),
	}
}
