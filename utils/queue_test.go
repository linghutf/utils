package utils

import (
	"testing"
)

func Test_Queue_Handle(t *testing.T) {
	q := NewQueue(1)
	t.Logf("len:%d,cap:%d", q.Length(), q.capacity)
	q.EnQueue(2)
	if q.Length() != 1 {
		t.Error("enqueue/length failed.")
	}

	q.DeQueue()
	if !q.IsEmpty() {
		t.Error("dequeue/isempty failed.")
	}

	q.EnQueue(12)
	if q.Front() != 12 {
		t.Error("front failed.")
	}
	if q.Back() != 12 {
		t.Error("back failed.")
	}
}
