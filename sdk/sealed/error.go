package sealed

import (
	"fmt"
	"strings"
)

type AggregatedUnsealError struct {
	UnsealErrors []UnsealError
}

func (e *AggregatedUnsealError) joinErrorMessages() string {
	var errorList []string
	for _, err := range e.UnsealErrors {
		errorList = append(errorList, err.Error.Error())
	}

	return strings.Join(errorList, ", ")
}

func (e *AggregatedUnsealError) Error() string {
	return fmt.Sprintf("unseal failed: %s.", e.joinErrorMessages())
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
