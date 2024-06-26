package sdk

import (
	"net/url"
	"strconv"
)

func addMapToUrlValues(data map[string]any, values *url.Values) {
	for key, value := range data {
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

}
