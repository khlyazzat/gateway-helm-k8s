package dto

type (
	UpdateProfileRequest struct {
		FirstName *string `json:"firstName,omitempty"`
		LastName  *string `json:"lastName,omitempty"`
		Age       *int    `json:"age,omitempty"`
		Address   *string `json:"address,omitempty"`
		Phone     *string `json:"phone,omitempty"`
	}
)

type (
	GetProfileResponse struct {
		Email     string  `json:"email"`
		FirstName *string `json:"firstName,omitempty"`
		LastName  *string `json:"lastName,omitempty"`
		Age       *int    `json:"age,omitempty"`
		Address   *string `json:"address,omitempty"`
		Phone     *string `json:"phone,omitempty"`
	}

	UpdateProfileResponse struct {
	}
)
