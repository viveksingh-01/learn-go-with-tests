package main

import "fmt"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return getGreetingPrefix(language) + name
}

func getGreetingPrefix(language string) (prefix string) {
	switch language {
	case "French":
		prefix = "Bonjour, "
	case "German":
		prefix = "Hallo, "
	case "Spanish":
		prefix = "Hola, "
	default:
		prefix = "Hello, "
	}
	return
}

func main() {
	user := "Mike"
	fmt.Print(Hello(user, ""))
}
