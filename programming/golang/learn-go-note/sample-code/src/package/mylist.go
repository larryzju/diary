package main

import (
  "errors"
  "fmt"
)

type Value int

type Node struct {
  Value
  next, prev *Node
}

type List struct {
  head, tail *Node
}

var ListEmptyError = errors.New( "List is Empty" )

func ( l *List ) Front () *Node {
  return l.head
}

func ( n *Node ) Next() *Node {
  return n.next
}

func New() *List {
  return &List{}
}

func (l *List) Push( v Value ) (n *Node) {
  n = &Node{Value:v}
  if l.head == nil {
    l.head = n
    l.tail = n
  } else {
    l.tail.next = n
    n.prev = l.tail
    l.tail = n
  }
  return
}

func (l *List) Pop() (v Value, e error ) {
  if l.tail != nil {
    n := l.tail
    v  = n.Value
    l.tail = n.prev
    if l.tail == nil {
      l.head = nil
    }
  } else {
    e = ListEmptyError
  }

  return
}



func main() {
  l := new(List)
  l.Push(1)
  l.Push(2)
  l.Push(3)

  for n := l.Front(); n != nil; n = n.Next() {
    fmt.Println( n.Value )
  }
}
