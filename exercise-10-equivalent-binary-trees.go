package tutorial

import (
	"regexp"
	"strconv"
	"strings"

	"golang.org/x/tour/tree"
)

// Walk traverses the tree from left to right and writes its elements to ch.
// It does not close so the channel is not rangable.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same compares two trees for equality in the manner of reflect.DeepEqual.  It
// converts trees to flat strings for comparison, which is not elegant but
// functional.
func Same(t1, t2 *tree.Tree) bool {
	s1 := toSliceOfStrings(t1)
	s2 := toSliceOfStrings(t2)

	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func toSliceOfStrings(t *tree.Tree) []string {
	s := regexp.
		MustCompile("[()]").
		ReplaceAllString(t.String(), " ")
	return strings.Fields(s)
}

func leftMostElem(t *tree.Tree) (int, error) {
	s := toSliceOfStrings(t)
	return strconv.Atoi(s[0])
}

func rightMostElem(t *tree.Tree) (int, error) {
	s := toSliceOfStrings(t)
	return strconv.Atoi(s[len(s)-1])
}
