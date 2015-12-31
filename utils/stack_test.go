package utils

import (
	"testing"
)

func Test_Stack_Handle(t *testing.T) {
	s := NewStack(5)
	if !s.IsEmpty() {
		t.Error("empty failed.")
	}
	s.Push(1)
	s.Push(2)
	if s.Top() != 2 {
		t.Error("top failed.")
	}
	s.Pop()
	if s.Length() != 1 {
		t.Error("pop failed.")
	}

}
