package gravity

import (
	"bytes"
	"fmt"
	"text/tabwriter"

	"github.com/go-gl/mathgl/mgl32"
)

type MapFunc func(r, g, b, a *float32)

func MapFloat4Array(target *[]float32, fn MapFunc) {
	for i := 0; i < len(*target); i += 4 {
		fn(&(*target)[i+0], &(*target)[i+1], &(*target)[i+2], &(*target)[i+3])
	}
}

func PrintVec3(v mgl32.Vec3) string {
	buf := new(bytes.Buffer)
	w := tabwriter.NewWriter(buf, 8, 1, 0, ' ', tabwriter.AlignRight)
	for _, i := range v {
		fmt.Fprintf(w, "%0.3f\t", i)
	}
	fmt.Fprint(w, "\t")
	w.Flush()
	return buf.String()
}
