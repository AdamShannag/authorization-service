package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SubjectPublicKeys holds the schema definition for the SubjectPublicKeys entity.
type SubjectPublicKeys struct {
	ent.Schema
}

// Fields of the SubjectPublicKeys.
func (SubjectPublicKeys) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
	}
}

// Edges of the SubjectPublicKeys.
func (SubjectPublicKeys) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("public_key_scope", PublicKeyScopes.Type).
			Unique(),
		edge.From("issuer_public_key_id", IssuerPublicKeys.Type).
			Unique().
			Ref("subject_public_key"),
	}
}
