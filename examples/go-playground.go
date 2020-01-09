package examples

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

type count32 int32

func (c *count32) increment() int32 {
	var next int32
	for {
		next = int32(*c) + 1
		if atomic.CompareAndSwapInt32((*int32)(c), int32(*c), next) {
			return next
		}
	}
}

func (c *count32) get() int32 {
	return atomic.LoadInt32((*int32)(c))
}

func main() {
	fmt.Println("GOMAXPROCS set from", runtime.GOMAXPROCS(runtime.NumCPU()), "to", runtime.GOMAXPROCS(0))
	test(10)
}

func test(n int) {
	var c count32
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			fmt.Println(c.increment(), c.get(), "values may differ") // increment and get may return different values
			wg.Done()
		}()
	}
	wg.Wait()
}
