package utils

import (
	"errors"
)

type Queue struct {
	data       []interface{}
	head, tail int
	capacity   int
}

func NewQueue(lens int) *Queue {
	if lens <= 0 {
		return &Queue{data: make([]interface{}, 10, 20), capacity: 20, head: 0, tail: 0}
	}
	return &Queue{data: make([]interface{}, lens, lens<<1), capacity: lens << 1, head: 0, tail: 0}
}

func (self *Queue) isFull() bool {
	return (self.tail + 1) == self.capacity
}

func (self *Queue) IsEmpty() bool {
	return self.tail == self.head
}

func (self *Queue) Length() int {
	return self.tail - self.head
}

func (self *Queue) increase() {
	if self.Length() < self.capacity {
		copy(self.data, self.data[self.head:self.tail])

		self.tail -= self.head
		self.head = 0
	} else {
		newdata := make([]interface{}, self.capacity<<1, self.capacity<<2)
		copy(newdata, self.data)
		self.data = newdata
		self.capacity <<= 2
	}
}

func (self *Queue) EnQueue(elem interface{}) {
	if self.isFull() {
		self.increase()
	}
	self.data[self.tail] = elem
	self.tail++ //tail point to the tail elem's next
}

func (self *Queue) DeQueue() error {
	if self.IsEmpty() {
		return errors.New("queue is empty.")
	}
	self.head++
	return nil
}

func (self *Queue) Front() interface{} {
	if self.IsEmpty() {
		return nil
	}
	return self.data[self.head]
}

func (self *Queue) Back() interface{} {
	if self.IsEmpty() {
		return nil
	}

	return self.data[self.tail-1]
}
