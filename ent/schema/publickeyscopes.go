package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"gopkg.in/square/go-jose.v2"
)

// PublicKeyScopes holds the schema definition for the PublicKeyScopes entity.
type PublicKeyScopes struct {
	ent.Schema
}

// Fields of the PublicKeyScopes.
func (PublicKeyScopes) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.JSON("json_web_key", jose.JSONWebKey{}),
		field.Strings("scopes"),
	}
}

// Edges of the PublicKeyScopes.
func (PublicKeyScopes) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("subject_public_key_id", SubjectPublicKeys.Type).
			Unique().
			Ref("public_key_scope"),
	}
}
