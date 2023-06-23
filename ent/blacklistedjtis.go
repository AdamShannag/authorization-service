// Code generated by ent, DO NOT EDIT.

package ent

import (
	"authorization-service/ent/blacklistedjtis"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// BlacklistedJTIs is the model entity for the BlacklistedJTIs schema.
type BlacklistedJTIs struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Expiry holds the value of the "expiry" field.
	Expiry       time.Time `json:"expiry,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BlacklistedJTIs) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case blacklistedjtis.FieldID:
			values[i] = new(sql.NullString)
		case blacklistedjtis.FieldExpiry:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BlacklistedJTIs fields.
func (bji *BlacklistedJTIs) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case blacklistedjtis.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				bji.ID = value.String
			}
		case blacklistedjtis.FieldExpiry:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expiry", values[i])
			} else if value.Valid {
				bji.Expiry = value.Time
			}
		default:
			bji.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BlacklistedJTIs.
// This includes values selected through modifiers, order, etc.
func (bji *BlacklistedJTIs) Value(name string) (ent.Value, error) {
	return bji.selectValues.Get(name)
}

// Update returns a builder for updating this BlacklistedJTIs.
// Note that you need to call BlacklistedJTIs.Unwrap() before calling this method if this BlacklistedJTIs
// was returned from a transaction, and the transaction was committed or rolled back.
func (bji *BlacklistedJTIs) Update() *BlacklistedJTIsUpdateOne {
	return NewBlacklistedJTIsClient(bji.config).UpdateOne(bji)
}

// Unwrap unwraps the BlacklistedJTIs entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bji *BlacklistedJTIs) Unwrap() *BlacklistedJTIs {
	_tx, ok := bji.config.driver.(*txDriver)
	if !ok {
		panic("ent: BlacklistedJTIs is not a transactional entity")
	}
	bji.config.driver = _tx.drv
	return bji
}

// String implements the fmt.Stringer.
func (bji *BlacklistedJTIs) String() string {
	var builder strings.Builder
	builder.WriteString("BlacklistedJTIs(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bji.ID))
	builder.WriteString("expiry=")
	builder.WriteString(bji.Expiry.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BlacklistedJTIsSlice is a parsable slice of BlacklistedJTIs.
type BlacklistedJTIsSlice []*BlacklistedJTIs
