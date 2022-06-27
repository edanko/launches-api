// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/edanko/nx/cmd/launch-api/internal/adapters/repositories/ent/kind"
	"github.com/google/uuid"
)

// Kind is the model entity for the Kind schema.
type Kind struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// Status holds the value of the "status" field.
	Status kind.Status `json:"status,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Kind) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case kind.FieldName, kind.FieldDescription, kind.FieldStatus:
			values[i] = new(sql.NullString)
		case kind.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Kind", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Kind fields.
func (k *Kind) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case kind.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				k.ID = *value
			}
		case kind.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				k.Name = value.String
			}
		case kind.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				k.Description = new(string)
				*k.Description = value.String
			}
		case kind.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				k.Status = kind.Status(value.String)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Kind.
// Note that you need to call Kind.Unwrap() before calling this method if this Kind
// was returned from a transaction, and the transaction was committed or rolled back.
func (k *Kind) Update() *KindUpdateOne {
	return (&KindClient{config: k.config}).UpdateOne(k)
}

// Unwrap unwraps the Kind entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (k *Kind) Unwrap() *Kind {
	tx, ok := k.config.driver.(*txDriver)
	if !ok {
		panic("ent: Kind is not a transactional entity")
	}
	k.config.driver = tx.drv
	return k
}

// String implements the fmt.Stringer.
func (k *Kind) String() string {
	var builder strings.Builder
	builder.WriteString("Kind(")
	builder.WriteString(fmt.Sprintf("id=%v", k.ID))
	builder.WriteString(", name=")
	builder.WriteString(k.Name)
	if v := k.Description; v != nil {
		builder.WriteString(", description=")
		builder.WriteString(*v)
	}
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", k.Status))
	builder.WriteByte(')')
	return builder.String()
}

// Kinds is a parsable slice of Kind.
type Kinds []*Kind

func (k Kinds) config(cfg config) {
	for _i := range k {
		k[_i].config = cfg
	}
}
