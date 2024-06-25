package sdk

import (
	"net/url"
	"strconv"
)

type urlValues interface {
	UrlValues() *url.Values
}

type urlValuesBuilder struct {
	data map[string]any
}

func (u *urlValuesBuilder) UrlValues() *url.Values {
	values := make(url.Values)

	for key, value := range u.data {
		var stringValue string

		switch v := value.(type) {
		case string:
			stringValue = v
		case int32:
			stringValue = strconv.FormatInt(int64(v), 10)
		case int64:
			stringValue = strconv.FormatInt(v, 10)
		default:
			continue
		}

		values.Set(key, stringValue)
	}

	return &values
}
