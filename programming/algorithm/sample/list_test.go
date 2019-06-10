package sample

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


func TestReverse(t *testing.T) {
	node4 := &List{4, nil}
	node3 := &List{3, node4}
	node2 := &List{2, node3}
	node1 := &List{1, node2}
	list := node1
	tail := Reverse(list)
	for p := tail; p != nil; p = p.next {
		t.Logf("%v -> ", p.v)
	}

	assert.Equal(t, tail, node4)
	assert.Equal(t, tail.next, node3)
	assert.Equal(t, tail.next.next, node2)
	assert.Equal(t, tail.next.next.next, node1)
	assert.Equal(t, tail.next.next.next.next, (*List)(nil))
}
