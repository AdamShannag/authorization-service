// Code generated by ent, DO NOT EDIT.

package issuerpublickeys

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the issuerpublickeys type in the database.
	Label = "issuer_public_keys"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeSubjectPublicKey holds the string denoting the subject_public_key edge name in mutations.
	EdgeSubjectPublicKey = "subject_public_key"
	// Table holds the table name of the issuerpublickeys in the database.
	Table = "issuer_public_keys"
	// SubjectPublicKeyTable is the table that holds the subject_public_key relation/edge.
	SubjectPublicKeyTable = "subject_public_keys"
	// SubjectPublicKeyInverseTable is the table name for the SubjectPublicKeys entity.
	// It exists in this package in order to avoid circular dependency with the "subjectpublickeys" package.
	SubjectPublicKeyInverseTable = "subject_public_keys"
	// SubjectPublicKeyColumn is the table column denoting the subject_public_key relation/edge.
	SubjectPublicKeyColumn = "issuer_public_keys_subject_public_key"
)

// Columns holds all SQL columns for issuerpublickeys fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the IssuerPublicKeys queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// BySubjectPublicKeyField orders the results by subject_public_key field.
func BySubjectPublicKeyField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSubjectPublicKeyStep(), sql.OrderByField(field, opts...))
	}
}
func newSubjectPublicKeyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SubjectPublicKeyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, SubjectPublicKeyTable, SubjectPublicKeyColumn),
	)
}
