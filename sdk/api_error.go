package sdk

// ApiError Provides access to the body, error and model on returned errors.
type ApiError struct {
	body  []byte
	error string
	model interface{}
}

// Error returns non-empty string if there was an error.
func (e ApiError) Error() string {
	return e.error
}

// Body returns the raw bytes of the response
func (e ApiError) Body() []byte {
	return e.body
}

// Model returns the unpacked model of the error
func (e ApiError) Model() interface{} {
	return e.model
}
