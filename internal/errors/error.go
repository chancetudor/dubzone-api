package errors

// Error is the type that implements the error interface.
type Error struct {
	Message    string   `json:"message"`
	Code       int      `json:"code"`
	Name       string   `json:"name"`
	Error      error    `json:"-"`
	Validation []string `json:"validation,omitempty"`
}

func BindError() *Error {
	return &Error{Message: "Error processing request.", Code: 400, Name: "BIND_ERROR"}
}

func ValidationError(errors []string) *Error {
	return &Error{Message: "A validation error occurred.", Code: 400, Name: "VALIDATION_ERROR", Validation: errors}
}
