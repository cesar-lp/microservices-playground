package common

// FieldError structure
type FieldError struct {
	FieldName    string `json:"fieldName"`
	Error        string `json:"error"`
	InvalidValue string `json:"invalidValue"`
}

func NewFieldError(fieldName, err, invalidValue string) FieldError {
	return FieldError{
		FieldName:    fieldName,
		Error:        err,
		InvalidValue: invalidValue,
	}
}
