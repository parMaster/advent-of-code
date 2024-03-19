package main

type Stack[T any] struct {
	items []T
}

// Returns new instance of Stack
func NewStack[T any](elem T) Stack[T] {
	items := make([]T, 0)
	return Stack[T]{items: items}
}

// Push appends elem to stack
func (s *Stack[T]) Push(elem T) {
	s.items = append(s.items, elem)
}

// Pop returns the last item of the stack, stack is reduced by the last element
func (s *Stack[T]) Pop() *T {
	if len(s.items) == 0 {
		return nil
	}
	result := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return &result
}

// Peek - returns the last item ot he stack, stack itself is untouched
func (s *Stack[T]) Peek() *T {
	if len(s.items) == 0 {
		return nil
	}
	return &s.items[len(s.items)-1]
}

// PopFirst - returns first element of Stack - FIFO behaviour for Stack struct
func (s *Stack[T]) PopFirst() *T {
	if len(s.items) == 0 {
		return nil
	}
	result := s.items[0]
	s.items = s.items[1:len(s.items)]
	return &result
}

// IsEmpty returns true if stack has no elements
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}
