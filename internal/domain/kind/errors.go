package kind

import (
	"errors"
)

var (
	ErrKindNotFound            = errors.New("kind not found")
	ErrKindValidateID          = errors.New("kind uuid validation failed")
	ErrKindValidateName        = errors.New("kind name cannot be empty")
	ErrKindValidateDescription = errors.New("kind description cannot be empty")
	ErrKindValidateStatus      = errors.New("kind status validation failed")
	ErrKindAlreadyExist        = errors.New("kind is already exist")
	ErrKindAlreadyPublished    = errors.New("kind is already published")
	ErrKindAlreadyDraft        = errors.New("kind is already draft")
)
