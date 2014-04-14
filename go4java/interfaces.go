package go4java

import (
	"fmt"
	"os"
)

type MyWriter interface {
	Write([]byte) (int, error)
}

func PrintHiTo(w MyWriter) {
	fmt.Fprintf(w, "hi")
}

func ExamplePrintHiTo() {
	PrintHiTo(os.Stdout)
}
