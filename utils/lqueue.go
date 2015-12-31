package utils

import (
	"errors"
)

//linking queue
//head->[empty node]<-tail: empty queue
//head->[empty node]->[x node]<-tail:x elem

type LinkQueue struct {
	head, tail *Node
	length     int
}

func NewLinkQueue() *LinkQueue {
	lq := new(LinkQueue)
	lq.head = NewNode()
	lq.head.next = nil
	lq.tail = lq.head
	return lq
}

func (self *LinkQueue) IsEmpty() bool {
	return self.head == self.tail
}

func (self *LinkQueue) EnQueue(elem interface{}) bool {
	p := NewNode()
	if p == nil {
		return false
	}
	p.Value = elem
	self.tail.next = p
	self.tail = p
	self.length++
	return true
}

func (self *LinkQueue) DeQueue() error {
	if self.head == self.tail {
		return errors.New("queue is empty.")
	}

	p := self.head.next
	self.head.next = p.next
	if self.tail == p {
		self.tail = self.head
	}
	self.length--
	return nil
}

func (self *LinkQueue) Clear() {
	self.tail = self.head
}

func (self *LinkQueue) Length() int {
	return self.length
}

func (self *LinkQueue) Front() interface{} {
	if self.IsEmpty() {
		return nil
	}
	return self.head.next.Value
}

func (self *LinkQueue) Back() interface{} {
	if self.IsEmpty() {
		return nil
	}
	return self.tail.Value
}
