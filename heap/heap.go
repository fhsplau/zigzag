package heap

import (
	"fmt"
)

// Comparator is
type Comparator interface {
	Compare(int, int) bool
}

// MinHeap is
type MinHeap struct{}

// Compare gets
func (mh MinHeap) Compare(parent, child int) bool {
	return parent > child
}

// MaxHeap is
type MaxHeap struct{}

// Compare gets
func (mh MaxHeap) Compare(parent, child int) bool {
	return child > parent
}

// Heap is heap
type Heap struct {
	next       int
	store      []int
	size       int
	comparator Comparator
}

// NewHeap creates a new heap
func NewHeap(size int, t Comparator) Heap {
	store := make([]int, size)
	return Heap{
		store:      store,
		comparator: t,
	}
}

// Build builds heap from an given array
func Build(arr []int, c Comparator) Heap {
	heap := NewHeap(len(arr), c)
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

	for h.comparator.Compare(h.store[p], h.store[i]) {
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
	nextCh := getNextChild(&sorted, h.comparator)
	pivot := len(sorted) - 1

	for pivot != 0 {
		swap(pivot, 0)
		parent := 0
		ch := nextCh(parent, pivot)

		for ch < pivot && h.comparator.Compare(sorted[parent], sorted[ch]) {
			swap(parent, ch)
			parent = ch
			ch = nextCh(parent, pivot)
		}

		pivot--
	}

	return sorted
}

func getNextChild(arr *[]int, c Comparator) func(int, int) int {
	return func(parent, pivot int) int {
		var ch int
		ch1 := 2*parent + 1
		ch2 := 2*parent + 2

		if ch2 >= pivot || c.Compare((*arr)[ch2], (*arr)[ch1]) {
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
