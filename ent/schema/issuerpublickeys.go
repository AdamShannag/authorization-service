package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// IssuerPublicKeys holds the schema definition for the IssuerPublicKeys entity.
type IssuerPublicKeys struct {
	ent.Schema
}

// Fields of the IssuerPublicKeys.
func (IssuerPublicKeys) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
	}
}

// Edges of the IssuerPublicKeys.
func (IssuerPublicKeys) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subject_public_key", SubjectPublicKeys.Type).
			Unique(),
	}
}
