package algods

type listNode[T comparable] struct {
	Next *listNode[T]
	Prev *listNode[T]
	Data T
}

type List[T comparable] struct {
	head   *listNode[T]
	tail   *listNode[T]
	Length int
}

func (l *List[T]) AddHead(value T) {
	var n = new(listNode[T])
	n.Data = value

	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		n.Next = l.head
		l.head.Prev = n
		l.head = n
	}

	l.Length++
}

func (l *List[T]) AddTail(value T) {
	var n = new(listNode[T])
	n.Data = value

	if l.head == nil {
		l.head = n
		l.tail = n
	} else {
		l.tail.Next = n
		l.tail = n
	}

	l.Length++
}

func (l *List[T]) Remove(value T) bool {
	for n := l.head; n != nil; n = n.Next {
		if value == n.Data {
			l.Length--

			// removing the only node
			if n.Prev == nil && n.Next == nil {
				l.head = nil
				l.tail = nil
				return true
			}

			// removing the head node
			if n.Prev == nil {
				l.head = n.Next
				l.head.Prev = nil
				return true
			}

			// remove the tail node
			if n.Next == nil {
				l.tail = n.Prev
				l.tail.Next = nil
				return true
			}

			n.Prev.Next = n.Next
			n.Next.Prev = n.Prev
			return true
		}
	}

	return false
}

func (l *List[T]) ForEach(action func(value T)) {
	for n := l.head; n != nil; n = n.Next {
		action(n.Data)
	}
}
