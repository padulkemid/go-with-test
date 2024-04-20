package generics

import "testing"

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) Push(val T)  {
	s.values = append(s.values, val)
}

func (s *Stack[T]) IsEmpty() bool  {
	return len(s.values) == 0
}

func (s *Stack[T]) Pop() (T, bool)  {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	index := len(s.values) - 1
	el := s.values[index]

	s.values = s.values[:index]

	return el, true
}

func assertEqual[T comparable](t testing.TB, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func assertNotEqual[T comparable](t testing.TB, got, want T)  {
	t.Helper()

	if got == want {
		t.Errorf("it shouldn't be the same, got %+v with %+v", got, want)
	}
}

func assertTrue(t testing.TB, got bool)  {
	t.Helper()

	if !got {
		t.Errorf("got %v want true", got)
	}
}

func assertFalse(t testing.TB, got bool)  {
	t.Helper()

	if got {
		t.Errorf("got %v want false", got)
	}
}

func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on int", func(t *testing.T) {
		assertEqual(t, 1, 1)
		assertNotEqual(t, 1, 2)
	})

	t.Run("assert in strings", func(t *testing.T) {
		assertEqual(t, "oi", "oi")
		assertNotEqual(t, "oi", "yeuwr")
	})

	t.Run("assert stacks", func(t *testing.T) {
		stackOfInts := new(Stack[int])

		// check if stack is empty
		assertTrue(t, stackOfInts.IsEmpty())

		// add a thing, then check its not empty
		stackOfInts.Push(231)
		assertFalse(t, stackOfInts.IsEmpty())

		// add another thing, pop it back again
		stackOfInts.Push(182)
		val, _ := stackOfInts.Pop()
		assertEqual(t, val, 182)

		val, _ = stackOfInts.Pop()
		assertEqual(t, val, 231)

		assertTrue(t, stackOfInts.IsEmpty())

		// can get the numbers we put in as numbers, not interface{}
		stackOfInts.Push(1)
		stackOfInts.Push(2)

		first, _ := stackOfInts.Pop()
		second, _ := stackOfInts.Pop()

		assertEqual(t, first + second, 3)
	})
}
