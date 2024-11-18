package sdk

import "errors"

// Error defines base interface of all errors returned by this SDK
type Error interface {
	error

	// Body returns the raw bytes of the response, if available.
	Body() []byte

	// Code returns the error code.
	Code() ErrorCode

	// Model returns the unpacked model of the error. When error was created with WrapWithApiError, it returns the original error.
	// If error was thrown after we get the HTTP Response, it contains model parsed from response body.
	Model() any
}

// ApiError Provides access to the body, error and model on returned errors.
type ApiError struct {
	body  []byte
	error string
	model any
	code  ErrorCode
}

func WrapWithApiError(err error) *ApiError {
	if err == nil {
		return nil
	}

	var apiError *ApiError
	if errors.As(err, &apiError) {
		return apiError
	}

	return &ApiError{
		error: err.Error(),
		code:  FAILED,
		model: err,
	}
}

// Error returns non-empty string if there was an error.
func (e ApiError) Error() string {
	return e.error
}

func (e ApiError) Body() []byte {
	return e.body
}

func (e ApiError) Model() any {
	return e.model
}

func (e ApiError) Code() ErrorCode {
	return e.code
}
