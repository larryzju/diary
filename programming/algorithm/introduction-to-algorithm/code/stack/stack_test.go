package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testStack5(t *testing.T, s Stack) {
	var err error
	var ele Element

	assert.NoError(t, s.Push(1))
	assert.NoError(t, s.Push(2))
	assert.NoError(t, s.Push(3))
	assert.NoError(t, s.Push(4))
	assert.NoError(t, s.Push(5))
	assert.Equal(t, OverflowError, s.Push(6))

	ele, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 5, ele)

	ele, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 4, ele)

	assert.NoError(t, s.Push(6))
	assert.NoError(t, s.Push(7))

	assert.Equal(t, OverflowError, s.Push(8))
	assert.False(t, s.Empty())

	ele, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 7, ele)

	ele, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 6, ele)

	assert.False(t, s.Empty())

	ele, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 3, ele)

	ele, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 2, ele)

	ele, err = s.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 1, ele)

	assert.True(t, s.Empty())

	ele, err = s.Pop()
	assert.Equal(t, UnderflowError, err)
	assert.Nil(t, ele)

	assert.True(t, s.Empty())
}

func TestStack(t *testing.T) {
	s := NewArrayStack(5)
	testStack5(t, s)
}
