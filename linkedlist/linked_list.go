package linkedlist

import (
	"fmt"
)

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
	current := l
	for {
		if current.Size() == 0 {
			current.root = el
			current.size++
			return
		}
		if current.next == nil {
			current.next = &LinkedList{}
		}
		current.size++
		current = current.next
	}
}

// Get gets
func (l *LinkedList) Get(index int) (int, error) {
	if index < 0 {
		return 0, fmt.Errorf("Index can't be negative")
	}

	if l.Size() == 0 {
		return 0, fmt.Errorf("Empty list")
	}

	if index >= l.Size() {
		return 0, fmt.Errorf("Index out of bounds")
	}

	current := l
	var count int

	for {
		if count == index {
			return current.root, nil
		}
		current = current.next
		count++
	}
}

// Remove removes
func (l *LinkedList) Remove(index int) error {
	if index < 0 {
		return fmt.Errorf("Index can't be negative")
	}

	if l.Size() == 0 {
		return fmt.Errorf("Empty list")
	}

	if index >= l.Size() {
		return fmt.Errorf("Index out of bounds")
	}
	return nil
}

// Contain checks if a linked list contains given element.
func (l *LinkedList) Contain(el int) bool {
	current := l
	for {
		if current == nil {
			return false
		}

		if current.root == el {
			return true
		}
		current = current.next
	}
}
