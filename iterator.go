package itertools

import "iter"

type Iterator[T any] interface {
	Iter() iter.Seq[T]
}
