package greeter

import (
	"fmt"
	"io"
)

func Greet(b io.Writer, name string) {
	fmt.Fprintf(b, "Hello %s", name)
}
