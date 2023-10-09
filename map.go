package contextx

import "sync"

type InternalValue struct {
	V map[any]interface{}

	M sync.RWMutex
}

func newInternalValue() *InternalValue {
	return &InternalValue{
		V: make(map[any]interface{}),
	}
}
