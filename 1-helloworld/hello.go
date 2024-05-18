package main

import "fmt"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	if language == "Spanish" {
		return "Hola, " + name
	}
	return "Hello, " + name
}

func main() {
	user := "Mike"
	fmt.Print(Hello(user, ""))
}
