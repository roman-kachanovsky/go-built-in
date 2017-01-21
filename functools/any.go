package functools

import (
	"reflect"
	"errors"
)

func any(function, slice interface{}) bool {
	in := reflect.ValueOf(slice)

	if in.Kind() != reflect.Slice {
		raise(errors.New("The passed collection is not a slice"), "Any")
	}

	if in.Len() == 0 {
		return false
	}

	fn := reflect.ValueOf(function)
	inType := in.Type().Elem()

	if !verifyAnyFuncType(fn, inType) {
		raise(errors.New("Function must be of type func(" + inType.String() +
			") bool or func(interface{}) bool"), "Any")
	}

	var param [1]reflect.Value

	for i := 0; i < in.Len(); i++ {
		param[0] = in.Index(i)

		if fn.Call(param[:])[0].Bool() {
			return true
		}
	}

	return false
}

func verifyAnyFuncType(fn reflect.Value, elType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}

	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 1 {
		return false
	}

	return ((fn.Type().In(0).Kind() == reflect.Interface || fn.Type().In(0) == elType) &&
		fn.Type().Out(0).Kind() == reflect.Bool)
}

func Any(slice interface{}) bool {
	return any(ToBool, slice)
}

func AnySafe(slice interface{}) (result bool, err error) {
	defer except(&err)
	result = any(ToBool, slice)
	return
}

func AnyFunc(function, slice interface{}) bool {
	return any(function, slice)
}

func AnyFuncSafe(function, slice interface{}) (result bool, err error) {
	defer except(&err)
	result = any(function, slice)
	return
}
