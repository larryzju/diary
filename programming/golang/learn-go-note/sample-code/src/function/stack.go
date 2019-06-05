package main

import (
  "fmt"
  "strconv"
)

type stack struct {
	i	int
	data	[10]int
}

func ( s *stack ) push ( k int ) {
	if s.i + 1 > 9 {
		return
	}

	s.data[s.i] = k
	s.i++
}

func ( s stack ) pop() int {
	s.i --
	ret := s.data[s.i]
	s.data[s.i] = 0
	return ret
}


func (s stack) String() string {
	var str string
	for i := 0; i < s.i; i++ {
		str = str + "[" +
		    strconv.Itoa(i) + ":" + strconv.Itoa( s.data[i] ) + "]"
	}
	return str
}

func main() {
	var s stack
	s.push(25)
	s.push(14)
	fmt.Printf( "stack %v\n", s )
}
