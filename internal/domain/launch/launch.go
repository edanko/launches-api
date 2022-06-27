package launch

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Launch struct {
	id          uuid.UUID
	name        string
	createdAt   time.Time
	updatedAt   time.Time
	completedAt *time.Time
	description *string
	reason      string
	status      Status
	files       *[]string
}

func (k *Launch) ID() uuid.UUID {
	return k.id
}

func (k *Launch) CreatedAt() time.Time {
	return k.createdAt
}

func (k *Launch) UpdatedAt() time.Time {
	return k.updatedAt
}

func (k *Launch) Name() string {
	return k.name
}

func (k *Launch) Description() *string {
	return k.description
}

func (k *Launch) Status() Status {
	return k.status
}

func NewLaunch(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	// name string,
	description *string,
	status string,
) (*Launch, error) {
	if id == uuid.Nil {
		return nil, ErrLaunchValidateID
	}
	// if name == "" {
	// 	return nil, ErrLaunchValidateName
	// }
	if description != nil && *description == "" {
		return nil, ErrLaunchValidateDescription
	}

	s, err := NewStatusFromString(status)
	if err != nil {
		return nil, errors.Wrap(err, ErrLaunchValidateStatus.Error())
	}

	launch := &Launch{
		id:        id,
		createdAt: createdAt,
		updatedAt: updatedAt,
		// name:        name,
		description: description,
		status:      s,
	}

	return launch, nil
}

func UnmarshalFromDB(
	id uuid.UUID,
	createdAt time.Time,
	updatedAt time.Time,
	// name string,
	description *string,
	status string,
) *Launch {
	launch := &Launch{
		id:        id,
		createdAt: createdAt,
		updatedAt: updatedAt,
		// name:        name,
		description: description,
		status:      MustNewStatusFromString(status),
	}

	return launch
}

func (k *Launch) ChangeName(newName string) {
	k.name = newName
}

func (k *Launch) ChangeDescription(newDescription string) {
	k.description = &newDescription
}

func (k *Launch) IsPublished() bool {
	return k.status == Todo
}

func (k *Launch) IsDraft() bool {
	return k.status == Started
}

func (k *Launch) MakePublished() error {
	if k.IsPublished() {
		return ErrLaunchAlreadyPublished
	}

	k.status = Todo
	return nil
}

func (k *Launch) MakeDraft() error {
	if k.IsDraft() {
		return ErrLaunchAlreadyDraft
	}

	k.status = Started
	return nil
}
