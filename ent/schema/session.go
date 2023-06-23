package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.JSON("expires_at", map[string]time.Time{}),
		field.String("username"),
		field.String("subject"),
		field.JSON("extra", map[string]any{}).Optional(),
	}
}

// Edges of the Session.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("requests", Request.Type),
	}
}
