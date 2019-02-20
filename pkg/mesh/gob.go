package mesh

import (
	"encoding/gob"
	"os"
)

func writeGob(path string, object interface{}) error {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	encoder := gob.NewEncoder(f)
	return encoder.Encode(object)
}

func readGob(path string, object interface{}) error {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	decoder := gob.NewDecoder(f)
	return decoder.Decode(object)
}
