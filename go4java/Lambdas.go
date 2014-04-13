package main

import "fmt"

// START_RUNNER OMIT
type Runner struct{}

func (r Runner) Run(f func()) { f() }

// START_TASK OMIT
type Task struct{ msg string }

func (t Task) Run() { fmt.Println("running", t.msg) }

// END_TASK OMIT

func main() {
	// RUN_LIT OMIT
	r := Runner{}
	r.Run(func() { fmt.Println("running a literal function") })

	// RUN_METHOD OMIT
	t := Task{"an method value"}
	r.Run(t.Run)
	// END_MAIN OMIT
}
