package dto

type (
	UserPathRequest struct {
		ID int64 `uri:"id"`
	}

	AddUserRequest struct {
		Name  string `json:"name" validate:"required"`
		Email string `json:"email" validate:"required,email"`
		Age   int    `json:"age" validate:"required"`
	}

	GetUserRequest struct {
		ID int64 `uri:"id"`
	}

	UpdateUserRequest struct {
		Name  *string `json:"name,omitempty"`
		Email *string `json:"email,omitempty"`
		Age   *int    `json:"age,omitempty"`
	}

	DeleteUserRequest struct {
		ID int64 `uri:"id"`
	}
)

type (
	AddUserResponse struct {
		ID int64 `json:"id"`
	}

	GetUserResponse struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Age   int    `json:"age"`
	}

	UpdateUserResponse struct {
		ID    int64  `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
		Age   int    `json:"age"`
	}

	DeleteUserResponse struct {
		BaseResponse
	}
)
