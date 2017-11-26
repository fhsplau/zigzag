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

func TestGetReturnsErrorIfListIsEmpty(t *testing.T) {
	l := LinkedList{}
	_, err := l.Get(0)

	if err == nil {
		t.Log("No error about empty list")
		t.Fail()
	}
}

func TestGetReturnsErrorIfIndexIsEqualToSize(t *testing.T) {
	l := LinkedList{size: 10}
	_, err := l.Get(10)

	if err == nil {
		t.Log("No error about index out of bounds")
		t.Fail()
	}
}

func TestGetReturnsErrorIfIndexIsGreaterThanSize(t *testing.T) {
	l := LinkedList{size: 10}
	_, err := l.Get(11)

	if err == nil {
		t.Log("No error about index out of bounds")
		t.Fail()
	}
}

func TestGetReturnsErrorIfIndexIsNegative(t *testing.T) {
	l := LinkedList{size: 10}
	_, err := l.Get(-1)

	if err == nil {
		t.Log("No error about negative index")
		t.Fail()
	}

}

func TestGetRootElement(t *testing.T) {
	l := LinkedList{}

	l.Add(10)

	el, _ := l.Get(0)
	areEqual(el, 10, t)
}

func TestGetOneOfTheNextElements(t *testing.T) {
	l := LinkedList{}
	l.Add(10)
	l.Add(20)
	l.Add(30)

	el, _ := l.Get(1)
	areEqual(el, 20, t)

	el, _ = l.Get(2)
	areEqual(el, 30, t)
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
		t.Logf("Should be %d, is %d", expected, is)
		t.Fail()
	}
}
