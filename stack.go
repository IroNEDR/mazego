package mazego

import (
	"sync"
)

const (
	EmptyStackError = mazeError("Stack is empty")
)

type Stack[T any] struct {
	mutex   sync.Mutex
	content []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{sync.Mutex{}, make([]T, 0)}
}

func (s *Stack[T]) Push(t T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.content = append(s.content, t)
}

func (s *Stack[T]) Pop() (T, error) {
	var t T
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.content) == 0 {
		return t, EmptyStackError
	}

	t = s.content[len(s.content)-1]
	s.content = s.content[:len(s.content)-1]
	return t, nil
}

func (s *Stack[T]) Peek() (T, error) {
	var t T
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if len(s.content) == 0 {
		return t, EmptyStackError
	}

	return s.content[len(s.content)-1], nil
}

type Node[T any] struct {
	State     T
	Parent    *Node[T]
	Cost      float32
	Heuristic float32
}
