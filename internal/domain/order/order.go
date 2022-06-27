package order

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Order struct {
	id          uuid.UUID
	name        string
	createdAt   time.Time
	updatedAt   time.Time
	description *string
	status      Status
}

func (k *Order) ID() uuid.UUID {
	return k.id
}

func (k *Order) CreatedAt() time.Time {
	return k.createdAt
}

func (k *Order) UpdatedAt() time.Time {
	return k.updatedAt
}

func (k *Order) Name() string {
	return k.name
}

func (k *Order) Description() *string {
	return k.description
}

func (k *Order) Status() Status {
	return k.status
}

func NewOrder(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	name string,
	description *string,
	status string,
) (*Order, error) {
	if id == uuid.Nil {
		return nil, ErrOrderValidateID
	}
	if name == "" {
		return nil, ErrOrderValidateName
	}
	if description != nil && *description == "" {
		return nil, ErrOrderValidateDescription
	}

	s, err := NewStatusFromString(status)
	if err != nil {
		return nil, errors.Wrap(err, ErrOrderValidateStatus.Error())
	}

	order := &Order{
		id:          id,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
		name:        name,
		description: description,
		status:      s,
	}

	return order, nil
}

func UnmarshalFromDB(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	name string,
	description *string,
	status string,
) *Order {
	order := &Order{
		id:          id,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
		name:        name,
		description: description,
		status:      MustNewStatusFromString(status),
	}

	return order
}

func (k *Order) ChangeName(newName string) {
	k.name = newName
}

func (k *Order) ChangeDescription(newDescription string) {
	k.description = &newDescription
}

func (k *Order) IsPublished() bool {
	return k.status == Published
}

func (k *Order) IsDraft() bool {
	return k.status == Draft
}

func (k *Order) MakePublished() error {
	if k.IsPublished() {
		return ErrOrderAlreadyPublished
	}

	k.status = Published
	return nil
}

func (k *Order) MakeDraft() error {
	if k.IsDraft() {
		return ErrOrderAlreadyDraft
	}

	k.status = Draft
	return nil
}
