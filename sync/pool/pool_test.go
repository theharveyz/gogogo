package pool

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			return 0
		},
	}

	a := pool.Get().(int)
	fmt.Println(a)
	pool.Put(1111)
	pool.Put(2)
	// pool.Put(&struct{ A int }{
	// 	A: 10,
	// })
	a = pool.Get().(int)
	fmt.Println(a)
	runtime.GC()
	a = pool.Get().(int)
	fmt.Println(a)
	// 0
	// 1111
	// 0
}
