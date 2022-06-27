// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/edanko/nx/cmd/launch-api/internal/adapters/ent/applicant"
	"github.com/edanko/nx/cmd/launch-api/internal/adapters/ent/kind"
	"github.com/edanko/nx/cmd/launch-api/internal/adapters/ent/launch"
	"github.com/edanko/nx/cmd/launch-api/internal/adapters/ent/order"
	"github.com/google/uuid"
)

// Launch is the model entity for the Launch schema.
type Launch struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status holds the value of the "status" field.
	Status launch.Status `json:"status,omitempty"`
	// Reason holds the value of the "reason" field.
	Reason string `json:"reason,omitempty"`
	// Description holds the value of the "description" field.
	Description *string `json:"description,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LaunchQuery when eager-loading is set.
	Edges              LaunchEdges `json:"edges"`
	applicant_launches *uuid.UUID
	kind_launches      *uuid.UUID
	order_launches     *uuid.UUID
}

// LaunchEdges holds the relations/edges for other nodes in the graph.
type LaunchEdges struct {
	// Order holds the value of the order edge.
	Order *Order `json:"order,omitempty"`
	// Kind holds the value of the kind edge.
	Kind *Kind `json:"kind,omitempty"`
	// Applicant holds the value of the applicant edge.
	Applicant *Applicant `json:"applicant,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LaunchEdges) OrderOrErr() (*Order, error) {
	if e.loadedTypes[0] {
		if e.Order == nil {
			// The edge order was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: order.Label}
		}
		return e.Order, nil
	}
	return nil, &NotLoadedError{edge: "order"}
}

// KindOrErr returns the Kind value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LaunchEdges) KindOrErr() (*Kind, error) {
	if e.loadedTypes[1] {
		if e.Kind == nil {
			// The edge kind was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: kind.Label}
		}
		return e.Kind, nil
	}
	return nil, &NotLoadedError{edge: "kind"}
}

// ApplicantOrErr returns the Applicant value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LaunchEdges) ApplicantOrErr() (*Applicant, error) {
	if e.loadedTypes[2] {
		if e.Applicant == nil {
			// The edge applicant was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: applicant.Label}
		}
		return e.Applicant, nil
	}
	return nil, &NotLoadedError{edge: "applicant"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Launch) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case launch.FieldStatus, launch.FieldReason, launch.FieldDescription:
			values[i] = new(sql.NullString)
		case launch.FieldCreatedAt, launch.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case launch.FieldID:
			values[i] = new(uuid.UUID)
		case launch.ForeignKeys[0]: // applicant_launches
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case launch.ForeignKeys[1]: // kind_launches
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case launch.ForeignKeys[2]: // order_launches
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Launch", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Launch fields.
func (l *Launch) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case launch.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				l.ID = *value
			}
		case launch.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				l.CreatedAt = value.Time
			}
		case launch.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				l.UpdatedAt = value.Time
			}
		case launch.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				l.Status = launch.Status(value.String)
			}
		case launch.FieldReason:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field reason", values[i])
			} else if value.Valid {
				l.Reason = value.String
			}
		case launch.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				l.Description = new(string)
				*l.Description = value.String
			}
		case launch.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field applicant_launches", values[i])
			} else if value.Valid {
				l.applicant_launches = new(uuid.UUID)
				*l.applicant_launches = *value.S.(*uuid.UUID)
			}
		case launch.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field kind_launches", values[i])
			} else if value.Valid {
				l.kind_launches = new(uuid.UUID)
				*l.kind_launches = *value.S.(*uuid.UUID)
			}
		case launch.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field order_launches", values[i])
			} else if value.Valid {
				l.order_launches = new(uuid.UUID)
				*l.order_launches = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryOrder queries the "order" edge of the Launch entity.
func (l *Launch) QueryOrder() *OrderQuery {
	return (&LaunchClient{config: l.config}).QueryOrder(l)
}

// QueryKind queries the "kind" edge of the Launch entity.
func (l *Launch) QueryKind() *KindQuery {
	return (&LaunchClient{config: l.config}).QueryKind(l)
}

// QueryApplicant queries the "applicant" edge of the Launch entity.
func (l *Launch) QueryApplicant() *ApplicantQuery {
	return (&LaunchClient{config: l.config}).QueryApplicant(l)
}

// Update returns a builder for updating this Launch.
// Note that you need to call Launch.Unwrap() before calling this method if this Launch
// was returned from a transaction, and the transaction was committed or rolled back.
func (l *Launch) Update() *LaunchUpdateOne {
	return (&LaunchClient{config: l.config}).UpdateOne(l)
}

// Unwrap unwraps the Launch entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (l *Launch) Unwrap() *Launch {
	tx, ok := l.config.driver.(*txDriver)
	if !ok {
		panic("ent: Launch is not a transactional entity")
	}
	l.config.driver = tx.drv
	return l
}

// String implements the fmt.Stringer.
func (l *Launch) String() string {
	var builder strings.Builder
	builder.WriteString("Launch(")
	builder.WriteString(fmt.Sprintf("id=%v", l.ID))
	builder.WriteString(", created_at=")
	builder.WriteString(l.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(l.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", l.Status))
	builder.WriteString(", reason=")
	builder.WriteString(l.Reason)
	if v := l.Description; v != nil {
		builder.WriteString(", description=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Launches is a parsable slice of Launch.
type Launches []*Launch

func (l Launches) config(cfg config) {
	for _i := range l {
		l[_i].config = cfg
	}
}
