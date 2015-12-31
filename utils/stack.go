package utils

import (
	"errors"
)

type Stack struct {
	Value    []interface{}
	cur      int
	capacity int
}

func NewStack(lens int) *Stack {
	if lens <= 0 {
		return &Stack{Value: make([]interface{}, 10, 20), cur: -1, capacity: 20}
	}
	return &Stack{Value: make([]interface{}, lens, lens<<1), cur: -1, capacity: lens << 1}
}

func (self *Stack) increase() {
	newValue := make([]interface{}, self.capacity<<1, self.capacity<<2)
	copy(newValue, self.Value)
	self.Value = newValue
	self.capacity <<= 2
}

func (self *Stack) Push(elem interface{}) {
	if self.isFull() {
		self.increase()
	}
	self.cur++
	self.Value[self.cur] = elem
}

func (self *Stack) Pop() error {
	if self.IsEmpty() {
		return errors.New("stack is empty.")
	}
	self.cur--
	return nil
}

func (self *Stack) Top() interface{} {
	if self.IsEmpty() {
		return nil
	}
	return self.Value[self.cur]
}

func (self *Stack) IsEmpty() bool {
	return self.cur == -1
}

func (self *Stack) isFull() bool {
	return self.cur == (self.capacity - 1)
}

func (self *Stack) Length() int {
	return self.cur + 1
}

//call Func from buttom to top if Func(op) is true
func (self *Stack) Traverse(Func func(elem interface{}) bool) {
	if self.IsEmpty() {
		return
	}
	for i := 0; i <= self.cur; i++ {
		if !Func(self.Value[self.cur]) {
			return
		}
	}
}
