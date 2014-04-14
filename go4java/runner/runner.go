// Package runner provides a Runner type that is used to define both CountingRunner
// and EmbeddedCountingRunner to show examples of how to use composition in Go.
package runner

import "fmt"

// A Task is a simple task that prints a message when run.
type Task struct{ Msg string }

func (t Task) Run() {
	fmt.Println("running", t.Msg)
}

// END_TASK OMIT

// A Runner provides a way of running tasks.
type Runner struct{ name string }

func NewRunner(name string) *Runner { return &Runner{name} }
func (r *Runner) Name() string      { return r.name }
func (r *Runner) Run(t Task)        { t.Run() }

func (r *Runner) RunAll(ts []Task) {
	for _, t := range ts {
		r.Run(t)
	}
}

// END_RUNNER OMIT

// A CountingRunner is a Runner that keeps a counter of the run tasks.
type CountingRunner struct {
	runner *Runner
	count  int
}

func NewCountingRunner(name string) *CountingRunner {
	return &CountingRunner{runner: NewRunner(name)}
}

func (r *CountingRunner) Run(t Task) {
	r.count++
	r.runner.Run(t)
}

func (r *CountingRunner) RunAll(ts []Task) {
	r.count += len(ts)
	r.runner.RunAll(ts)
}

func (r *CountingRunner) Count() int { return r.count }

func (r *CountingRunner) Name() string { return r.runner.Name() }
