package queue

import (
	"errors"
)

var (
	OverflowError  = errors.New("overflow")
	UnderflowError = errors.New("underflow")
)

type Element interface {
}

type Queue interface {
	Enqueue(e Element) error
	Dequeue() (e Element, err error)
}
