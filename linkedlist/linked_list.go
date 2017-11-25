package linkedlist

// LinkedList struct represents linked list
type LinkedList struct {
	next *LinkedList
	root int
	size int
}

// Size returns linked list's size
func (l *LinkedList) Size() int {
	return l.size
}

// Add adds a new element to the linked list
func (l *LinkedList) Add(el int) {
	if l.size == 0 {
		l.root = el
	} else {
		if l.next == nil {
			l.next = &LinkedList{}
		}
		l.next.Add(el)
	}
	l.size++
}

// Contain checks if a linked list contains given element.
func (l *LinkedList) Contain(el int) bool {
	if l.root == el {
		return true
	}
	if l.next != nil {
		return l.next.Contain(el)
	}
	return false
}
