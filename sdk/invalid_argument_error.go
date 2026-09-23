package sdk

import "fmt"

// InvalidArgumentError Returned before a request is sent when an argument passed to an operation
// cannot be used. It reports the name of the parameter that was rejected, the value it was given,
// and why that value is not accepted. Its code is ErrorCode_REQUEST_CANNOT_BE_PARSED, since the
// request could not be formed from the given arguments.
type InvalidArgumentError struct {
	error     string
	parameter string
	value     string
	code      ErrorCode
}

func newInvalidArgumentError(parameter string, value string, message string) *InvalidArgumentError {
	return &InvalidArgumentError{
		error:     fmt.Sprintf("invalid value %q for %s: %s", value, parameter, message),
		parameter: parameter,
		value:     value,
		code:      ErrorCode_REQUEST_CANNOT_BE_PARSED,
	}
}

func (e *InvalidArgumentError) Error() string {
	return e.error
}

func (e *InvalidArgumentError) Body() []byte {
	return nil
}

func (e *InvalidArgumentError) Model() any {
	return nil
}

// Parameter returns the name of the parameter that was rejected.
func (e *InvalidArgumentError) Parameter() string {
	return e.parameter
}

// Value returns the rejected value.
func (e *InvalidArgumentError) Value() string {
	return e.value
}

func (e *InvalidArgumentError) Code() ErrorCode {
	return e.code
}
