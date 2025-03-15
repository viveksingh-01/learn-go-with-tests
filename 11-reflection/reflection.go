package reflection

import "reflect"

func Walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	// val has a method NumField which returns the number of fields in the value.
	// This lets us iterate over the fields and call fn which passes our test.
	for i := range val.NumField() {
		// val.Field(i) returns the i-th field of the value.
		// We can then call String() on the field to get the string value.
		field := val.Field(i)
		// Kind helps us to check the type of the field
		if field.Kind() == reflect.String {
			fn(field.String())
		}
		// We inspect its Kind and if it happens to be a struct we just call walk again on that inner struct.
		if field.Kind() == reflect.Struct {
			Walk(field.Interface(), fn)
		}
	}
}
