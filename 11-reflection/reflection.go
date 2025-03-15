package reflection

import "reflect"

func Walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	// If the value is a pointer, we need to get the value it points to.
	// This is because we can't call NumField on a pointer.
	// We can use Elem() to get the value the pointer points to.
	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	// val has a method NumField which returns the number of fields in the value.
	// This lets us iterate over the fields and call fn which passes our test.
	for i := range val.NumField() {
		// val.Field(i) returns the i-th field of the value.
		// We can then call String() on the field to get the string value.
		field := val.Field(i)
		// Kind helps us to check the type of the field
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			Walk(field.Interface(), fn)
		}
	}
}
