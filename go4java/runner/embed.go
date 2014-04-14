package runner

// EmbeddingCountingRunner is completely equivalent to CountingRunner,
// but uses struct embedding to avoid the boilerplate of redeclaring
// the Name method.
type EmbeddingCountingRunner struct {
	*Runner // HL
	count   int
}

func NewEmbeddingCountingRunner(name string) *EmbeddingCountingRunner {
	return &EmbeddingCountingRunner{NewRunner(name), 0}
}

func (r *EmbeddingCountingRunner) Run(t Task) {
	r.count++
	r.Runner.Run(t) // HL
}

func (r *EmbeddingCountingRunner) RunAll(ts []Task) {
	r.count += len(ts)
	r.Runner.RunAll(ts) // HL
}

func (r *EmbeddingCountingRunner) Count() int { return r.count }
