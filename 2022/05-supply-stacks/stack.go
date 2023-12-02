package main

type Stack struct {
	items []string
}

func NewStack() *Stack {
	return &Stack{}
}

// push ...
func (s *Stack) push(elem string) {
	s.items = append(s.items, elem)
}

// pop ...
func (s *Stack) pop() string {
	if len(s.items) == 0 {
		return ""
	}
	result := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return result
}

// peek - pop without popping
func (s *Stack) peek() string {
	if len(s.items) == 0 {
		return ""
	}
	return s.items[len(s.items)-1]
}

// popFirst - returns first element of Stack - FIFO behaviour for Stack struct
func (s *Stack) popFirst() string {
	if len(s.items) == 0 {
		return ""
	}
	result := s.items[0]
	s.items = s.items[1:len(s.items)]
	return result
}

// isEmpty ...
func (s *Stack) isEmpty() bool {
	return len(s.items) == 0
}
