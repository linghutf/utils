package utils

import (
	"testing"
)

func Test_LinkQueue_Handler(t *testing.T) {
	lq := NewLinkQueue()
	if !lq.IsEmpty() {
		t.Error("IsEmpty failed.")
	}

	lq.EnQueue(12)
	if lq.Front() != lq.Back() {
		t.Logf("front:%d.back:%d\n", lq.Front(), lq.Back())
		t.Error("Front/Back failed.")
	}

	lq.DeQueue()
	if !lq.IsEmpty() {
		t.Error("IsEmpty failed.")
	}

	lq.EnQueue(13)
	lq.EnQueue(14)
	if lq.Length() != 2 {
		t.Error("Length failed.")
	}

	lq.DeQueue()
	lq.Clear()
	if lq.Front() != nil {
		t.Logf("front:%d\n", lq.Front())
		t.Error("Clear failed.")
	}
	if lq.Back() != nil {
		t.Logf("back:%d\n", lq.Back())
		t.Error("Clear failed.")
	}
}
