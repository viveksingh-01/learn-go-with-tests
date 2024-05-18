package main

import "fmt"

func Hello(name string) string {
	if name == "" {
		return "Hello, world"
	}
	return "Hello, " + name
}

func main() {
	user := "Mike";
	fmt.Print(Hello(user))
}