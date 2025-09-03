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
		case []string:
			for _, s := range v {
				values.Add(key, s)
			}
			continue
		case *[]string:
			if v == nil {
				continue
			}
			for _, s := range *v {
				values.Add(key, s)
			}
			continue
		case []any:
			tmp := make([]string, 0, len(v))
			ok := true
			for _, e := range v {
				switch s := e.(type) {
				case string:
					tmp = append(tmp, s)
				case *string:
					if s != nil {
						tmp = append(tmp, *s)
					} else {
						ok = false
					}
				default:
					ok = false
				}
				if !ok {
					break
				}
			}
			if ok {
				for _, s := range tmp {
					values.Add(key, s)
				}
			}
			continue
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

		case float32:
			stringValue = strconv.FormatFloat(float64(v), 'f', -1, 32)
		case *float32:
			if v == nil {
				continue
			}
			stringValue = strconv.FormatFloat(float64(*v), 'f', -1, 32)

		default:
			continue
		}

		values.Set(key, stringValue)
	}

}
