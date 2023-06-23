// Code generated by ent, DO NOT EDIT.

package issuerpublickeys

import (
	"authorization-service/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.IssuerPublicKeys {
	return predicate.IssuerPublicKeys(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.IssuerPublicKeys {
	return predicate.IssuerPublicKeys(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.IssuerPublicKeys {
	return predicate.IssuerPublicKeys(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.IssuerPublicKeys {
	return predicate.IssuerPublicKeys(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.IssuerPublicKeys {
	return predicate.IssuerPublicKeys(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.IssuerPublicKeys {
	return predicate.IssuerPublicKeys(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.IssuerPublicKeys {
	return predicate.IssuerPublicKeys(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.IssuerPublicKeys {
	return predicate.IssuerPublicKeys(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.IssuerPublicKeys {
	return predicate.IssuerPublicKeys(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.IssuerPublicKeys {
	return predicate.IssuerPublicKeys(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.IssuerPublicKeys {
	return predicate.IssuerPublicKeys(sql.FieldContainsFold(FieldID, id))
}

// HasSubjectPublicKey applies the HasEdge predicate on the "subject_public_key" edge.
func HasSubjectPublicKey() predicate.IssuerPublicKeys {
	return predicate.IssuerPublicKeys(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SubjectPublicKeyTable, SubjectPublicKeyColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubjectPublicKeyWith applies the HasEdge predicate on the "subject_public_key" edge with a given conditions (other predicates).
func HasSubjectPublicKeyWith(preds ...predicate.SubjectPublicKeys) predicate.IssuerPublicKeys {
	return predicate.IssuerPublicKeys(func(s *sql.Selector) {
		step := newSubjectPublicKeyStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.IssuerPublicKeys) predicate.IssuerPublicKeys {
	return predicate.IssuerPublicKeys(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.IssuerPublicKeys) predicate.IssuerPublicKeys {
	return predicate.IssuerPublicKeys(func(s *sql.Selector) {
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
func Not(p predicate.IssuerPublicKeys) predicate.IssuerPublicKeys {
	return predicate.IssuerPublicKeys(func(s *sql.Selector) {
		p(s.Not())
	})
}
