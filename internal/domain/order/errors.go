package order

import (
	"errors"
)

var (
	ErrOrderNotFound            = errors.New("order not found")
	ErrOrderValidateID          = errors.New("order uuid validation failed")
	ErrOrderValidateName        = errors.New("order name cannot be empty")
	ErrOrderValidateDescription = errors.New("order description cannot be empty")
	ErrOrderValidateStatus      = errors.New("order status validation failed")
	ErrOrderAlreadyExist        = errors.New("order is already exist")
	ErrOrderAlreadyPublished    = errors.New("order is already published")
	ErrOrderAlreadyDraft        = errors.New("order is already draft")
)
