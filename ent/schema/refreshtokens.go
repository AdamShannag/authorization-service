package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RefreshTokens holds the schema definition for the RefreshTokens entity.
type RefreshTokens struct {
	ent.Schema
}

// Fields of the RefreshTokens.
func (RefreshTokens) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.Bool("active"),
	}
}

// Edges of the RefreshTokens.
func (RefreshTokens) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("request_id", Request.Type).
			Unique().
			Ref("refresh_token"),
	}
}
