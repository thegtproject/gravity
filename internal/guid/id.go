package id

import (
	"sync/atomic"
)

var counter uint32

// Next ...
func Next() uint32 {
	return atomic.AddUint32(
		&counter, 1,
	)
}
