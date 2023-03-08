package algods

import (
	"fmt"
	"golang.org/x/exp/constraints"
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

func TestBinaryTree_Add_From_Empty_Adds_Root(t *testing.T) {
	tree := new(BinaryTree[int])
	tree.Add(5)

	if !orderMatches(tree.InOrder, []int{5}) {
		t.Error("Tree does not contain the expected content (in-order)")
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
