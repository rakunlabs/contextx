package contextx

import (
	"context"
)

type ctxInternalKey string

const ctxInternal ctxInternalKey = "internal"

var unUsedValue = newData()

// Value gets value from context's map.
//
// Returns same as map's value and ok.
func Value[T any](ctx context.Context, key any) (T, bool) { //nolint:ireturn // ignore
	var o *Data
	if ctx == nil {
		o = unUsedValue
	} else {
		var ok bool
		if o, ok = Internal(ctx); !ok {
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
func WithValue(ctx context.Context, key, value any) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	o, ok := Internal(ctx)
	if !ok {
		return WithValue(Init(ctx), key, value)
	}

	o.M.Lock()
	defer o.M.Unlock()

	o.V[key] = value

	return ctx
}

// Init context with internal value.
//
// If context is nil, it use context.Background().
func Init(ctx context.Context) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}

	return context.WithValue(ctx, ctxInternal, newData())
}

// Internal gets internal data value from context.
//
// If context is nil return nil and false.
func Internal(ctx context.Context) (*Data, bool) {
	if ctx == nil {
		return nil, false
	}

	v, ok := ctx.Value(ctxInternal).(*Data)

	return v, ok
}
