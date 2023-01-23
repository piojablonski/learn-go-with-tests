package countdown

import (
	"fmt"
	"io"
)

type Sleeper interface {
	Sleep()
}

func Countdown(buf io.Writer, sleeper Sleeper) {
	for i := 0; i < 3; i++ {

	}
	for i := 3; i > 0; i-- {
		// v := []byte(string(i) + "\n")
		// buf.Write(v)
		fmt.Fprintln(buf, i)
		sleeper.Sleep()
	}
	fmt.Fprint(buf, "go!")
}
