package main

import "fmt"

const helloEng = "Hello"

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}

	prefix := makePrefix(lang)

	return fmt.Sprintf("%s %s!", prefix, name)
}

func makePrefix(lang string) string {
	var prefix string
	switch lang {
	case "Spanish":
		prefix = "Hola"
	case "French":
		prefix = "Salut"
	default:
		prefix = helloEng
	}
	return prefix
}

func main() {
	fmt.Println(Hello("Piotr", ""))
}
