package itertools

import (
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_map(t *testing.T) {
	type testcase[T any] struct {
		in     []T
		out    []T
		mapper func(e T) T
	}

	for _, tc := range []testcase[int]{
		{
			in:     []int{1, 2, 3, 4, 5},
			out:    []int{10, 20, 30, 40, 50},
			mapper: func(elem int) int { return elem * 10 },
		},
		{
			in:     []int{},
			out:    []int{},
			mapper: func(elem int) int { return elem * 0 },
		},
		{
			in:     []int{1},
			out:    []int{1},
			mapper: func(elem int) int { return elem * elem },
		},
	} {
		r := Collect(Map(FromSlice(tc.in), tc.mapper))

		if diff := cmp.Diff(r, tc.out); diff != "" {
			t.Error(diff)
		}
	}

	for _, tc := range []testcase[string]{
		{
			[]string{"foo", "bar", "baz"},
			[]string{"FOO", "BAR", "BAZ"},
			func(s string) string { return strings.ToUpper(s) },
		},
		{
			[]string{},
			[]string{},
			func(s string) string { return "" },
		},
		{
			[]string{"f"},
			[]string{"b"},
			func(s string) string { return "b" },
		},
	} {
		r := Collect(Map(FromSlice(tc.in), tc.mapper))

		if diff := cmp.Diff(r, tc.out); diff != "" {
			t.Error(diff)
		}
	}
}

func Test_map_change(t *testing.T) {
	type testcase[T, K any] struct {
		in     []T
		out    []K
		mapper func(T) K
	}

	{
		testcases := []testcase[int, int]{
			{
				[]int{1, 2, 3},
				[]int{2, 3, 4},
				func(n int) int { return n + 1 },
			},
			{
				[]int{},
				[]int{},
				func(n int) int { return n },
			},
			{
				[]int{1, 2, 3, 4},
				[]int{1, 4, 9, 16},
				func(n int) int { return n * n },
			},
			{
				[]int{1},
				[]int{0},
				func(n int) int { return 0 },
			},
		}

		for _, tc := range testcases {
			r := Collect(MapChange(FromSlice(tc.in), tc.mapper))

			if diff := cmp.Diff(r, tc.out); diff != "" {
				t.Error(diff)
			}
		}
	}
	{
		mapper := strconv.Itoa
		testcases := []testcase[int, string]{
			{
				[]int{1, 2, 3},
				[]string{"1", "2", "3"},
				func(n int) string { return mapper(n) },
			},
			{
				[]int{1},
				[]string{"1"},
				func(n int) string { return mapper(n) },
			},
			{
				[]int{},
				[]string{},
				func(n int) string { return mapper(n) },
			},
		}

		for _, tc := range testcases {
			r := Collect(MapChange(FromSlice(tc.in), tc.mapper))

			if diff := cmp.Diff(r, tc.out); diff != "" {
				t.Error(diff)
			}
		}
	}
	{
		testcases := []testcase[int, byte]{
			{
				[]int{1, 2, 3},
				[]byte{1, 2, 3},
				func(n int) byte { return byte(n) },
			},
			{
				[]int{1},
				[]byte{1},
				func(n int) byte { return byte(n) },
			},
			{
				[]int{},
				[]byte{},
				func(n int) byte { return byte(n) },
			},
		}

		for _, tc := range testcases {
			r := Collect(MapChange(FromSlice(tc.in), tc.mapper))

			if diff := cmp.Diff(r, tc.out); diff != "" {
				t.Error(diff)
			}
		}
	}
}

func Test_filter(t *testing.T) {
	type testcase[T any] struct {
		in     []T
		out    []T
		filter func(T) bool
	}

	for _, tc := range []testcase[int]{
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			[]int{0, 2, 4, 6, 8, 10},
			func(n int) bool { return n%2 == 0 },
		},
		{
			[]int{1},
			[]int{},
			func(n int) bool { return n > 10 },
		},
		{
			[]int{-2, -3, 2, 3},
			[]int{2, 3},
			func(n int) bool { return n > 0 },
		},
	} {
		r := Collect(Filter(FromSlice(tc.in), tc.filter))

		if diff := cmp.Diff(r, tc.out); diff != "" {
			t.Error(diff)
		}
	}
}

func Test_zip(t *testing.T) {
	type testcase[K comparable, V any] struct {
		seq    []K
		joined []V
		expect map[K]V
	}

	for _, tc := range []testcase[int, int]{
		{
			[]int{1, 2, 3},
			[]int{10, 20, 30},
			map[int]int{1: 10, 2: 20, 3: 30},
		},
		{
			[]int{10, 20, 30},
			[]int{1, 2, 3},
			map[int]int{10: 1, 20: 2, 30: 3},
		},
		{
			[]int{1, 2, 3},
			[]int{10, 20, 30, 40, 50},
			map[int]int{1: 10, 2: 20, 3: 30},
		},
		{
			[]int{1, 2, 3, 4, 5},
			[]int{10, 20, 30},
			map[int]int{1: 10, 2: 20, 3: 30},
		},
		{
			[]int{},
			[]int{1, 2, 3},
			map[int]int{},
		},
		{
			[]int{1, 2, 3},
			[]int{},
			map[int]int{},
		},
	} {
		r := CollectPair(Zip(FromSlice(tc.seq), FromSlice(tc.joined)))

		if diff := cmp.Diff(r, tc.expect); diff != "" {
			t.Error(diff)
		}
	}
}

func Test_enumerate(t *testing.T) {
	type testcase[T any] struct {
		in  []T
		out map[int]T
	}

	for _, tc := range []testcase[int]{
		{
			[]int{1, 2, 3, 4, 5},
			map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5},
		},
		{
			[]int{10, 20, 30},
			map[int]int{0: 10, 1: 20, 2: 30},
		},
		{
			[]int{},
			map[int]int{},
		},
		{
			[]int{1},
			map[int]int{0: 1},
		},
	} {
		r := CollectPair(Enumerate(FromSlice(tc.in)))

		if diff := cmp.Diff(r, tc.out); diff != "" {
			t.Error(diff)
		}
	}
}
