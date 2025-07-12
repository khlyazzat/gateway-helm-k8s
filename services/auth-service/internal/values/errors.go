package values

import (
	"errors"
)

var (
	// ErrWrongLoginOrPassword = NewHTTPError(http.StatusBadRequest, "api.invalid_login_or_password")
	// ErrSetRefreshToken         = NewHTTPError(http.StatusBadRequest, "api.set_refresh_token")
	// ErrParseToken = NewHTTPError(http.StatusBadRequest, "api.parse_token")
	// ErrGetRefreshToken = NewHTTPError(http.StatusBadRequest, "api.cannot_get_refresh_token")

	ErrWrongLoginOrPassword = errors.New("invalid_login_or_password")
	ErrSetRefreshToken      = errors.New("set_refresh_token")
	ErrParseToken           = errors.New("parse_token")
	ErrGetRefreshToken      = errors.New("cannot_get_refresh_token")

	ErrNotAPointer          = errors.New("struct_is_not_a_pointer")
	ErrEmailExists          = errors.New("email already exists")
	ErrUserNotFound         = errors.New("user not found")
	ErrInvalidToken         = errors.New("invalid_token")
	ErrInvalidClaims        = errors.New("invalid_claims")
	ErrInvalidSigningMethod = errors.New("invalid_signing_method")
	ErrInvalidSigningToken  = errors.New("invalid_signing_token")
	ErrRedisKeyExists       = errors.New("key_does_not_exist")
	// ErrSelectFilter         = errors.New("db.select_query_filter_error")
)
