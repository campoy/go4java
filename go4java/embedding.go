package main

import "fmt"

type Task struct{ msg string }

func (t Task) Run() { fmt.Println("running", t.msg) }

type Runner struct{ name string }

func NewRunner(name string) *Runner { return &Runner{name} }
func (r *Runner) Name() string      { return r.name }
func (r *Runner) Run(t Task)        { t.Run() }
func (r *Runner) RunAll(ts []Task) {
	for _, t := range ts {
		r.Run(t)
	}
}

// COUNTING OMIT
type CountingRunner struct {
	*Runner // HL
	count   int
}

func NewCountingRunner(name string) *CountingRunner {
	return &CountingRunner{NewRunner(name), 0}
}

func (r *CountingRunner) Run(t Task) {
	r.count++
	r.Runner.Run(t)
}

func (r *CountingRunner) RunAll(ts []Task) {
	r.count += len(ts)
	r.Runner.RunAll(ts)
}

func (r *CountingRunner) Count() int { return r.count }

// MAIN OMIT
func main() {
	r := NewCountingRunner("my runner")
	tasks := []Task{{"one"}, {"two"}, {"three"}}
	r.RunAll(tasks)
	fmt.Printf("%s ran %d tasks\n", r.Name(), r.Count())
}
