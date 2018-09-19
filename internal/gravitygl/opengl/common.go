package opengl

import (
	"github.com/go-gl/gl/v4.3-core/gl"
)

// GetString ...
func GetString(name uint32) string {
	return gl.GoStr(
		gl.GetString(name),
	)
}
