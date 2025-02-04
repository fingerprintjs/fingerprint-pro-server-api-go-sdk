package sdk

import (
	"net/url"
	"strconv"
)

func addMapToUrlValues(data map[string]any, values *url.Values) {
	for key, value := range data {
		var stringValue string

		if value == nil {
			continue
		}

		switch v := value.(type) {
		case bool:
			stringValue = strconv.FormatBool(v)
		case *bool:
			if v == nil {
				continue
			}

			stringValue = strconv.FormatBool(*v)

		case *string:
			if v == nil {
				continue
			}

			stringValue = *v
		case string:
			stringValue = v

		case *int32:
			if v == nil {
				continue
			}

			stringValue = strconv.FormatInt(int64(*v), 10)
		case int32:
			stringValue = strconv.FormatInt(int64(v), 10)

		case *int64:
			if v == nil {
				continue
			}

			stringValue = strconv.FormatInt(*v, 10)
		case int64:
			stringValue = strconv.FormatInt(v, 10)

		default:
			continue
		}

		values.Set(key, stringValue)
	}

}
