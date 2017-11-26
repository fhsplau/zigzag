package linkedlist

import (
	"testing"
)

func TestSizeOfEmptyLinkedListIsZero(t *testing.T) {
	l := LinkedList{}

	areEqual(l.Size(), 0, t)
}

func TestAfterAddingOneElementSizeIsOne(t *testing.T) {
	l := LinkedList{}

	l.Add(10)

	areEqual(l.Size(), 1, t)
}

func TestAfterAddingTwoElementsSizeIsTwo(t *testing.T) {
	l := LinkedList{}

	l.Add(10)
	l.Add(11)

	areEqual(l.Size(), 2, t)
}

func TestEmptyLinkedListDoesNotContainAnElement(t *testing.T) {
	l := LinkedList{}

	if l.Contain(10) {
		t.Log("Emtpy list should not contain any element")
		t.Fail()
	}
}

func TestAddToEmptyLinkedList(t *testing.T) {
	l := LinkedList{}

	l.Add(10)

	areEqual(l.root, 10, t)
}

func TestAddElementToNotEmptyList(t *testing.T) {
	l := LinkedList{}

	l.Add(10)
	l.Add(9)

	areEqual(l.next.root, 9, t)
}

func TestSizeOfNextIsOneSmallerThanTheWhole(t *testing.T) {
	l := LinkedList{}
	l.Add(10)
	l.Add(9)
	areEqual(l.next.Size(), 1, t)
}

func TestNotEmptyListDoesNotContainAnElement(t *testing.T) {
	l := LinkedList{}
	l.Add(10)
	if l.Contain(11) {
		t.Log("List should not contain element")
		t.Fail()
	}
}

func TestLinkedListContainsRootElement(t *testing.T) {
	l := LinkedList{}
	l.Add(10)

	if !l.Contain(10) {
		t.Log("List should contain root element")
		t.Fail()
	}
}

func TestLinkedListContainsElement(t *testing.T) {
	l := LinkedList{}
	l.Add(10)
	l.Add(11)
	l.Add(100)

	if !l.Contain(100) {
		t.Log("List should contain 100")
		t.Fail()
	}
}

func BenchmarkAddToList(t *testing.B) {
	l := LinkedList{}

	for i := 0; i < t.N; i++ {
		l.Add(10)
	}
}

func BenchmarkContainElement(b *testing.B) {
	c := 1000
	l := LinkedList{}

	for i := 0; i < c; i++ {
		l.Add(i)
	}

	for n := 0; n < b.N; n++ {
		l.Contain(c + 1)
	}
}

func areEqual(is, expected int, t *testing.T) {
	if is != expected {
		t.Logf("Should be %d, is %d", is, expected)
		t.Fail()
	}
}
