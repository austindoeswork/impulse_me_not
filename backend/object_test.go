package main

import (
	"testing"
)

func makeObjArray(size int) []Object {
	ret := make([]Object, size)
	for i := 0; i < size; i++ {
		ret[i] = Object{UUID: i}
	}
	return ret
}

func TestObjectByUser(t *testing.T) {

}
