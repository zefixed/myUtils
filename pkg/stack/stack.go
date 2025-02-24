package stack

import (
	"errors"
)

type Stack[T comparable] struct {
	data []T
}

var (
	EmptyStack = errors.New("stack is empty")
	NotInStack = errors.New("not in stack")
)

func (s *Stack[T]) Push(elem T) {
	s.data = append(s.data, elem)
}

func (s *Stack[T]) Pop() (error, T) {
	var elem T
	if len(s.data) == 0 {
		return EmptyStack, elem
	}

	elem = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return nil, elem
}

func (s *Stack[T]) Top() (error, T) {
	var elem T
	if len(s.data) == 0 {
		return EmptyStack, elem
	}

	elem = s.data[len(s.data)-1]
	return nil, elem
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func (s *Stack[T]) Clear() {
	s.data = []T{}
}

func (s *Stack[T]) Contains(elem T) bool {
	if len(s.data) == 0 {
		return false
	}

	for _, el := range s.data {
		if el == elem {
			return true
		}
	}

	return false
}
