package linkedlist

import (
	"errors"
	"fmt"
)

type Node[T comparable] struct {
	Val  T
	Next *Node[T]
}

var (
	EmptyList = errors.New("list is empty")
	NotInList = errors.New("value is not in the list")
)

func (n *Node[T]) Append(val T) error {
	if n == nil {
		return EmptyList
	}

	for n.Next != nil {
		n = n.Next
	}

	n.Next = &Node[T]{Val: val}
	return nil
}

func (n *Node[T]) AppendAfter(val, afterVal T) error {
	if n == nil {
		return EmptyList
	}

	for n != nil {
		if n.Val == afterVal {
			n.Next = &Node[T]{Val: val, Next: n.Next}
			return nil
		}
		n = n.Next
	}

	return NotInList
}

func (n *Node[T]) AppendBefore(val, beforeVal T) (error, *Node[T]) {
	head := n
	if head == nil {
		return EmptyList, head
	}

	if head.Val == beforeVal {
		return nil, &Node[T]{Val: val, Next: head}
	}

	prev := head
	n = n.Next
	for n != nil {
		if n.Val == beforeVal {
			prev.Next = &Node[T]{Val: val, Next: n}
			return nil, head
		}
		prev = n
		n = n.Next
	}

	return NotInList, head
}

func (n *Node[T]) Prepend(val T) *Node[T] {
	head := n

	if head == nil {
		head = &Node[T]{Val: val}
		return head
	}

	head = &Node[T]{Val: val, Next: n}
	return head
}

func (n *Node[T]) String(separator string) string {
	if n == nil {
		return ""
	}

	var res string
	for n != nil {
		res += fmt.Sprintf("%v", n.Val)
		if n.Next != nil {
			res += fmt.Sprintf("%v", separator)
		} else {
			res += "\n"
		}
		n = n.Next
	}

	return res
}

func (n *Node[T]) Len() int {
	if n == nil {
		return 0
	}

	ln := 1
	for n.Next != nil {
		ln++
		n = n.Next
	}

	return ln
}

func (n *Node[T]) Index(val T) (error, int) {
	if n == nil {
		return EmptyList, -1
	}

	idx := 0
	for n != nil {
		if n.Val == val {
			return nil, idx
		}
		idx++
		n = n.Next
	}

	return NotInList, -1
}
