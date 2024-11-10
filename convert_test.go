package itertools

import (
	"iter"
	"testing"
)

func Test_from_slice(t *testing.T) {
	testcases := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3},
		{1},
		{1, 2},
		{2, 1},
		{},
	}

	for _, tc := range testcases {
		seq := FromSlice(tc)

		var (
			index      int
			next, stop = iter.Pull(seq)
		)

		for {
			n, ok := next()

			if !ok {
				break
			}
			if index >= len(tc) {
				break
			}

			num := tc[index]

			if num != n {
				t.Errorf("index: %d\n[slice]: %d\n[iter]: %d", index, num, n)
			}

			index++
		}

		if index < len(tc) {
			t.Errorf("slice len: %d; iter len: %d", len(tc), index)
		}
		if n, ok := next(); ok == true {
			for ok == true {
				t.Error("unexpected elem from iter:", n)
				n, ok = next()
			}
		}

		stop()
	}
}

func Test_from_map(t *testing.T) {
	testcases := []map[int]int{
		{0: 1, 1: 2, 2: 3},
		{10: 1, 20: 2, 30: 3},
		{100: 10, 200: 10, 300: 30},
		{100: 10, 200: 100, 300: 100},
		{},
		{0: 1},
		{10: 10},
		{20: 0},
	}

	for _, tc := range testcases {
		seq := FromMap(tc)

		var (
			count      int
			next, stop = iter.Pull2(seq)
		)

		for k, v, ok := next(); ok == true; k, v, ok = next() {
			count++

			if v2, exists := tc[k]; !exists {
				t.Error("key not exists in original map:", k)
			} else if v2 != v {
				t.Errorf("on key %v value in original map: %v, in iter: %v", k, v2, v)
			}
		}

		if count != len(tc) {
			t.Errorf("length of map: %v, length of iter: %v", len(tc), count)
		}

		stop()
	}
}
