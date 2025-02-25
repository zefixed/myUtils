package Queue

import (
	"container/list"
	"errors"
)

type Queue[T comparable] struct {
	list *list.List
}

var EmptyQueue = errors.New("queue is empty")

func New[T comparable]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(val T) {
	if q.list == nil {
		q.list = list.New()
	}

	q.list.PushBack(val)
}

func (q *Queue[T]) Dequeue() (error, T) {
	if q.list == nil {
		var elem T
		return EmptyQueue, elem
	}

	elem := q.list.Front()
	q.list.Remove(elem)
	return nil, elem.Value.(T)
}

func (q *Queue[T]) Front() (error, T) {
	if q.list == nil {
		var elem T
		return EmptyQueue, elem
	}
	return nil, q.list.Front().Value.(T)
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list == nil
}

func (q *Queue[T]) Size() int {
	if q.list == nil {
		return 0
	}

	return q.list.Len()
}
