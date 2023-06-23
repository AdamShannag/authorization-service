package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// PKCES holds the schema definition for the PKCES entity.
type PKCES struct {
	ent.Schema
}

// Fields of the PKCES.
func (PKCES) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
	}
}

// Edges of the PKCES.
func (PKCES) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("request_id", Request.Type).
			Unique().
			Ref("pkce"),
	}
}
