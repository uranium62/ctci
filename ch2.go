package ctci

import (
	"bytes"
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type List struct {
	Head   *Node
	length int
}

func (l *List) Push(v int) *List {
	n := new(Node)
	n.Value = v
	n.Next = l.Head
	l.Head = n
	l.length++
	return l
}

func (l *List) Pop() (int, bool) {
	if l.Head == nil {
		return 0, false
	}
	n := l.Head
	l.Head = n.Next
	l.length--
	return n.Value, true
}

func (l *List) String() string {
	var buf bytes.Buffer

	for n := l.Head; n != nil; n = n.Next {
		buf.WriteString(fmt.Sprintf("%v", n.Value))
		buf.WriteString(" -> ")
	}
	buf.WriteString("nil")

	return buf.String()
}

func (l *List) Equals(a *List) bool {
	if l.length != a.length {
		return false
	}

	var i, j *Node = nil, nil

	for i, j = l.Head, a.Head; i != nil && j != nil; i, j = i.Next, j.Next {
		if i.Value != j.Value {
			return false
		}
	}

	if i == nil && j == nil {
		return true
	} else {
		return false
	}

}

func NewList(items []int) *List {
	list := new(List)
	for _, it := range items {
		list.Push(it)
	}
	return list
}

func (l *List) RemoveDups() {
	table := make(map[int]bool)
	prev := l.Head

	for n := l.Head; n != nil; n = n.Next {
		if _, ok := table[n.Value]; ok {
			prev.Next = n.Next
			l.length--
		} else {
			table[n.Value] = true
			prev = n
		}
	}

}

func (l *List) NthToLast(k int) (*Node, bool) {
	if k < 0 {
		return nil, false
	}

	p1 := l.Head
	p2 := l.Head

	for i := 0; i < k-1; i++ {
		if p2 == nil {
			return nil, false
		}
		p2 = p2.Next
	}

	if p2 == nil {
		return nil, false
	}

	for p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1, true
}

func (l *List) DeleteNode(n *Node) bool {
	if n == nil || n.Next == nil {
		return false
	}

	n.Value = n.Next.Value
	n.Next = n.Next.Next
	return true
}

func (l *List) Partition(v int) {
	before := &List{}
	after := &List{}

	for e := l.Head; e != nil; {
		next := e.Next
		if e.Value < v {
			e.Next = before.Head
			before.Head = e
		} else {
			e.Next = after.Head
			after.Head = e
		}
		e = next
	}

	if before == nil {
		l.Head = after.Head
		return
	}
	l.Head = before.Head
	e := l.Head
	for e.Next != nil {
		e = e.Next
	}
	e.Next = after.Head
}
