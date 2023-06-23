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
	return []ent.Field{
		field.String("id").Unique(),
		field.Bool("active"),
	}
}

// Edges of the AuthorizeCodes.
func (AuthorizeCodes) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("request_id", Request.Type).
			Unique().
			Ref("authorize_code"),
	}
}
