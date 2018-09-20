//+build !debug

package mtx

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

// MTX scheduler is used for managing system calls that are required
// to come from the main os thread of the application.
var (
	queue           CallQueue
	mtxCallQueueCap int
	total, max      uint64
	current         int64
	doneSignal      chan struct{}
	startTime       time.Time
)

// Init ...
func Init(maxQueue int) {
	queue = make(CallQueue, maxQueue)
	mtxCallQueueCap = maxQueue
}

// Run ...
func Run(f func()) {
	startTime = time.Now()
	doneSignal = make(chan struct{}, 1)
	go func() {
		f()
		doneSignal <- struct{}{}
	}()
	handler()
}

func handler() {
	running := true
	for running == true {
		select {
		case mtxcall := <-queue:
			mtxcall()
			currdec()
		case <-doneSignal:
			running = false
		}
	}
}

// CallNoneBlocking (none-blocking) queues the callfunc on MTX and proceeds
func CallNoneBlocking(f CallFunc) {
	queue <- f
	incTotal()
	if cnt := incCurrent(); cnt > 0 {
		if i := uint64(cnt); i > getMax() {
			setMax(i)
		}
	}
}

// Call blocks until it has been processed
func Call(f CallFunc) {
	done := make(chan struct{})
	queue <- func() {
		f()
		done <- struct{}{}
	}
	<-done
}

// StatsCallsPerSec ...
func StatsCallsPerSec() float64 {
	val := float64(atomic.LoadUint64(&total))
	return val / time.Since(startTime).Seconds()
}

func setMax(val uint64) {
	atomic.StoreUint64(&max, val)
}

func getMax() uint64 {
	return atomic.LoadUint64(&max)
}

func incTotal() {
	atomic.AddUint64(&total, 1)
}

func incCurrent() int64 {
	return atomic.AddInt64(&current, 1)
}

func currdec() {
	atomic.AddInt64(&current, -1)
}

// CallQueue ...
type CallQueue = chan func()

// CallFunc ...
type CallFunc = func()

func init() {
	runtime.LockOSThread()
}

// PrintStats ...
func PrintStats() {
	fmt.Printf(statsTemplate, mtxCallQueueCap,
		current, StatsCallsPerSec(),
		total, max-1,
	)
}

const statsTemplate = `
MTX Stats (Cap = %d)
Current: %d      Calls/sec: %f
Total:   %d
Max:     %d
`
