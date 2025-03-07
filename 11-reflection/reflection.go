package reflection

import "reflect"

// Go challenge: write a function which takes a struct x and calls fn for all strings fields found inside.
// We may come across scenarios though where we want to write a function where we don't know
// the type at compile time.
// Go lets us get around this with the type interface{} which we can think of as just any type.
// We can use the reflect package to inspect the type of a variable at runtime.

// Walk walks through the fields of a struct and calls a callback function for each field.
// The callback function is expected to take a string as input.
func Walk(x interface{}, fn func(input string)) {
	// The reflect package has a function ValueOf which returns us a Value of a given variable.
	// This has ways for us to inspect a value, including its fields
	val := reflect.ValueOf(x)

	// val has a method NumField which returns the number of fields in the value.
	// This lets us iterate over the fields and call fn which passes our test.
	for i := 0; i < val.NumField(); i++ {
		// val.Field(i) returns the i-th field of the value.
		// We can then call String() on the field to get the string value.
		field := val.Field(i)
		fn(field.String())
	}
}
