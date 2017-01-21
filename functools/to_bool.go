package functools

import (
	"reflect"
	"errors"
)

/*
'ToBool' function returns 'true' if numeric 'value' parameter isn't equal to zero,
string or iterable collections aren't empty, and for bool 'value' parameter it returns
original value.

	ToBool(value) bool
	ToBoolSafe(value) (bool, err)
*/

func toBool(value interface{}) bool {
	if value == nil {
		return false
	}

	switch value.(type) {
	case int:
		return value.(int) != 0
	case int8:
		return value.(int8) != 0
	case int16:
		return value.(int16) != 0
	case int32:
		return value.(int32) != 0
	case int64:
		return value.(int64) != 0
	case uint:
		return value.(uint) != 0
	case uint8:
		return value.(uint8) != 0
	case uint16:
		return value.(uint16) != 0
	case uint32:
		return value.(uint32) != 0
	case uint64:
		return value.(uint64) != 0
	case float32:
		return value.(float32) != 0.0
	case float64:
		return value.(float64) != 0.0
	case string:
		return value.(string) != ""
	case bool:
		return value.(bool)
	default:
		r := reflect.TypeOf(value)

		if r.Kind() == reflect.Array || r.Kind() == reflect.Slice || r.Kind() == reflect.Map {
			return reflect.ValueOf(value).Len() > 0
		}

		raise(errors.New("Unexpected type (" + r.String() + ") of value"), "ToBool")
	}
	return false
}

func ToBool(value interface{}) bool {
	return toBool(value)
}

func ToBoolSafe(value interface{}) (result bool, err error) {
	defer except(&err)
	result = toBool(value)
	return
}
