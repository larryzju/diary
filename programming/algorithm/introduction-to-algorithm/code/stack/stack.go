package stack

import (
	"errors"
)

var (
	UnderflowError = errors.New("underflow")
	OverflowError  = errors.New("overflow")
)

type Element interface {
}

type Stack interface {
	Push(e Element) error
	Pop() (Element, error)
	Empty() bool
}
