package itertools

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_range(t *testing.T) {
	type testcase struct {
		r Iterator[int]
		s []int
	}
	testcases := []testcase{
		{
			NewRange(10),
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			NewRange(5),
			[]int{0, 1, 2, 3, 4},
		},
		{
			NewRange(1),
			[]int{0},
		},
		{
			NewRange(0),
			[]int{},
		},
		{
			NewRange(-5),
			[]int{},
		},
		{
			NewRange(-1),
			[]int{},
		},
		{
			NewRange(5).WithStart(1),
			[]int{1, 2, 3, 4},
		},
		{
			NewRange(1).WithStart(1),
			[]int{},
		},
		{
			NewRange(5 + 1),
			[]int{0, 1, 2, 3, 4, 5},
		},
		{
			NewRange(5 + 1).WithStart(1),
			[]int{1, 2, 3, 4, 5},
		},
		{
			NewRange(5).WithStep(2),
			[]int{0, 2, 4},
		},
		{
			NewRange(5).WithStart(1).WithStep(2),
			[]int{1, 3},
		},
		{
			NewRange(5 + 1).WithStart(1).WithStep(2),
			[]int{1, 3, 5},
		},
	}

	for _, tc := range testcases {
		r := Collect(tc.r.Iter())

		if diff := cmp.Diff(r, tc.s); diff != "" {
			t.Log("range:", r)
			t.Log("slice:", tc.s)
			t.Error(diff)
		}
	}
}
