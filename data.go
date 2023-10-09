package contextx

import "sync"

type Data struct {
	V map[any]interface{}

	M sync.RWMutex
}

func newData() *Data {
	return &Data{
		V: make(map[any]interface{}),
	}
}
