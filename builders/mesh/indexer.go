package meshbuilder

import (
	"github.com/thegtproject/gravity/geometry"
)

// Indexer ...
type Indexer struct {
	unique []geometry.Vec3
	table  map[geometry.Vec3]int
}

// Add ...
func (idxr *Indexer) Add(vec geometry.Vec3) (index int) {
	if i, ok := idxr.table[vec]; !ok {
		idxr.pushUnique(vec)
		idxr.table[vec] = len(idxr.unique)
	} else {
		return i
	}
	return idxr.table[vec]
}

func (idxr *Indexer) pushUnique(vec geometry.Vec3) {
	idxr.unique = append(idxr.unique, vec)
}
