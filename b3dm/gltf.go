package b3dm

import (
	"io"
	"errors"

	"github.com/qmuntal/gltf"
)

func loadGltfFromByte(reader io.Reader) (*gltf.Document, error) {
	dec := gltf.NewDecoder(reader)
	doc := new(gltf.Document)
	if err := dec.Decode(doc); err != nil {
		return nil, errors.New("failed to decode the glTF doc")
	}
	return doc, nil
}
