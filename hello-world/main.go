package main

import "fmt"

const helloEng = "Hello"

func Hello(name string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s %s!", helloEng, name)
}

func main() {
	fmt.Println(Hello("Piotr"))
}
