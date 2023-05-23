package core

import "testing"

func TestSimpleString(t *testing.T) {
	s := NewSDSFromValue([]byte{0, 33, 49})
	s.Print()
}
