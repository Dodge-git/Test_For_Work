package service

import "errors"

var (
	ErrDepartmentNotFound = errors.New("department not found")
	ErrInvalidName        = errors.New("invalid name")
	ErrCycleDetected      = errors.New("cycle detected")
	ErrDepthExceeded      = errors.New("max depth is 5")
	ErrInvalidParent      = errors.New("invalid parent")
	ErrDuplicateName      = errors.New("duplicate name in same parent")
	ErrParentNotFound     = errors.New("parent department not found")
	ErrInvalidDeleteMode  = errors.New("invalid delete mode")
	ErrReassignRequired   = errors.New("reassign required")
	ErrInvalidFullName    = errors.New("invalid full name")
	ErrInvalidPosition    = errors.New("invalid position")
)
