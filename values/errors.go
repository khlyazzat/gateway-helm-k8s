package values

import "errors"

var (
	ErrInvalidSigningToken  = errors.New("invalid_signing_token")
	ErrInvalidSigningMethod = errors.New("invalid_signing_method")
	ErrInvalidToken         = errors.New("invalid_token")
	ErrInvalidClaims        = errors.New("invalid_claims")
)
