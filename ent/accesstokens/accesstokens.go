// Code generated by ent, DO NOT EDIT.

package accesstokens

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the accesstokens type in the database.
	Label = "access_tokens"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// EdgeRequestID holds the string denoting the request_id edge name in mutations.
	EdgeRequestID = "request_id"
	// Table holds the table name of the accesstokens in the database.
	Table = "access_tokens"
	// RequestIDTable is the table that holds the request_id relation/edge.
	RequestIDTable = "access_tokens"
	// RequestIDInverseTable is the table name for the Request entity.
	// It exists in this package in order to avoid circular dependency with the "request" package.
	RequestIDInverseTable = "requests"
	// RequestIDColumn is the table column denoting the request_id relation/edge.
	RequestIDColumn = "request_access_token"
)

// Columns holds all SQL columns for accesstokens fields.
var Columns = []string{
	FieldID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "access_tokens"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"request_access_token",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the AccessTokens queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByRequestIDField orders the results by request_id field.
func ByRequestIDField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRequestIDStep(), sql.OrderByField(field, opts...))
	}
}
func newRequestIDStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RequestIDInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, RequestIDTable, RequestIDColumn),
	)
}