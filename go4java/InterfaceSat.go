package main

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

func main() {
	PrintHiTo(os.Stdout)
}