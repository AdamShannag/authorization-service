package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// AccessTokens holds the schema definition for the AccessTokens entity.
type AccessTokens struct {
	ent.Schema
}

// Fields of the AccessTokens.
func (AccessTokens) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
	}
}

// Edges of the AccessTokens.
func (AccessTokens) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("request_id", Request.Type).
			Unique().
			Ref("access_token"),
	}
}
