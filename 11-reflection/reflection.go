package reflection

import "reflect"

func Walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Struct:
		// NumField which returns the number of fields in the value.
		for i := range val.NumField() {
			Walk(val.Field(i).Interface(), fn)
		}
	case reflect.Slice, reflect.Array:
		// Len returns the number of elements in the slice.
		for i := range val.Len() {
			Walk(val.Index(i).Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}
}

func getValue(x any) reflect.Value {
	val := reflect.ValueOf(x)
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}
	return val
}
