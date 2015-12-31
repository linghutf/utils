package utils

import (
	"fmt"
	"testing"
)

func TestNewSet(t *testing.T) {
	set1 := NewSet(1, 2, 3, 2, 4)
	fmt.Println(set1)

	set2 := NewSet()
	fmt.Println(set2)
}

func Test_Set_Insert(t *testing.T) {
	set := NewSet(1, 2, 3, 4)
	set.Insert(5)
	if !set.IsInSet(5) {
		t.Error("insert failed!")
	}
}
func TestIsInSet(t *testing.T) {
	set := NewSet(1, 2, 3, 4)
	if !set.IsInSet(1) {
		t.Error("isIn!")
	}
	if set.IsInSet(5) {
		t.Error("not In!")
	}
}

func Test_Set_ToArray(t *testing.T) {
	set := NewSet(1, 2, 3, 4)
	a := set.ToArray()
	for _, n := range a {
		if !set.IsInSet(n) {
			t.Error("to array failed.")
		}
	}
}

func Test_Set_Remove(t *testing.T) {
	set := NewSet(1, 2, 3, 4)

	if set.Remove(4) != nil {
		t.Error("remove failed.")
	}
	if set.Remove(5) == nil {
		t.Error("remove failed.")
	}
}
