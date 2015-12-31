package utils

import (
	"testing"
)

func Test_List_Add(t *testing.T) {
	list := NewList(1, 2, 3)
	list.Add(4, 5, 6, 8)
	if list.Length() != 7 {
		t.Fatal("Add failed.")
	}
}

func Test_List_Remove(t *testing.T) {
	list := NewList(1, 2, 3, 4, 5, 5)
	list.RemoveIf(func(elem interface{}) bool {
		return elem == 5
	})
	if list.Length() != 4 {
		t.Logf("%v\n", list)
		t.Error("remove failed.")
	}
	/*
		list.Add(2, 2, 2)
		list.RemoveFirst(2, 10)
		if list.Length() != 2 {
			t.Logf("%v\n", list)
			t.Error("remove failed.")
		}
	*/
	list1 := NewList(13, 13, 15)

	list1.Remove(13)
	if list1.Length() != 1 {
		t.Logf("%v\n", list1)
		t.Error("remove failed.")
	}

	list2 := NewList(23, 23, 23, 24, 25, 23)
	list2.RemoveFirst(23, 4)
	if list2.Length() != 2 {
		t.Logf("%v\n", list2)
		t.Error("remove failed.")
	}

	list2.Clear()
	if list2.length != 0 {
		t.Logf("%v\n", list2)
		t.Error("clear failed.")
	}
}

func Test_List_Index(t *testing.T) {
	list := NewList(1, 2, 3, 4, 5, 6, 7, 7, 8)

	lens := list.Length()
	//t.Logf("Length:[%d]...\n", lens)
	f11, _ := list.IndexOf(0)
	//t.Logf("Elem:%d\n", f11)
	f12, _ := list.IndexOf(-lens)
	//t.Logf("Elem:%d\n", f12)
	f21, _ := list.LastIndexOf(-1)
	f22, _ := list.LastIndexOf(lens - 1)

	t11, _ := list.IndexOf(lens - 1)
	t12, _ := list.IndexOf(-1)

	t21, _ := list.LastIndexOf(0)
	t22, _ := list.LastIndexOf(-lens)

	t.Logf("%d %d %d %d\n", f11, f12, f21, f22)
	t.Logf("%d %d %d %d\n", t11, t12, t21, t22)
	if f11 != f21 || t11 != t21 {
		t.Error("index failed.")
	}

	if _, err := list.IndexOf(lens); err == nil {
		t.Error("out of index failed.")
	}
}

func Test_List_Pop(t *testing.T) {
	list := NewList(1, 2, 3, 4, 5, 6, 7, 7, 8)
	emplist := NewList()
	e := list.PopBack()
	f := list.PopFront()
	n1 := emplist.PopFront()
	n2 := emplist.PopBack()
	if e != 8 || f != 1 || n1 != nil || n2 != n1 {
		t.Error("Pop failed.")
	}
}
