package utils

import (
	"testing"
)

func Test_Ring_Handler(t *testing.T) {
	r := NewRing(3)
	r.EnRing(1)
	r.EnRing(1)
	r.EnRing(2)
	if !r.IsFull() {
		t.Error("IsFull failed.")
	}

	r.DeRing()
	r.DeRing()
	if r.Length() != 1 {
		t.Error("Length failed.")
	}
	r.DeRing()
	if !r.IsEmpty() {
		t.Error("IsEmpty failed.")
	}
	r.EnRing(12)
	if r.Front() != 12 || r.Back() != 12 {
		t.Error("Front/Back failed.")
	}
}
