package algods

import "testing"

func TestDeque_EnqueueHead_Increases_Counter(t *testing.T) {
	q := new(Deque[int])

	for i := 0; i < 100; i++ {
		if count := q.Count(); count != i {
			t.Error("Deque count incorrect (expected, actual)", i, count)
		}

		q.EnqueueHead(i)
	}
}

func TestDeque_EnqueueTail_Increases_Counter(t *testing.T) {
	q := new(Deque[int])

	for i := 0; i < 100; i++ {
		if count := q.Count(); count != i {
			t.Error("Deque count incorrect (expected, actual)", i, count)
		}

		q.EnqueueTail(i)
	}
}

func TestDeque_EnqueueHead_Enqueues_Values(t *testing.T) {
	q := new(Deque[int])

	for i := 0; i < 100; i++ {
		q.EnqueueHead(i)

		found, value := q.PeekHead()

		if !found {
			t.Error("PeekHead should have returned a value")
		}

		if value != i {
			t.Error("PeekHead returned the wrong value (expected, actual)", i, value)
		}
	}
}

func TestDeque_EnqueueTail_Enqueues_Values(t *testing.T) {
	q := new(Deque[int])

	for i := 0; i < 100; i++ {
		q.EnqueueTail(i)

		found, value := q.PeekTail()

		if !found {
			t.Error("PeekTail should have returned a value")
		}

		if value != i {
			t.Error("PeekTail returned the wrong value (expected, actual)", i, value)
		}
	}
}

func TestDeque_DequeueHead_Dequeues_Values(t *testing.T) {
	q := new(Deque[int])

	for i := 0; i < 100; i++ {
		q.EnqueueHead(i)
	}

	for i := 99; i >= 0; i-- {
		found, value := q.DequeueHead()

		if !found {
			t.Error("DequeueHead should have returned a value")
		}

		if value != i {
			t.Error("DequeueHead returned the wrong value (expected, actual)", i, value)
		}
	}
}

func TestDeque_DequeueTail_Dequeues_Values(t *testing.T) {
	q := new(Deque[int])

	for i := 0; i < 100; i++ {
		q.EnqueueTail(i)
	}

	for i := 99; i >= 0; i-- {
		found, value := q.DequeueTail()

		if !found {
			t.Error("DequeueTail should have returned a value")
		}

		if value != i {
			t.Error("DequeueTail returned the wrong value (expected, actual)", i, value)
		}
	}
}
