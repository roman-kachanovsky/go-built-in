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
	case int, int8, int16, int32, int64:
		return reflect.ValueOf(value).Int() != 0
	case uint, uint8, uint16, uint32, uint64:
		return reflect.ValueOf(value).Uint() != 0
	case float32, float64:
		return reflect.ValueOf(value).Float() != 0.0
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
