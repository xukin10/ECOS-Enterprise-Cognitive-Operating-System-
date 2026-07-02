package runtime

// Bootstrap is a convenience builder for creating a pre-configured Runtime.
type Bootstrap struct {
	engines []*Engine
}

// NewBootstrap creates an empty Bootstrap builder.
func NewBootstrap() *Bootstrap {
	return &Bootstrap{engines: make([]*Engine, 0)}
}

// Add appends an engine to the bootstrap configuration.
func (b *Bootstrap) Add(engine *Engine) *Bootstrap {
	b.engines = append(b.engines, engine)
	return b
}

// Build creates a Runtime with all added engines registered.
func (b *Bootstrap) Build() (*Runtime, error) {
	rt := New()
	for _, e := range b.engines {
		if err := rt.Register(e); err != nil {
			return nil, err
		}
	}
	return rt, nil
}

// Quick is a one-shot function to create a Runtime from engines, equivalent to
// NewBootstrap().Add(engines...).Build().
func Quick(engines ...*Engine) (*Runtime, error) {
	return NewBootstrap().Add(engines...).Build()
}
