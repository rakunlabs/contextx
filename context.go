package contextx

import (
	"context"
)

type CtxKey string

const CtxInternal CtxKey = "internal"

var unUsedValue = newData()

// Value gets value from context's map.
//
// Returns same as map's value and ok.
func Value[T any](ctx context.Context, key any, opts ...Option) (T, bool) { //nolint:ireturn // ignore
	var o *Data
	if ctx == nil {
		o = unUsedValue
	} else {
		var ok bool
		if o, ok = Internal(ctx, opts...); !ok {
			o = unUsedValue
		}
	}

	o.M.RLock()
	defer o.M.RUnlock()

	ret, ok := o.V[key].(T)

	return ret, ok
}

// WithValue sets value to context's map.
//
// If context is nil, it use context.Background().
// If context's map is nil, it will init context's map and add value to ctx.
func WithValue(ctx context.Context, key, value any, opts ...Option) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	o, ok := Internal(ctx, opts...)
	if !ok {
		return WithValue(Init(ctx, opts...), key, value, opts...)
	}

	o.M.Lock()
	defer o.M.Unlock()

	o.V[key] = value

	return ctx
}

// Init context with internal value.
//
// If context is nil, it use context.Background().
func Init(ctx context.Context, opts ...Option) context.Context {
	opt := newOption(opts...)

	if ctx == nil {
		ctx = context.Background()
	}

	return context.WithValue(ctx, opt.Key, newData())
}

// Internal gets internal data value from context.
//
// If context is nil return nil and false.
func Internal(ctx context.Context, opts ...Option) (*Data, bool) {
	opt := newOption(opts...)

	if ctx == nil {
		return nil, false
	}

	v, ok := ctx.Value(opt.Key).(*Data)

	return v, ok
}
