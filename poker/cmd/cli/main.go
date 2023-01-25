package main

import (
	"fmt"
	"github.com/piojablonski/learn-go-with-tests/poker/store"
	"github.com/piojablonski/learn-go-with-tests/poker/store/filesystem"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("poker command line utility")
	f, err := os.OpenFile("db", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("cannot open file %q", err)
	}
	store, err := filesystem.NewStore(f)
	if err != nil {
		log.Fatalf("cannot create store %v", err)
	}
	cli := CLI{store, os.Stdin}
	for {
		cli.PlayPoker()
	}
}

func (cli *CLI) PlayPoker() {
	var name string
	fmt.Fscanln(cli.in, &name)
	cli.store.RecordWin(name)
}

func NewCLI(store store.PlayerStore, in io.Reader) *CLI {
	return &CLI{store, in}
}

type CLI struct {
	store store.PlayerStore
	in    io.Reader
}
