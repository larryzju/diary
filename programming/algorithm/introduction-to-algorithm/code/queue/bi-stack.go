package queue

import (
	"github.com/larryzju/ita/stack"
)

type BiStackQueue struct {
	a, b stack.Stack
	useA bool
}

func NewBiStackQueue(a, b stack.Stack) *BiStackQueue {
	return &BiStackQueue{a, b, true}
}

func (q *BiStackQueue) Enqueue(e Element) error {
	if q.a.Push(e) == stack.OverflowError {
		return OverflowError
	}
	return nil
}

func (q *BiStackQueue) Dequeue() (e Element, err error) {
	if e, err := q.b.Pop(); err == nil {
		return e, nil
	}

	for {
		ele, err := q.a.Pop()
		if err == stack.UnderflowError {
			break
		}

		if q.b.Push(ele) != nil {
			// assume b has equal or more space than a
			panic("no space in b")
		}
	}

	if e, err = q.b.Pop(); err == nil {
		return e, nil
	}

	return nil, UnderflowError
}
