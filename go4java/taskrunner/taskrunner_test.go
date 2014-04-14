package taskrunner_test

import (
	"fmt"

	"github.com/campoy/go4java/go4java/taskrunner"
)

func ExampleRunner() {
	r := taskrunner.NewRunner("my runner")
	r.Run(func() { fmt.Println("running a literal function") })

	t := taskrunner.Task{"a method value"}
	r.Run(t.Run)
}
