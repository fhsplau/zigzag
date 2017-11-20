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
	size      int
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
		swap(&(h.store), p, i)
		i = p
		p = h.GetParent(i)
	}

	h.next++
	h.size++
}

// GetParent gets parent of the i-th element
func (h *Heap) GetParent(i int) int {
	return (i - 1) / 2
}

// Size returns heap size
func (h *Heap) Size() int {
	return h.size
}

// Sort sorts the heap
func (h *Heap) Sort() []int {
	var sorted []int
	if h.Size() == 0 {
		return []int{}
	}

	sorted = append(sorted, h.store[0:h.next]...)
	swap := swapper(&sorted)
	nextCh := getNextChild(&sorted, h.doHeapify)
	pivot := len(sorted) - 1

	for pivot != 0 {
		swap(pivot, 0)
		parent := 0
		ch := nextCh(parent, pivot)

		for ch < pivot && h.doHeapify(sorted[parent], sorted[ch]) {
			swap(parent, ch)
			parent = ch
			ch = nextCh(parent, pivot)
		}

		pivot--
	}

	return sorted
}

func getNextChild(arr *[]int, f func(int, int) bool) func(int, int) int {
	return func(parent, pivot int) int {
		var ch int
		ch1 := 2*parent + 1
		ch2 := 2*parent + 2

		if ch2 >= pivot || f((*arr)[ch2], (*arr)[ch1]) {
			ch = ch1
		} else {
			ch = ch2
		}

		return ch
	}
}

func swapper(arr *[]int) func(int, int) {
	return func(first, second int) {
		tmp := (*arr)[first]
		(*arr)[first] = (*arr)[second]
		(*arr)[second] = tmp
	}
}

func swap(arr *[]int, first int, second int) {
	tmp := (*arr)[first]
	(*arr)[first] = (*arr)[second]
	(*arr)[second] = tmp
}
