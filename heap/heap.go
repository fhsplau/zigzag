package heap

import (
	"fmt"
)

// Type is
type Type interface {
	GetHeapifyCondition() func(int, int) bool
}

// MinHeap is
type MinHeap struct{}

// GetHeapifyCondition gets
func (mh MinHeap) GetHeapifyCondition() func(int, int) bool {
	return func(p int, i int) bool {
		return p > i
	}
}

// MaxHeap is
type MaxHeap struct{}

// GetHeapifyCondition gets
func (mh MaxHeap) GetHeapifyCondition() func(int, int) bool {
	return func(p int, i int) bool {
		return i > p
	}
}

// Heap is heap
type Heap struct {
	next      int
	store     []int
	doHeapify func(int, int) bool
}

// NewHeap creates a new heap
func NewHeap(size int, t Type) Heap {
	store := make([]int, size)
	return Heap{
		store:     store,
		doHeapify: t.GetHeapifyCondition(),
	}
}

// Build builds heap from an given array
func Build(arr []int, t Type) Heap {
	heap := NewHeap(len(arr), t)
	for _, el := range arr {
		heap.Insert(el)
	}
	return heap
}

func (h Heap) String() string {
	s := "["

	for _, el := range h.store {
		s += fmt.Sprintf("%d, ", el)
	}

	return s + "]"
}

// Insert inserts a new element on the heap.
func (h *Heap) Insert(el int) {
	h.store[h.next] = el
	i := h.next
	p := h.GetParent(i)

	for h.doHeapify(h.store[p], h.store[i]) {
		tmp := h.store[p]
		h.store[p] = h.store[i]
		h.store[i] = tmp
		i = p
		p = h.GetParent(i)
	}

	h.next++
}

// GetParent gets parent of the i-th element
func (h *Heap) GetParent(i int) int {
	return (i - 1) / 2
}
