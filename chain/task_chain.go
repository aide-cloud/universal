package chain

type (
	Interface func() error

	Chain struct {
		tasks []Interface
	}

	Option func(*Chain)
)

// WithTask replaces the tasks with the given chain.
func WithTask(chain ...Interface) Option {
	return func(c *Chain) {
		c.tasks = append(c.tasks, chain...)
	}
}

// NewChain creates a new chain.
func NewChain(opts ...Option) *Chain {
	c := &Chain{}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Do execute the chain.
func (c *Chain) Do() error {
	for _, task := range c.tasks {
		if err := task(); err != nil {
			return err
		}
	}
	return nil
}

// Append appends a task to the chain.
func (c *Chain) Append(chain ...Interface) *Chain {
	c.tasks = append(c.tasks, chain...)
	return c
}
