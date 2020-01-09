package atomics

import "sync/atomic"

type count64 int64

func (c *count64) Increment() {
	for {
		next := int64(*c) + 1
		if atomic.CompareAndSwapInt64((*int64)(c), int64(*c), next) {
			return
		}
	}
}

func (c *count64) IncrementWithReturn() int64 {
	var next int64
	for {
		next = int64(*c) + 1
		if atomic.CompareAndSwapInt64((*int64)(c), int64(*c), next) {
			return next
		}
	}
}

func (c *count64) Get() int64 {
	return atomic.LoadInt64((*int64)(c))
}

func (c *count64) Reset() {
	for {
		zero := int64(0)
		if atomic.CompareAndSwapInt64((*int64)(c), int64(*c), zero) {
			return
		}
	}
}

var TotalCounts count64
