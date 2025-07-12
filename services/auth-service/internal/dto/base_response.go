package dto

type (
	BaseResponse struct {
		Status bool            `json:"status"`
		Errors []ErrorResponse `json:"errors,omitempty"`
	}

	ErrorResponse struct {
		Message string `json:"message"`
		Field   string `json:"field,omitempty"`
		Tag     string `json:"tag,omitempty"`
	}
)
