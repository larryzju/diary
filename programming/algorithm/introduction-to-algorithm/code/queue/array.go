package queue

type ArrayQueue struct {
	array []Element
	head  int
	tail  int
}

func NewArrayQueue(cap int) *ArrayQueue {
	return &ArrayQueue{
		make([]Element, cap+1),
		0,
		0,
	}
}

func (q *ArrayQueue) Len() int {
	return len(q.array)
}

func (q *ArrayQueue) Enqueue(e Element) error {
	if (q.tail+1)%q.Len() == q.head {
		return OverflowError
	}

	q.array[q.tail] = e
	q.tail = (q.tail + 1) % q.Len()
	return nil
}

func (q *ArrayQueue) Dequeue() (e Element, err error) {
	if q.tail == q.head {
		return nil, UnderflowError
	}

	e = q.array[q.head]
	q.head = (q.head + 1) % q.Len()
	return
}
