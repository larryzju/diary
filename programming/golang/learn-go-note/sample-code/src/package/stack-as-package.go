package stack

type Stack struct{
	n    int
	data [10]int
}

func (s *Stack) Push( v int ) {
	if s.n+1 > 10 {
		return
	}

	s.data[s.n] = v
	s.n++
}

func (s *Stack) Pop()( v int ) {
	s.n--
	v = s.data[s.n]
	return
}
