package contextx

type option struct {
	Key any
}

func newOption(opts ...Option) *option {
	o := &option{
		Key: CtxInternal,
	}
	for _, opt := range opts {
		opt(o)
	}

	return o
}

type Option func(*option)

// WithKey sets the key for the option.
func WithKey(key any) Option {
	return func(o *option) {
		o.Key = key
	}
}
