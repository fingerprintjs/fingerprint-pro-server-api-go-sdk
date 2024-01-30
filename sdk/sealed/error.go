package sealed

type AggregatedUnsealError struct {
	UnsealErrors []UnsealError
}

func (e *AggregatedUnsealError) Error() string {
	return "unseal failed"
}

func NewAggregatedUnsealError() *AggregatedUnsealError {
	return &AggregatedUnsealError{}
}

func (e *AggregatedUnsealError) Add(error UnsealError) {
	e.UnsealErrors = append(e.UnsealErrors, error)
}

type UnsealError struct {
	Error error
	Key   DecryptionKey
}

func NewUnsealError(err error, key DecryptionKey) UnsealError {
	return UnsealError{Error: err, Key: key}
}
