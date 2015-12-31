package utils

import (
	"errors"
)

type GenericSet interface {
	//Get() interface{}
	Remove(key interface{}) error
	Insert(elem interface{})
}

type Set struct {
	Value    map[interface{}]bool
	capacity int
}

func NewSet(elems ...interface{}) *Set {
	lens := len(elems)
	set := Set{Value: make(map[interface{}]bool, lens), capacity: lens}
	for _, elem := range elems {
		set.Value[elem] = true
	}
	return &set
}

func (self *Set) Insert(elem interface{}) {
	self.Value[elem] = true
}

func (self *Set) IsInSet(elem interface{}) bool {
	return self.Value[elem] == true
}

func (self *Set) ToArray() []interface{} {
	arr := make([]interface{}, len(self.Value))
	i := 0
	for k, _ := range self.Value {
		arr[i] = k
		i++
	}
	return arr
}

func (self *Set) Remove(key interface{}) error {
	if self.capacity == 0 {
		return errors.New("set is empty!")
	}
	for k, _ := range self.Value {
		if k == key {
			delete(self.Value, k)
			return nil
		}
	}
	return errors.New("key is not in set!")
}

func (self *Set) Length() int {
	return len(self.Value)
}
