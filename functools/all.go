package functools

import (
	"reflect"
	"errors"
)

func all(function, slice interface{}) bool {
	in := reflect.ValueOf(slice)

	if in.Kind() != reflect.Slice {
		raise(errors.New("The passed collection is not a slice"), "All")
	}

	if in.Len() == 0 {
		return false
	}

	fn := reflect.ValueOf(function)
	inType := in.Type().Elem()

	if !verifyAllFuncType(fn, inType) {
		raise(errors.New("Function must be of type func(" + inType.String() +
			") bool or func(interface{}) bool"), "All")
	}

	var param [1]reflect.Value

	for i := 0; i < in.Len(); i++ {
		param[0] = in.Index(i)

		if !fn.Call(param[:])[0].Bool() {
			return false
		}
	}

	return true
}

func verifyAllFuncType(fn reflect.Value, elType reflect.Type) bool {
	if fn.Kind() != reflect.Func {
		return false
	}

	if fn.Type().NumIn() != 1 || fn.Type().NumOut() != 1 {
		return false
	}

	return ((fn.Type().In(0).Kind() == reflect.Interface || fn.Type().In(0) == elType) &&
		fn.Type().Out(0).Kind() == reflect.Bool)
}

func All(slice interface{}) bool {
	return all(ToBool, slice)
}

func AllSafe(slice interface{}) (result bool, err error) {
	defer except(&err)
	result = all(ToBool, slice)
	return
}

func AllFunc(function, slice interface{}) bool {
	return all(function, slice)
}

func AllFuncSafe(function, slice interface{}) (result bool, err error) {
	defer except(&err)
	result = all(function, slice)
	return
}
