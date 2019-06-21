package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDoubleLinkedList(t *testing.T) {
	lst := NewDoubleLinkedList()
	node1 := lst.InsertAfter(lst.Head, 1)
	node2 := lst.InsertAfter(node1, 2)
	node3 := lst.InsertAfter(node1, 3)
	node4 := lst.InsertAfter(node2, 4)
	node5 := lst.InsertAfter(node3, 5)

	assert.Equal(t, node1, lst.Search(1))
	assert.Equal(t, node2, lst.Search(2))
	assert.Equal(t, node3, lst.Search(3))
	assert.Equal(t, node4, lst.Search(4))
	assert.Equal(t, node5, lst.Search(5))
	assert.Nil(t, nil, lst.Search(0))

	c := []Element{}
	for p := lst.Head.next; p != nil; p = p.next {
		c = append(c, p.value)
	}

	assert.Equal(t, 5, len(c))
	assert.Equal(t, c[0], 1)
	assert.Equal(t, c[1], 3)
	assert.Equal(t, c[2], 5)
	assert.Equal(t, c[3], 2)
	assert.Equal(t, c[4], 4)
}
