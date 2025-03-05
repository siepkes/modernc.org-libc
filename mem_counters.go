//go:build memory.counters

package libc

import "expvar"

func init() {
	expvar.Publish("memory.allocator", expvar.Func(func() interface{} {
		return MemStat()
	}))
}
