package linkedlist

import (
	"testing"
)

func TestSizeOfEmptyLinkedListIsZero(t *testing.T) {
	l := New()

	shouldBeEqual(l.Size(), 0, t)
}

func TestAfterAddingOneElementSizeIsOne(t *testing.T) {
	l := New(10)
	shouldBeEqual(l.Size(), 1, t)
}

func TestAfterAddingTwoElementsSizeIsTwo(t *testing.T) {
	l := New(10, 11)
	shouldBeEqual(l.Size(), 2, t)
}

func TestEmptyLinkedListDoesNotContainAnElement(t *testing.T) {
	l := LinkedList{}

	if l.Contain(10) {
		t.Log("Emtpy list should not contain any element")
		t.Fail()
	}
}

func TestAddToEmptyLinkedList(t *testing.T) {
	l := New()
	l.Add(10)
	shouldBeEqual(l.root, 10, t)
}

func TestAddElementToNotEmptyList(t *testing.T) {
	l := New()
	l.Add(10)
	l.Add(9)
	shouldBeEqual(l.next.root, 9, t)
}

func TestSizeOfNextIsOneSmallerThanTheWhole(t *testing.T) {
	l := New(10, 9)
	shouldBeEqual(l.next.Size(), 1, t)
}

func TestNotEmptyListDoesNotContainAnElement(t *testing.T) {
	l := New(10)
	if l.Contain(11) {
		t.Log("List should not contain element")
		t.Fail()
	}
}

func TestLinkedListContainsRootElement(t *testing.T) {
	l := New(10)

	if !l.Contain(10) {
		t.Log("List should contain root element")
		t.Fail()
	}
}

func TestLinkedListContainsElement(t *testing.T) {
	l := New(10, 11, 100)
	if !l.Contain(100) {
		t.Log("List should contain 100")
		t.Fail()
	}
}

func TestGetReturnsErrorIfListIsEmpty(t *testing.T) {
	l := LinkedList{}
	_, err := l.Get(0)

	shouldBeError(err, "empty list", t)
}

func TestGetReturnsErrorIfIndexIsEqualToSize(t *testing.T) {
	l := LinkedList{size: 10}
	_, err := l.Get(10)

	shouldBeError(err, "index out of bounds", t)
}

func TestGetReturnsErrorIfIndexIsGreaterThanSize(t *testing.T) {
	l := LinkedList{size: 10}
	_, err := l.Get(11)

	shouldBeError(err, "index out of bounds", t)
}

func TestGetReturnsErrorIfIndexIsNegative(t *testing.T) {
	l := LinkedList{size: 10}
	_, err := l.Get(-1)

	shouldBeError(err, "negative index", t)
}

func TestGetRootElement(t *testing.T) {
	l := New(10)
	el, _ := l.Get(0)
	shouldBeEqual(el, 10, t)
}

func TestGetOneOfTheNextElements(t *testing.T) {
	l := New(10, 20, 30)
	el, _ := l.Get(1)
	shouldBeEqual(el, 20, t)

	el, _ = l.Get(2)
	shouldBeEqual(el, 30, t)
}

func TestRemoveFromEmptyList(t *testing.T) {
	l := LinkedList{}

	err := l.Remove(0)
	shouldBeError(err, "empty list", t)
}

func TestRemoveNegativeIndex(t *testing.T) {
	l := LinkedList{size: 10}

	err := l.Remove(-1)
	shouldBeError(err, "negative index", t)
}

func TestRemoveIndexEqualToSize(t *testing.T) {
	l := LinkedList{size: 10}

	err := l.Remove(10)
	shouldBeError(err, "index out of bounds", t)
}

func TestRemoveIndexGreaterThanSize(t *testing.T) {
	l := LinkedList{size: 10}
	err := l.Remove(11)

	shouldBeError(err, "index out of bounds", t)
}

func TestRemoveRootElement(t *testing.T) {
	l := New(10)
	l.Remove(0)
	if l.Contain(10) {
		t.Log("Root should be removed")
		t.Fail()
	}
}

func TestRemoveRootElementInBiggerList(t *testing.T) {
	l := New(10, 20)
	l.Remove(0)
	shouldBeEqual(l.root, 20, t)
}

func TestRemoveMiddleElement(t *testing.T) {
	l := New(10, 20, 30)
	l.Remove(1)
	if l.Contain(20) {
		t.Log("Should not contain 20")
		t.Fail()
	}

	shouldBeEqual(l.next.root, 30, t)
}

func TestAfterRemovingRootSizeIsZero_OneElement(t *testing.T) {
	l := New(10)
	l.Remove(0)

	shouldBeEqual(l.Size(), 0, t)
}

func TestAfterRemovingRootSizeIsOne_TwoElements(t *testing.T) {
	l := New(10, 11)
	l.Remove(0)

	shouldBeEqual(l.Size(), 1, t)
}

func TestSizeAfterRemovingMiddleElement(t *testing.T) {
	l := New(10, 11, 12)

	l.Remove(1)
	shouldBeEqual(l.Size(), 2, t)
	shouldBeEqual(l.next.Size(), 1, t)
}

func TestSizeAfterRemovingLastElement(t *testing.T) {
	l := New(10, 11, 12)

	l.Remove(2)
	shouldBeEqual(l.Size(), 2, t)
	shouldBeEqual(l.next.Size(), 1, t)
}

func TestStringEmptyList(t *testing.T) {
	l := New()

	s := l.String()
	if s != "nil" {
		t.Log("Should be nil, is ", s)
		t.Fail()
	}
}

func TestStringListWithElements(t *testing.T) {
	l := New(10, 11)

	s := l.String()
	if s != "10 -> 11 -> nil" {
		t.Log("Should be 10 -> 11 -> nil, is ", s)
		t.Fail()
	}
}

func BenchmarkAddToList(t *testing.B) {
	l := New()

	for i := 0; i < t.N; i++ {
		l.Add(10)
	}
}

func BenchmarkContainElement(b *testing.B) {
	c := 1000
	l := New()

	for i := 0; i < c; i++ {
		l.Add(i)
	}

	for n := 0; n < b.N; n++ {
		l.Contain(c + 1)
	}
}

func shouldBeEqual(is, expected int, t *testing.T) {
	if is != expected {
		t.Logf("Should be %d, is %d", expected, is)
		t.Fail()
	}
}

func shouldBeError(err error, msg string, t *testing.T) {
	if err == nil {
		t.Log("No error about ", msg)
		t.Fail()
	}
}
