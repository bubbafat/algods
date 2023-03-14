package algods

type Deque[T any] struct {
	data []T
}

func (q *Deque[T]) Count() int {
	return len(q.data)
}

func (q *Deque[T]) PeekHead() (bool, T) {
	if len(q.data) > 0 {
		return true, q.data[0]
	}

	var empty T
	return false, empty
}

func (q *Deque[T]) PeekTail() (bool, T) {
	if len(q.data) > 0 {
		return true, q.data[len(q.data)-1]
	}

	var empty T
	return false, empty
}

func (q *Deque[T]) DequeueHead() (bool, T) {
	var result T

	if len(q.data) == 0 {
		return false, result
	}

	result, q.data = q.data[0], q.data[1:]

	return true, result
}

func (q *Deque[T]) DequeueTail() (bool, T) {
	var result T

	if len(q.data) == 0 {
		return false, result
	}

	result, q.data = q.data[len(q.data)-1], q.data[:len(q.data)-1]

	return true, result
}

func (q *Deque[T]) EnqueueHead(value T) {
	q.data = append([]T{value}, q.data...)
}

func (q *Deque[T]) EnqueueTail(value T) {
	q.data = append(q.data, value)
}
