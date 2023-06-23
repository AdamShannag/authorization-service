package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// BlacklistedJTIs holds the schema definition for the BlacklistedJTIs entity.
type BlacklistedJTIs struct {
	ent.Schema
}

// Fields of the BlacklistedJTIs.
func (BlacklistedJTIs) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.Time("expiry"),
	}
}

// Edges of the BlacklistedJTIs.
func (BlacklistedJTIs) Edges() []ent.Edge {
	return nil
}
