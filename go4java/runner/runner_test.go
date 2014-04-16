package runner_test

import (
	"fmt"

	"github.com/campoy/go4java/go4java/runner"
)

func ExampleRunCounter() {
	r := runner.NewRunCounter("my runner")
	r.RunAll([]runner.Task{
		{"one"},
		{"two"},
		{"three"},
	})
	fmt.Printf("%s ran %d tasks\n", r.Name(), r.Count())
	// Output:
	// running one
	// running two
	// running three
	// my runner ran 3 tasks
}

func ExampleRunCounter2() {
	r := runner.NewRunCounter2("my runner")
	r.RunAll([]runner.Task{
		{"one"},
		{"two"},
		{"three"},
	})
	fmt.Printf("%s ran %d tasks\n", r.Name(), r.Count())
	// Output:
	// running one
	// running two
	// running three
	// my runner ran 3 tasks
}
