// Code generated by ent, DO NOT EDIT.

package accesstokens

import (
	"authorization-service/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.AccessTokens {
	return predicate.AccessTokens(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.AccessTokens {
	return predicate.AccessTokens(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.AccessTokens {
	return predicate.AccessTokens(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.AccessTokens {
	return predicate.AccessTokens(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.AccessTokens {
	return predicate.AccessTokens(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.AccessTokens {
	return predicate.AccessTokens(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.AccessTokens {
	return predicate.AccessTokens(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.AccessTokens {
	return predicate.AccessTokens(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.AccessTokens {
	return predicate.AccessTokens(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.AccessTokens {
	return predicate.AccessTokens(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.AccessTokens {
	return predicate.AccessTokens(sql.FieldContainsFold(FieldID, id))
}

// HasRequestID applies the HasEdge predicate on the "request_id" edge.
func HasRequestID() predicate.AccessTokens {
	return predicate.AccessTokens(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, RequestIDTable, RequestIDColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRequestIDWith applies the HasEdge predicate on the "request_id" edge with a given conditions (other predicates).
func HasRequestIDWith(preds ...predicate.Request) predicate.AccessTokens {
	return predicate.AccessTokens(func(s *sql.Selector) {
		step := newRequestIDStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AccessTokens) predicate.AccessTokens {
	return predicate.AccessTokens(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AccessTokens) predicate.AccessTokens {
	return predicate.AccessTokens(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AccessTokens) predicate.AccessTokens {
	return predicate.AccessTokens(func(s *sql.Selector) {
		p(s.Not())
	})
}
