package utils

import (
	"errors"
)

type Ring struct {
	value       []interface{}
	capacity    int
	front, rear int
}

func NewRing(lens int) *Ring {
	if lens <= 0 {
		return &Ring{value: make([]interface{}, 10), capacity: 10, front: 0, rear: 0}
	}
	return &Ring{value: make([]interface{}, lens), capacity: lens, front: 0, rear: 0}
}

func (self *Ring) EnRing(elem interface{}) bool {
	if self.IsFull() {
		return false
	}
	self.value[self.rear] = elem
	self.rear = (self.rear + 1) % self.capacity
	return true
}

func (self *Ring) DeRing() error {
	if self.IsEmpty() {
		return errors.New("Ring is empty.")
	}
	self.front = (self.front + 1) % self.capacity
	return nil
}

func (self *Ring) Front() interface{} {
	if self.IsEmpty() {
		return nil
	}
	return self.value[self.front]
}
func (self *Ring) Back() interface{} {
	if self.IsEmpty() {
		return nil
	}
	return self.value[self.rear-1]
}

func (self *Ring) Clear() {
	self.front = self.rear
}

func (self *Ring) IsEmpty() bool {
	return self.front == self.rear
}

func (self *Ring) IsFull() bool {
	return (self.rear+1)%self.capacity == self.front
}

func (self *Ring) Length() int {
	if self.IsFull() {
		return self.capacity
	}
	return (self.rear - self.front + self.capacity) % self.capacity
}
