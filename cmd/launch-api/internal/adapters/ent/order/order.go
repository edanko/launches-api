// Code generated by entc, DO NOT EDIT.

package order

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the order type in the database.
	Label = "order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// EdgeLaunches holds the string denoting the launches edge name in mutations.
	EdgeLaunches = "launches"
	// Table holds the table name of the order in the database.
	Table = "orders"
	// LaunchesTable is the table that holds the launches relation/edge.
	LaunchesTable = "launches"
	// LaunchesInverseTable is the table name for the Launch entity.
	// It exists in this package in order to avoid circular dependency with the "launch" package.
	LaunchesInverseTable = "launches"
	// LaunchesColumn is the table column denoting the launches relation/edge.
	LaunchesColumn = "order_launches"
)

// Columns holds all SQL columns for order fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldName,
	FieldDescription,
	FieldStatus,
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

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Status defines the type for the "status" enum field.
type Status string

// StatusDraft is the default value of the Status enum.
const DefaultStatus = StatusDraft

// Status values.
const (
	StatusPublished Status = "published"
	StatusDraft     Status = "draft"
)

func (s Status) String() string {
	return string(s)
}

// StatusValidator is a validator for the "status" field enum values. It is called by the builders before save.
func StatusValidator(s Status) error {
	switch s {
	case StatusPublished, StatusDraft:
		return nil
	default:
		return fmt.Errorf("order: invalid enum value for status field: %q", s)
	}
}
