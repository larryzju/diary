package stack

type ArrayStack struct {
	array []Element
	top   int
	cap   int
}

func NewArrayStack(cap int) *ArrayStack {
	return &ArrayStack{
		make([]Element, cap),
		0,
		cap,
	}
}

func (s *ArrayStack) Push(e Element) error {
	if s.top >= s.cap {
		return OverflowError
	}
	s.array[s.top] = e
	s.top += 1
	return nil
}

func (s *ArrayStack) Pop() (Element, error) {
	if s.top == 0 {
		return nil, UnderflowError
	}

	ele := s.array[s.top-1]
	s.top -= 1
	return ele, nil
}

func (s *ArrayStack) Empty() bool {
	return s.top == 0
}
