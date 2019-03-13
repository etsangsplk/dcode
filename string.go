package dcode

import (
	tree "github.com/bmatsuo/go-jsontree"
)

// String decodes any JSON field into a string
func String() Decoder {
	return newDecoder(func(t *tree.JsonTree) (interface{}, error) {
		ret, err := t.String()
		if err != nil {
			return nil, err
		}
		return ret, nil
	})
}
