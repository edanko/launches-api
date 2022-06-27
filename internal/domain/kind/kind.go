package kind

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Kind struct {
	id          uuid.UUID
	name        string
	createdAt   time.Time
	updatedAt   time.Time
	description *string
	status      Status
}

func (k *Kind) ID() uuid.UUID {
	return k.id
}

func (k *Kind) CreatedAt() time.Time {
	return k.createdAt
}

func (k *Kind) UpdatedAt() time.Time {
	return k.updatedAt
}

func (k *Kind) Name() string {
	return k.name
}

func (k *Kind) Description() *string {
	return k.description
}

func (k *Kind) Status() Status {
	return k.status
}

func NewKind(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	name string,
	description *string,
	status string,
) (*Kind, error) {
	if id == uuid.Nil {
		return nil, ErrKindValidateID
	}
	if name == "" {
		return nil, ErrKindValidateName
	}
	if description != nil && *description == "" {
		return nil, ErrKindValidateDescription
	}

	s, err := NewStatusFromString(status)
	if err != nil {
		return nil, errors.Wrap(err, ErrKindValidateStatus.Error())
	}

	kind := &Kind{
		id:          id,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
		name:        name,
		description: description,
		status:      s,
	}

	return kind, nil
}

func UnmarshalFromDB(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	name string,
	description *string,
	status string,
) *Kind {
	kind := &Kind{
		id:          id,
		createdAt:   createdAt,
		updatedAt:   updatedAt,
		name:        name,
		description: description,
		status:      statusValues[status],
	}

	return kind
}

func (k *Kind) ChangeName(newName string) {
	k.name = newName
}

func (k *Kind) ChangeDescription(newDescription string) {
	k.description = &newDescription
}

func (k *Kind) IsPublished() bool {
	return k.status == Published
}

func (k *Kind) IsDraft() bool {
	return k.status == Draft
}

func (k *Kind) MakePublished() error {
	if k.IsPublished() {
		return ErrKindAlreadyPublished
	}

	k.status = Published
	return nil
}

func (k *Kind) MakeDraft() error {
	if k.IsDraft() {
		return ErrKindAlreadyDraft
	}

	k.status = Draft
	return nil
}
