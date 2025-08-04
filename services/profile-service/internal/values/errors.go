package values

import (
	"errors"
)

var (
	ErrNotAPointer  = errors.New("struct_is_not_a_pointer")
	ErrEmailExists  = errors.New("email already exists")
	ErrUserNotFound = errors.New("user not found")
)
