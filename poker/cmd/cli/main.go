package main

import (
	"fmt"
	"github.com/piojablonski/learn-go-with-tests/poker/store"
)

func main() {
	fmt.Println("poker command line utility")
}

func RecordWin(store store.PlayerStore, name string) {
	store.RecordWin(name)
}
