package log

import "fmt"

var log = newLogRedirector()

// GetRedirector ...
func GetRedirector() *Redirector {
	return log
}

// Redirector ...
type Redirector struct {
	Print   PrintFn
	Printf  PrintfFn
	Println PrintlnFn
}

func newLogRedirector() *Redirector {
	r := &Redirector{}
	r.SetDefaults()
	return r
}

type (
	// PrintFn ...
	PrintFn = func(a ...interface{})
	// PrintfFn ...
	PrintfFn = func(format string, a ...interface{})
	// PrintlnFn ...
	PrintlnFn = func(a ...interface{})
)

// SetDefaults ...
func (r *Redirector) SetDefaults() {
	r.Print = func(a ...interface{}) {
		fmt.Print(a...)
	}

	r.Printf = func(format string, a ...interface{}) {
		fmt.Printf(format, a...)
	}

	r.Println = func(a ...interface{}) {
		fmt.Println(a...)
	}
}

// SetPrint ...
func (r *Redirector) SetPrint(fn PrintFn) {
	r.Print = fn
}

// SetPrintf ...
func (r *Redirector) SetPrintf(fn PrintfFn) {
	r.Printf = fn
}

// SetPrintln ...
func (r *Redirector) SetPrintln(fn PrintlnFn) {
	r.Println = fn
}
