package algods

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math/rand"
	"reflect"
	"testing"
)

func TestBinaryTree_Add_Updates_Count(t *testing.T) {
	tree := new(BinaryTree[int])

	for i := 1; i < 25; i++ {
		tree.Add(i)
		if tree.Count != i {
			t.Error("Adding one item should have updated Count")
		}
	}
}

func TestBinaryTree_Add_Builds_Correct_Tree_PostOrder(t *testing.T) {
	tree := new(BinaryTree[int])
	tree.Add(5)
	tree.Add(3)
	tree.Add(4)
	tree.Add(2)
	tree.Add(7)
	tree.Add(6)
	tree.Add(8)

	if !orderMatches(tree.PostOrder, []int{2, 4, 3, 6, 8, 7, 5}) {
		t.Error("Tree does not contain the expected content (post-order)")
	}
}

func TestBinaryTree_Add_Builds_Correct_Tree_InOrder(t *testing.T) {
	tree := new(BinaryTree[int])
	tree.Add(5)
	tree.Add(3)
	tree.Add(4)
	tree.Add(2)
	tree.Add(7)
	tree.Add(6)
	tree.Add(8)

	if !orderMatches(tree.InOrder, []int{2, 3, 4, 5, 6, 7, 8}) {
		t.Error("Tree does not contain the expected content (in-order)")
	}
}

func TestBinaryTree_Add_Builds_Correct_Tree_PreOrder(t *testing.T) {
	tree := new(BinaryTree[int])
	tree.Add(5)
	tree.Add(3)
	tree.Add(4)
	tree.Add(2)
	tree.Add(7)
	tree.Add(6)
	tree.Add(8)

	if !orderMatches(tree.PreOrder, []int{5, 3, 2, 4, 7, 6, 8}) {
		t.Error("Tree does not contain the expected content (pre-order)")
	}
}

func TestBinaryTree_Remove_Empty(t *testing.T) {
	tree := new(BinaryTree[int])

	if tree.Remove(10) {
		t.Error("Removing a value from an empty tree should return false")
	}

	if tree.Count != 0 {
		t.Error("After empty remove, count was not 0: ", tree.Count)
	}
}

func TestBinaryTree_Remove_Root_Sole(t *testing.T) {
	tree := new(BinaryTree[int])
	tree.Add(10)

	if !tree.Remove(10) {
		t.Error("Removing a value from an empty tree should return true")
	}

	if tree.Count != 0 {
		t.Error("After remove, count was not 0: ", tree.Count)
	}
}

func TestBinaryTree_Remove_Root_LeftChild(t *testing.T) {
	tree := new(BinaryTree[int])
	tree.Add(5)
	tree.Add(4)

	if !tree.Remove(5) {
		t.Error("Removing a value from an empty tree should return true")
	}

	if tree.Count != 1 {
		t.Error("After remove, count was not 0: ", tree.Count)
	}

	if !orderMatches(tree.InOrder, []int{4}) {
		t.Error("Tree does not contain the expected content (in-order)")
	}
}

func TestBinaryTree_Remove_Root_RightChild(t *testing.T) {
	tree := new(BinaryTree[int])
	tree.Add(5)
	tree.Add(6)

	if !tree.Remove(5) {
		t.Error("Removing a value from an empty tree should return true")
	}

	if tree.Count != 1 {
		t.Error("After remove, count was incorrect: ", tree.Count)
	}

	if !orderMatches(tree.InOrder, []int{6}) {
		t.Error("Tree does not contain the expected content (in-order)")
	}
}

func TestBinaryTree_Remove_Root_TwoChildren(t *testing.T) {
	tree := new(BinaryTree[int])
	tree.Add(5)
	tree.Add(4)
	tree.Add(6)

	if !tree.Remove(5) {
		t.Error("Removing a value from an empty tree should return true")
	}

	if tree.Count != 2 {
		t.Error("After remove, count was incorrect: ", tree.Count)
	}

	if !orderMatches(tree.PreOrder, []int{6, 4}) {
		t.Error("Tree does not contain the expected content (pre-order)")
	}
}

func TestBinaryTree_Remove_Leaf_Left(t *testing.T) {
	tree := new(BinaryTree[int])
	tree.Add(5)
	tree.Add(4)
	tree.Add(6)

	if !tree.Remove(4) {
		t.Error("Removing a value from an empty tree should return true")
	}

	if tree.Count != 2 {
		t.Error("After remove, count was incorrect: ", tree.Count)
	}

	if !orderMatches(tree.PreOrder, []int{5, 6}) {
		t.Error("Tree does not contain the expected content (pre-order)")
	}
}

func TestBinaryTree_Remove_Leaf_Right(t *testing.T) {
	tree := new(BinaryTree[int])
	tree.Add(5)
	tree.Add(4)
	tree.Add(6)

	if !tree.Remove(6) {
		t.Error("Removing a value from an empty tree should return true")
	}

	if tree.Count != 2 {
		t.Error("After remove, count was incorrect: ", tree.Count)
	}

	if !orderMatches(tree.PreOrder, []int{5, 4}) {
		t.Error("Tree does not contain the expected content (pre-order)")
	}
}

func TestBinaryTree_Remove_Inner_Left_Only(t *testing.T) {
	tree := new(BinaryTree[int])
	tree.Add(5)
	tree.Add(3)
	tree.Add(2)

	if !tree.Remove(3) {
		t.Error("Removing a value from an empty tree should return true")
	}

	if tree.Count != 2 {
		t.Error("After remove, count was incorrect: ", tree.Count)
	}

	if !orderMatches(tree.InOrder, []int{2, 5}) {
		t.Error("Tree does not contain the expected content (pre-order)")
	}
}

func TestBinaryTree_Remove_Inner_Right_Only(t *testing.T) {
	tree := new(BinaryTree[int])
	tree.Add(5)
	tree.Add(3)
	tree.Add(4)

	if !tree.Remove(3) {
		t.Error("Removing a value from an empty tree should return true")
	}

	if tree.Count != 2 {
		t.Error("After remove, count was incorrect: ", tree.Count)
	}

	if !orderMatches(tree.InOrder, []int{4, 5}) {
		t.Error("Tree does not contain the expected content (pre-order)")
	}
}

func TestBinaryTree_Remove_Inner_Both(t *testing.T) {
	tree := new(BinaryTree[int])
	tree.Add(5)
	tree.Add(3)
	tree.Add(4)
	tree.Add(2)

	if !tree.Remove(3) {
		t.Error("Removing a value from an empty tree should return true")
	}

	if tree.Count != 3 {
		t.Error("After remove, count was incorrect: ", tree.Count)
	}

	if !orderMatches(tree.InOrder, []int{2, 4, 5}) {
		t.Error("Tree does not contain the expected content (pre-order)")
	}
}

func TestBinaryTree_Remove_Inner_Deep(t *testing.T) {
	tree := new(BinaryTree[int])
	start := []int{5, 20, 10, 30, 7, 15, 25, 35, 6, 9, 23, 27, 28}
	expected := []int{5, 6, 7, 9, 10, 15, 20, 23, 25, 27, 28, 35}

	for _, value := range start {
		tree.Add(value)
	}

	if !tree.Remove(30) {
		t.Error("Removing a value from an empty tree should return true")
	}

	if tree.Count != len(start)-1 {
		t.Error("After remove, count was incorrect: ", tree.Count)
	}

	if !orderMatches(tree.InOrder, expected) {
		t.Error("Tree does not contain the expected content (in-order)")
	}
}

func TestBinaryTree_Add_From_Empty_Adds_Root(t *testing.T) {
	tree := new(BinaryTree[int])
	tree.Add(5)

	if !orderMatches(tree.InOrder, []int{5}) {
		t.Error("Tree does not contain the expected content (in-order)")
	}
}

func TestBinaryTree_Random_Big(t *testing.T) {
	tree := new(BinaryTree[int])
	data := make([]int, 0)

	for i := 0; i < 25000; i++ {
		value := rand.Int()
		data = append(data, value)
		tree.Add(value)
	}

	if tree.Count != len(data) {
		t.Error("Count should be equal to data size: ", len(data))
	}

	for _, v := range data {
		if !tree.Contains(v) {
			t.Error("Tree should contain ", v)
		}
	}

	for _, v := range data {
		if !tree.Remove(v) {
			t.Error("Should have been able to remove from tree: ", v)
		}
	}

	if tree.Count != 0 {
		t.Error("Tree should be empty")
	}
}

func TestBinaryTree_Contains_Empty_Tree(t *testing.T) {
	tree := new(BinaryTree[int])
	if tree.Contains(0) {
		t.Error("The tree should not contain any values")
	}
}

func TestBinaryTree_Contains_Only_Root(t *testing.T) {
	tree := new(BinaryTree[int])
	tree.Add(5)

	if !tree.Contains(5) {
		t.Error("The tree should contain 5")
	}

	if tree.Contains(0) {
		t.Error("The tree should not contain 0")
	}
}

func TestBinaryTree_Contains_Many(t *testing.T) {
	tree := new(BinaryTree[int])
	start := []int{5, 20, 10, 30, 7, 15, 25, 35, 6, 9, 23, 27, 28}
	missing := []int{1, 3, 12, 22, 29, 36}

	for _, value := range start {
		tree.Add(value)
	}

	for _, value := range start {
		if !tree.Contains(value) {
			t.Error("The tree should contain the value: ", value)
		}
	}

	for _, value := range missing {
		if tree.Contains(value) {
			t.Error("The tree should not contain the value: ", value)
		}
	}

}

func orderMatches[T constraints.Ordered](order func(func(value T)), expected []T) bool {
	actual := make([]T, 0)

	order(func(value T) {
		actual = append(actual, value)
	})

	if !reflect.DeepEqual(expected, actual) {
		fmt.Printf("Expected: %v\n", expected)
		fmt.Printf("Actual:   %v\n", actual)
		return false
	}

	return true
}
