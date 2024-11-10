package itertools

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_map_pair(t *testing.T) {
	type testcase[K comparable, T any] struct {
		in     map[K]T
		out    map[K]T
		mapper func(K, T) (K, T)
	}

	{
		testcases := []testcase[int, int]{
			{
				map[int]int{1: 10, 2: 20, 3: 30},
				map[int]int{10: 1, 20: 2, 30: 3},
				func(k int, v int) (int, int) { return v, k },
			},
			{
				map[int]int{1: 1, 2: 2, 3: 3},
				map[int]int{1: 10, 2: 20, 3: 30},
				func(k int, v int) (int, int) { return k, v * 10 },
			},
			{
				map[int]int{1: 1, 2: 2, 3: 3},
				map[int]int{10: 1, 20: 2, 30: 3},
				func(k int, v int) (int, int) { return k * 10, v },
			},
		}

		for _, tc := range testcases {
			r := CollectPair(MapPair(FromMap(tc.in), tc.mapper))

			if diff := cmp.Diff(r, tc.out); diff != "" {
				t.Error(diff)
			}
		}
	}
	{
		testcases := []testcase[int, string]{
			{
				map[int]string{1: "bob", 2: "sam", 3: "jack"},
				map[int]string{1: "bobson", 2: "samson", 3: "jackson"},
				func(k int, v string) (int, string) { return k, v + "son" },
			},
			{
				map[int]string{1: "", 2: "bob", 3: "", 4: "jack"},
				map[int]string{1: "empty", 2: "bob", 3: "empty", 4: "jack"},
				func(k int, v string) (int, string) {
					value := ""
					if v == "" {
						value = "empty"
					} else {
						value = v
					}
					return k, value
				},
			},
		}

		for _, tc := range testcases {
			r := CollectPair(MapPair(FromMap(tc.in), tc.mapper))

			if diff := cmp.Diff(r, tc.out); diff != "" {
				t.Error(diff)
			}
		}
	}
}

func Test_filter_pair(t *testing.T) {
	type testcase[K comparable, V any] struct {
		in     map[K]V
		out    map[K]V
		filter func(K, V) bool
	}

	{
		testcases := []testcase[int, int]{
			{
				map[int]int{1: 10, 2: 11, 3: 12, 4: 20},
				map[int]int{1: 10, 3: 12, 4: 20},
				func(k, v int) bool { return v%2 == 0 },
			},
			{
				map[int]int{1: 1, 3: 3, 2: 3, 4: 3, 5: 9, 6: 5},
				map[int]int{},
				func(k, v int) bool { return v%2 == 0 },
			},
			{
				map[int]int{1: -2, 2: -1, 3: 0, 4: 1, 5: 2},
				map[int]int{4: 1, 5: 2},
				func(k, v int) bool { return v > 0 },
			},
		}

		for _, tc := range testcases {
			r := CollectPair(FilterPair(FromMap(tc.in), tc.filter))

			if diff := cmp.Diff(r, tc.out); diff != "" {
				t.Error(diff)
			}
		}
	}

	{
		testcases := []testcase[int, string]{
			{
				map[int]string{1: "foo", 2: "bar", 3: "jack"},
				map[int]string{3: "jack"},
				func(k int, v string) bool { return len(v) > 3 },
			},
			{
				map[int]string{1: "boo"},
				map[int]string{},
				func(k int, v string) bool { return k > 0 && v == "foo" },
			},
		}

		for _, tc := range testcases {
			r := CollectPair(FilterPair(FromMap(tc.in), tc.filter))

			if diff := cmp.Diff(r, tc.out); diff != "" {
				t.Error(diff)
			}
		}
	}
}
