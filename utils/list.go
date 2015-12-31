package utils

import (
	"errors"
	//"log"
)

type GenericList interface {
	Add(elem interface{})
	Remove(key interface{}) error
	Clear() error
	Insert(index int, elem interface{}) error
	IsEmpty() bool
	Length() int
	//can be negetive
	IndexOf(index int) (interface{}, error)
	//LastIndexOf(index int) (interface{}, error)
}

type List struct {
	data             []interface{}
	length, capacity int
}

func NewList(elems ...interface{}) *List {
	lens := len(elems)

	if lens == 0 {
		return &List{data: make([]interface{}, 10, 20), length: 0, capacity: 20}
	}
	list := &List{length: lens, capacity: lens * 2, data: make([]interface{}, lens*2)}
	for i, n := range elems {
		list.data[i] = n
	}
	return list
}

func (self *List) Add(elems ...interface{}) {
	lens := len(elems)
	if lens <= 0 {
		return
	} else if (lens + self.length) > self.capacity {
		tmp := make([]interface{}, (self.length+lens)*2)
		copy(tmp, self.data)
		self.data = tmp
	}

	for _, elem := range elems {
		self.data[self.length] = elem
		self.length++
	}

}

func (self *List) RemoveIf(Func func(elem interface{}) bool) error {
	if self.IsEmpty() {
		return errors.New("List is Empty.")
	}
	newdata := make([]interface{}, self.length, self.capacity)
	index := 0
	//range error
	for i := 0; i < self.length; i++ {
		if Func(self.data[i]) {
			continue
		}
		newdata[index] = self.data[i]
		index++
	}
	if index == self.length {
		return errors.New("element not found.")
	}

	self.data = newdata[:index]
	self.length = index
	return nil
}

func (self *List) Remove(key interface{}) error {
	return self.RemoveIf(func(elem interface{}) bool {
		return elem == key
	})
}

func (self *List) RemoveFirst(key interface{}, times int) error {
	if self.IsEmpty() {
		return errors.New("List is Empty.")
	}
	if times <= 0 {
		return errors.New("remove time must be non-negative.")
	}
	count := 0
	newdata := make([]interface{}, self.length, self.capacity)
	index := 0
	//range error
	for i := 0; i < self.length; i++ {
		if key == self.data[i] {
			count++
			//have done remove
			if count == times {
				//copy tail of origin
				tail := self.length - 1
				for i < tail {
					newdata[index] = self.data[i]
					i++
					index++
				}
				break //out of for-range
			}
			continue //miss the equal number
		}
		//copy not euqal number
		newdata[index] = self.data[i]
		index++
	}
	if index == self.length {
		return errors.New("element not found.")
	}

	self.data = newdata[:index]
	self.length = index
	return nil
}

func (self *List) Clear() {
	//self = NewList() //not effective
	for i := 0; i < self.length; i++ {
		self.data[i] = nil
	}
	self.length = 0
}

func (self *List) IndexOf(index int) (interface{}, error) {
	if index >= self.length || index < (-self.length) {
		return nil, errors.New("index out of range.")
	}
	if index < 0 {
		return self.data[self.length+index], nil
	} else {
		return self.data[index], nil
	}
}

func (self *List) LastIndexOf(index int) (interface{}, error) {
	if index >= self.length || index < (-self.length) {
		return nil, errors.New("index out of range.")
	}
	if index >= 0 {
		return self.data[self.length-1-index], nil
	} else {
		return self.data[-index-1], nil
	}
}

func (self *List) IsEmpty() bool {
	return self.length == 0
}

func (self *List) Length() int {
	return self.length
}

func (self *List) PopFront() interface{} {
	if self.length == 0 {
		return nil
	}
	elem := self.data[0]
	self.data = self.data[1:]
	self.length--
	return elem
}

//type test
/*
func IsSameType(a *List,elem interface{})bool{
	return reflect,Type
}
*/
func (self *List) PopBack() interface{} {
	if self.length == 0 {
		return nil
	}
	elem := self.data[self.length-1]
	self.data = self.data[:self.length-2]
	self.length--
	return elem
}
