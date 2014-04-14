package taskrunner

import "fmt"

type Runner struct {
	name string
}

func NewRunner(name string) *Runner {
	return &Runner{name}
}

// Run runs the given function.
func (r Runner) Run(f func()) {
	f()
}

// A Task is a simple task that prints a message when run.
type Task struct {
	Msg string
}

func (t Task) Run() {
	fmt.Println("running", t.Msg)
}
