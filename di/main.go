package main

import (
	"di/greeter"
	"log"
	"net/http"
	"os"
)

func greetHandler(rw http.ResponseWriter, req *http.Request) {
	greeter.Greet(rw, "Piotr in a web")
}

func main() {
	greeter.Greet(os.Stdout, "Piotr")
	http.HandleFunc("/", greetHandler)
	err := http.ListenAndServe(":5001", nil)
	if err != nil {
		log.Fatalf("init error %q", err)
	}

}
