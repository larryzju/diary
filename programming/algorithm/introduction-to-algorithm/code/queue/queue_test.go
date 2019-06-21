package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testQueue(t *testing.T, q Queue) {
	var ele Element
	var err error

	ele, err = q.Dequeue()
	assert.Equal(t, UnderflowError, err)

	assert.NoError(t, q.Enqueue(1))
	assert.NoError(t, q.Enqueue(2))
	assert.NoError(t, q.Enqueue(3))
	assert.NoError(t, q.Enqueue(4))
	assert.Equal(t, OverflowError, q.Enqueue(5))

	ele, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 1, ele)

	ele, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 2, ele)

	assert.NoError(t, q.Enqueue(6))
	assert.NoError(t, q.Enqueue(7))
	assert.Equal(t, OverflowError, q.Enqueue(8))

	ele, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 3, ele)

	ele, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 4, ele)

	ele, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 6, ele)

	ele, err = q.Dequeue()
	assert.NoError(t, err)
	assert.Equal(t, 7, ele)

	ele, err = q.Dequeue()
	assert.Equal(t, UnderflowError, err)
}

func TestQueue(t *testing.T) {
	q1 := NewArrayQueue(4)
	testQueue(t, q1)
}
