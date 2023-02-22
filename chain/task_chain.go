package chain

import "sync"

type (
	Interface func() error

	Chain struct {
		locker sync.Mutex
		tasks  []Interface
	}

	Option func(*Chain)
)

// NewChain creates a new chain.
func NewChain(opts ...Option) *Chain {
	c := &Chain{}
	c.locker.Lock()
	defer c.locker.Unlock()
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// AddTask adds a task to the chain.
func AddTask(chain ...Interface) Option {
	return func(c *Chain) {
		c.locker.Lock()
		defer c.locker.Unlock()
		c.tasks = append(c.tasks, chain...)
	}
}

// AppendTask appends a task to the chain.
func AppendTask(chain ...Interface) Option {
	return func(c *Chain) {
		c.locker.Lock()
		defer c.locker.Unlock()
		c.tasks = append(c.tasks, chain...)
	}
}

// Do execute the chain.
func (c *Chain) Do() error {
	c.locker.Lock()
	defer c.locker.Unlock()
	for _, task := range c.tasks {
		if err := task(); err != nil {
			return err
		}
	}
	return nil
}

// Append appends a task to the chain.
func (c *Chain) Append(chain ...Interface) *Chain {
	c.locker.Lock()
	defer c.locker.Unlock()
	c.tasks = append(c.tasks, chain...)
	return c
}
