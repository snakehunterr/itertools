package itertools

import "iter"

func NewRange(stop int) Range {
	return Range{
		Start: 0,
		Stop:  stop,
		Step:  1,
	}
}

type Range struct {
	Start int
	Stop  int
	Step  int
}

func (r Range) WithStart(start int) Range {
	return Range{
		Start: start,
		Stop:  r.Stop,
		Step:  r.Step,
	}
}

func (r Range) WithStep(step int) Range {
	return Range{
		Start: r.Start,
		Stop:  r.Stop,
		Step:  step,
	}
}

// implement Iterator interface (Iterator[int])
// for Range struct
func (r Range) Iter() iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := r.Start; i < r.Stop; i += r.Step {
			if !yield(i) {
				return
			}
		}
	}
}
