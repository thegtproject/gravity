package gravity

// Callbacks ...
var Callbacks = struct {
	CaptureModeOnChange func(bool)
}{
	CaptureModeOnChange: func(b bool) {},
}
