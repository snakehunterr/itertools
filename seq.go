package itertools

import (
	"iter"
	"math"
)

// Map applies the provided mapper function to each element in the input sequence,
// yielding a new sequence with the mapped elements.
func Map[T any](seq iter.Seq[T], mapper func(T) T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := range seq {
			if !yield(mapper(e)) {
				return
			}
		}
	}
}

// Filter returns a new sequence that contains only the elements from the input
// sequence that pass the provided filter function.
func Filter[T any](seq iter.Seq[T], filter func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := range seq {
			if !filter(e) {
				continue
			}
			if !yield(e) {
				return
			}
		}
	}
}

// Collect takes an input sequence and returns a slice containing all the elements
// from the sequence. This is a convenience function that allows you to easily
// convert a sequence into a slice.
func Collect[T any](seq iter.Seq[T]) []T {
	s := make([]T, 0, 1<<10)

	for e := range seq {
		s = append(s, e)
	}

	return s
}

func MapChange[T, K any](seq iter.Seq[T], mapper func(T) K) iter.Seq[K] {
	return func(yield func(K) bool) {
		for e := range seq {
			if !yield(mapper(e)) {
				return
			}
		}
	}
}

func Zip[T, K any](seq iter.Seq[T], other iter.Seq[K]) iter.Seq2[T, K] {
	return func(yield func(T, K) bool) {
		snext, sstop := iter.Pull(seq)
		onext, ostop := iter.Pull(other)

		defer func() {
			sstop()
			ostop()
		}()

		var (
			elem1 T
			elem2 K
			ok    bool
		)
		for {
			elem1, ok = snext()
			if !ok {
				return
			}

			elem2, ok = onext()
			if !ok {
				return
			}

			if !yield(elem1, elem2) {
				return
			}
		}
	}
}

func Enumerate[T any](seq iter.Seq[T]) iter.Seq2[int, T] {
	return Zip(NewRange(math.MaxInt).Iter(), seq)
}
