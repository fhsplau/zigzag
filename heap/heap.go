package heap

// Heap is heap
type Heap struct {
	next  int
	store []int
}

// Insert inserts element on the heap.
func (h *Heap) Insert(el int) {
	h.store[h.next] = el
	i := h.next
	p := h.GetParent(i)

	for h.store[p] > h.store[i] {
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
