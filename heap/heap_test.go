package heap

import "testing"

func TestCanAddRootElement(t *testing.T) {
	heap := NewHeap(1, MinHeap{})

	heap.Insert(1)

	checkIfMatches(heap.store[0], 1, t)
}

func TestCanAddElementToNonEmptyHeap(t *testing.T) {
	heap := NewHeap(2, MinHeap{})
	heap.Insert(1)
	heap.Insert(2)

	checkIfMatches(heap.store[1], 2, t)
}

func TestAddSmallerElement(t *testing.T) {
	heap := NewHeap(2, MinHeap{})
	heap.Insert(2)
	heap.Insert(1)

	checkIfMatches(heap.store[0], 1, t)
}

func TestNewElementIsTheSmallest_Left(t *testing.T) {
	heap := NewHeap(4, MinHeap{})

	heap.Insert(1)
	heap.Insert(2)
	heap.Insert(3)
	heap.Insert(-1)

	checkIfMatches(heap.store[0], -1, t)
	checkIfMatches(heap.store[1], 1, t)
	checkIfMatches(heap.store[2], 3, t)
	checkIfMatches(heap.store[3], 2, t)
}

func TestNewElementIsTheSmallest_Right(t *testing.T) {
	heap := NewHeap(10, MinHeap{})

	heap.Insert(1)
	heap.Insert(2)
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(5)
	heap.Insert(-1)

	checkIfMatches(heap.store[0], -1, t)
	checkIfMatches(heap.store[1], 2, t)
	checkIfMatches(heap.store[2], 1, t)
	checkIfMatches(heap.store[3], 4, t)
	checkIfMatches(heap.store[4], 5, t)
	checkIfMatches(heap.store[5], 3, t)
}

func TestGetParent_Left(t *testing.T) {
	store := make([]int, 2)
	heap := Heap{store: store}

	checkIfMatches(heap.GetParent(1), 0, t)
}

func TestGetParent_Right(t *testing.T) {
	store := make([]int, 10)
	heap := Heap{store: store}

	checkIfMatches(heap.GetParent(4), 1, t)
}

func checkIfMatches(is int, shouldBe int, t *testing.T) {
	if is != shouldBe {
		t.Logf("Should be %d, is: %d", shouldBe, is)
		t.Fail()
	}
}
