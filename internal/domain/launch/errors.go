package launch

import (
	"errors"
)

var (
	ErrLaunchNotFound            = errors.New("launch not found")
	ErrLaunchValidateID          = errors.New("launch uuid validation failed")
	ErrLaunchValidateName        = errors.New("launch name cannot be empty")
	ErrLaunchValidateDescription = errors.New("launch description cannot be empty")
	ErrLaunchValidateStatus      = errors.New("launch status validation failed")
	ErrLaunchAlreadyExist        = errors.New("launch is already exist")
	ErrLaunchAlreadyPublished    = errors.New("launch is already published")
	ErrLaunchAlreadyDraft        = errors.New("launch is already draft")
)
