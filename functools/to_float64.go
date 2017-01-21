package functools

import (
	"reflect"
	"errors"
)

/*
'ToFloat64' function converts numeric 'value' to float64 '(-)value.0',
'true' to 1.0 and 'false' to 0.0, 'nil' to 0.0 and raised exception for other types.

	ToFloat64(value) float64
	ToFloat64Safe(value) (float64, err)
*/

func toFloat64(value interface{}) float64 {
	if value == nil {
		return 0.0
	}

	switch value.(type) {
	case int:
		return float64(value.(int))
	case int8:
		return float64(value.(int8))
	case int16:
		return float64(value.(int16))
	case int32:
		return float64(value.(int32))
	case int64:
		return float64(value.(int64))
	case uint:
		return float64(value.(uint))
	case uint8:
		return float64(value.(uint8))
	case uint16:
		return float64(value.(uint16))
	case uint32:
		return float64(value.(uint32))
	case uint64:
		return float64(value.(uint64))
	case float32:
		return float64(value.(float32))
	case float64:
		return value.(float64)
	case bool:
		if value.(bool) {
			return 1.0
		}
		return 0.0
	default:
		r := reflect.TypeOf(value)
		raise(errors.New("ToFloat64 can't cast "+ r.String() +" value to float64"), "ToFloat64")
	}
	return 0.0
}

func ToFloat64(value interface{}) float64 {
	return toFloat64(value)
}

func ToFloat64Safe(value interface{}) (result float64, err error) {
	defer except(&err)
	result = toFloat64(value)
	return
}
