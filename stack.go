package faq

type Stack struct {
	items []interface{}
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(item interface{}) {
	newItems := append(s.items, item)
	s.items = newItems
}

func (s *Stack) Peek(last int) interface{} {
	itemCount := len(s.items)
	peekIndex := (itemCount - last) - 1

	if peekIndex > itemCount || peekIndex < 0 {
		return nil
	}

	return s.items[peekIndex]
}

func (s *Stack) Pop() interface{} {
	lastIndex := len(s.items) - 1
	popped := s.items[lastIndex]
	s.items = s.items[:len(s.items)-1]

	return popped
}

func (s *Stack) Length() int {
	return len(s.items)
}
