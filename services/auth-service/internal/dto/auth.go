package dto

type (
	SignUpRequest struct {
		Email    string
		Password string
		Name     string
		Age      int
	}

	SignInRequest struct {
		Email    string
		Password string
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
