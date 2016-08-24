package ctci

import (
	"testing"
)

func TestList_RemoveDups(t *testing.T) {
	cases := []struct {
		input  *List
		expect *List
	}{
		{NewList([]int{}),
			NewList([]int{})},
		{NewList([]int{1}),
			NewList([]int{1})},
		{NewList([]int{1, 2, 3}),
			NewList([]int{1, 2, 3})},
		{NewList([]int{1, 2, 2, 3}),
			NewList([]int{1, 2, 3})},
		{NewList([]int{1, 1, 1, 2, 3, 3}),
			NewList([]int{1, 2, 3})},
	}

	for _, it := range cases {
		it.input.RemoveDups()

		if !it.input.Equals(it.expect) {
			t.Errorf("RemoveDups: input %q, not equals expect %q", it.input, it.expect)
		}
	}
}

func TestList_NthToLast(t *testing.T) {
	list := NewList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	cases := []struct {
		k   int
		res bool
		val int
	}{
		{-1, false, 0},
		{10, false, 0},
		{2, true, 2},
		{9, true, 9},
	}
	for _, it := range cases {

		n, ok := list.NthToLast(it.k)

		if ok != it.res {
			t.Errorf("NthToLast: input %q, not contains NthToLast(%d)", list, it.k)
		}
		if (n != nil) && n.Value != it.val {
			t.Errorf("NthToLast: input %q, the NthToLast{%d} return %d, but expect %d", list, it.k, n.Value, it.val)
		}
	}
}

func TestList_DeleteNode(t *testing.T) {
	cases := []struct {
		input  *List
		expect *List
	}{
		{NewList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}),
			NewList([]int{1, 2, 3, 4, 6, 7, 8, 9})},
	}

	for _, it := range cases {
		node, ok := it.input.NthToLast(5)
		ok = it.input.DeleteNode(node)

		if !ok && !it.input.Equals(it.expect) {
			t.Errorf("DeleteNode: input %q, not equals expect %q", it.input, it.expect)
		}
	}

}
