package sdk

// TooManyRequestsError Provides access to the body, error and model on returned 429 TooManyRequests error
type TooManyRequestsError struct {
	error      string
	retryAfter int64
	body       []byte
	model      interface{}
	code       string
}

func (e *TooManyRequestsError) Model() interface{} {
	return e.model
}

func (e *TooManyRequestsError) Error() string {
	return e.error
}

func (e *TooManyRequestsError) Body() []byte {
	return e.body
}

func (e *TooManyRequestsError) RetryAfter() int64 {
	return e.retryAfter
}

func (e *TooManyRequestsError) Code() string {
	return e.code
}
