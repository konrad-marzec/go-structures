package linkedlist

type Node[T comparable] struct {
	prev  *Node[T]
	next  *Node[T]
	value T
}

func (n *Node[T]) Value() T {
	return n.value
}

type List[T comparable] struct {
	head *Node[T]
	tail *Node[T]
	size int
}

func New[T comparable]() *List[T] {
	return &List[T]{}
}

func (l *List[T]) Insert(value T) *Node[T] {
	n := &Node[T]{prev: l.tail, next: nil, value: value}

	if l.head == nil {
		l.head = n
	} else {
		l.tail.next = n
	}

	l.tail = n
	l.size += 1

	return n
}

func (l *List[T]) Remove(value T) {
	if l.head != nil {
		c := l.head
		for c != nil && c.value != value {
			c = c.next
		}

		if c == nil {
			return
		}

		if c.value == value {
			l.size -= 1

			if c.prev != nil {
				c.prev.next = c.next

				if c.next != nil {
					c.next.prev = c.prev
				} else {
					l.tail = c.prev
				}
			} else {
				l.head = c.next

				if l.head == nil {
					l.tail = c.next
				} else {
					l.head.prev = nil
				}
			}
		}
	}
}

func (l *List[T]) Shift() *Node[T] {
	r := l.head

	if l.head != nil {
		l.size -= 1
		l.head = l.head.next

		if l.head == nil {
			l.tail = l.head
		} else {
			l.head.prev = nil
		}
	}

	return r
}

func (l *List[T]) Pop() *Node[T] {
	r := l.tail

	if l.tail != nil {
		l.size -= 1

		l.tail = l.tail.prev

		if l.tail == nil {
			l.head = l.tail
		} else {
			l.tail.next = nil
		}
	}

	return r
}

func (l *List[T]) Size() int {
	return l.size
}
