package id

import (
	"sync/atomic"
)

// GUID ...
type GUID = uint32

type idgenerator struct {
	counter GUID
}

var defaultGenerator = idgenerator{}

// NextID ...
func NextID() GUID {
	return atomic.AddUint32(
		&defaultGenerator.counter, 1,
	)
}
