package sdk

import (
	"net/url"
	"reflect"
	"strconv"
)

func addStructToURLQuery(query *url.Values, opts interface{}) {
	v := reflect.ValueOf(opts)
	t := v.Type()

	if t.Kind() == reflect.Pointer {
		v = v.Elem()
		t = v.Type()
	}

	// Ensure the provided value is a struct
	if t.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()

		// Get the query tag
		tag := field.Tag.Get("url")
		if tag == "" {
			tag = field.Name
		}

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

		if stringValue != "" {
			query.Add(tag, stringValue)
		}
	}
}
