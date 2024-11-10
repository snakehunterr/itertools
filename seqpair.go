package itertools

import (
	"iter"
)

func MapPair[K, V any](seq iter.Seq2[K, V], mapper func(K, V) (K, V)) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !yield(mapper(k, v)) {
				return
			}
		}
	}
}

func FilterPair[K, V any](seq iter.Seq2[K, V], filter func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if !filter(k, v) {
				continue
			}

			if !yield(k, v) {
				return
			}
		}
	}
}

func CollectPair[K comparable, V any](seq iter.Seq2[K, V]) map[K]V {
	m := map[K]V{}

	for k, v := range seq {
		m[k] = v
	}

	return m
}
