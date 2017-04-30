package tutorial

import (
	"math/rand"
	"reflect"
	"testing"

	"golang.org/x/tour/tree"
)

func TestBinaryTreeWalk(t *testing.T) {
	t1 := tree.New(1)

	test := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := make([]int, 10)

	ch := make(chan int, 10)
	go Walk(t1, ch)

	for i := range result {
		result[i] = <-ch
	}
	if !reflect.DeepEqual(test, result) {
		t.Error("Tree channel data do not equal test data.")
	}
}

// The following two functions were copied from
// https://github.com/golang/tour/blob/master/tree/tree.go and adapted to
// create a descending tree for testing
func newDescendingTree(k int) *tree.Tree {
	var t *tree.Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, (1-v)*k)
	}
	return t
}

func insert(t *tree.Tree, v int) *tree.Tree {
	if t == nil {
		return &tree.Tree{Left: nil, Value: v, Right: nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}
	return t
}

func TestBinaryTreeSame(t *testing.T) {
	t1 := tree.New(1)
	t2 := tree.New(1)
	t3 := &tree.Tree{
		Left:  new(tree.Tree),
		Value: 0,
		Right: new(tree.Tree),
	}
	t4 := newDescendingTree(1)

	if !Same(t1, t2) {
		t.Error("Same function did not return expected outcome.")
	}

	if Same(t1, t3) {
		t.Error("Same function did not return expected outcome.")
	}

	if Same(t1, t4) {
		t.Error("Same function did not return expected outcome.")
	}
}

func TestBinaryTreeLeftMostElem(t *testing.T) {
	t1 := tree.New(1)
	if v, err := leftMostElem(t1); err == nil && v != 1 {
		t.Error("Left most element was not returned as expected.")
	}
}

func TestBinaryTreeRightMostElem(t *testing.T) {
	t1 := tree.New(1)
	if v, err := rightMostElem(t1); err == nil && v != 10 {
		t.Error("Right most element was not returned as expected.")
	}
}
