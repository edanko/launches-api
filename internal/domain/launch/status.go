package launch

import (
	"github.com/pkg/errors"
)

var (
	Todo      = Status{"todo"}
	Started   = Status{"started"}
	Completed = Status{"completed"}
)

var statusValues = map[string]Status{
	"todo":      Todo,
	"started":   Started,
	"completed": Completed,
}

type Status struct {
	s string
}

func NewStatusFromString(s string) (Status, error) {
	if _, ok := statusValues[s]; !ok {
		return Status{}, errors.Errorf("unknown status value: %s", s)
	}
	return statusValues[s], nil
}

func MustNewStatusFromString(s string) Status {
	return statusValues[s]
}

func (h Status) IsZero() bool {
	return h == Status{}
}

func (h Status) String() string {
	return h.s
}
