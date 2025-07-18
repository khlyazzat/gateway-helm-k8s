package dto

type (
	SignUpRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
		Name     string `json:"name" validate:"required"`
		Age      int    `json:"age" validate:"required"`
	}
	SignInRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	RefreshRequest struct {
		Token string `json:"token" validate:"required"`
	}
)

type (
	SignUpResponse struct {
	}

	TokenResponse struct {
		AuthToken    string `json:"authToken"`
		RefreshToken string `json:"refreshToken"`
		TTL          int    `json:"ttl"`
	}
)
