package algods

import (
	"fmt"
	"reflect"
	"testing"
)

func TestList_Length_Default_0(t *testing.T) {
	list := new(List[int])
	if list.Length != 0 {
		t.Errorf("Length was %v but expected %v", list.Length, 0)
	}
}

func TestList_AddHead_Increments_Length(t *testing.T) {
	list := new(List[int])
	list.AddHead(1)
	list.AddHead(2)
	list.AddHead(3)

	if list.Length != 3 {
		t.Errorf("Length was %v but expected %v", list.Length, 3)
	}
}

func TestList_AddHead_Adds(t *testing.T) {
	list := new(List[int])
	list.AddHead(1)
	list.AddHead(2)
	list.AddHead(3)

	if !listContains(list, 3, 2, 1) {
		t.Error("List did not contain the expected values")
	}
}

func TestList_AddTail_Adds_To_Tail(t *testing.T) {
	list := new(List[int])
	list.AddTail(1)
	list.AddTail(2)
	list.AddTail(3)

	if !listContains(list, 1, 2, 3) {
		t.Error("List did not contain the expected values")
	}
}

func TestList_Add_Adds_As_Indicated(t *testing.T) {
	list := new(List[int])
	list.AddTail(2)
	list.AddHead(1)
	list.AddTail(3)
	list.AddHead(0)
	list.AddTail(4)

	if !listContains(list, 0, 1, 2, 3, 4) {
		t.Error("List did not contain the expected values")
	}
}

func TestList_Remove_Actually_Removes(t *testing.T) {
	list := new(List[int])
	list.AddHead(1)
	list.AddHead(2)
	list.AddHead(3)

	list.Remove(2)
	if !listContains(list, 3, 1) {
		t.Error("List did not contain the expected values")
	}

	list.Remove(3)
	if !listContains(list, 1) {
		t.Error("List did not contain the expected values")
	}

	list.Remove(1)
	if !listContains(list) {
		t.Error("List did not contain the expected values")
	}
}

func TestList_Remove_Decrements_Length(t *testing.T) {
	list := new(List[int])
	list.AddHead(1)
	list.AddHead(2)
	list.AddHead(3)

	list.Remove(2)
	if list.Length != 2 {
		t.Errorf("Length was %v but expected %v", list.Length, 2)
	}

	list.Remove(3)
	if list.Length != 1 {
		t.Errorf("Length was %v but expected %v", list.Length, 1)
	}

	list.Remove(1)
	if list.Length != 0 {
		t.Errorf("Length was %v but expected %v", list.Length, 0)
	}
}

func TestList_Remove_Does_Not_Decrement_Length_When_Nothing_Removed(t *testing.T) {
	list := new(List[int])

	list.Remove(10)
	if list.Length != 0 {
		t.Errorf("Length was %v but expected %v", list.Length, 0)
	}

	list.AddHead(1)
	list.Remove(2)
	if list.Length != 1 {
		t.Errorf("Length was %v but expected %v", list.Length, 1)
	}
}

func TestList_Remove_Returns_False_When_Empty(t *testing.T) {
	list := new(List[int])
	if list.Remove(3) {
		t.Error("List should not return true when nothing was removed")
	}
}

func TestList_Remove_Returns_False_When_Nothing_Remove(t *testing.T) {
	list := new(List[int])
	list.AddHead(1)
	list.AddHead(2)
	list.AddHead(3)
	if list.Remove(4) {
		t.Error("List should not return true when nothing was removed")
	}
}

func TestList_Remove_Returns_True_When_Removing(t *testing.T) {
	list := new(List[int])
	list.AddHead(1)
	if !list.Remove(1) {
		t.Error("List should return true when something was removed")
	}
}

func TestList_Remove_Removes_Correct_Duplicate(t *testing.T) {
	list := new(List[int])
	list.AddHead(1)
	list.AddHead(2)
	list.AddHead(1)
	list.AddHead(2)
	list.AddHead(1) // 1 2 1 2 1

	if !listContains(list, 1, 2, 1, 2, 1) {
		t.Error("List did not contain the expected values")
	}

	list.Remove(2)
	if !listContains(list, 1, 1, 2, 1) {
		t.Error("List did not contain the expected values")
	}

	list.Remove(2)
	if !listContains(list, 1, 1, 1) {
		t.Error("List did not contain the expected values")
	}

	list.Remove(1)
	if !listContains(list, 1, 1) {
		t.Error("List did not contain the expected values")
	}

	list.Remove(1)
	if !listContains(list, 1) {
		t.Error("List did not contain the expected values")
	}
}

func listToArray[T comparable](list *List[T]) []T {
	output := make([]T, 0)
	list.ForEach(func(value T) {
		output = append(output, value)
	})

	return output
}

func listContains[T comparable](list *List[T], values ...T) bool {
	if list.Length == 0 && len(values) == 0 {
		return true
	}

	data := listToArray(list)

	if reflect.DeepEqual(data, values) {
		return true
	}

	fmt.Println("Actual:   ", data)
	fmt.Println("Expected: ", values)

	return false
}
