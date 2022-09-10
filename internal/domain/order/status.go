package order

import (
	"github.com/pkg/errors"
)

var (
	Published = Status{"published"}
	Draft     = Status{"draft"}
)

var statusValues = map[string]Status{
	"published": Published,
	"draft":     Draft,
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
