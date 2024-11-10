package itertools

import (
	"fmt"
	"testing"
)

func Benchmark_map(b *testing.B) {
	size := 100_000

	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = i
	}

	for i := 10; i <= size; i *= 10 {
		b.Run(
			fmt.Sprintf("Map-%d", i),
			func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					seq := Map(
						Map(
							Filter(
								FromSlice(nums[:i]),
								func(n int) bool { return n%2 == 0 },
							),
							func(n int) int { return n * n },
						),
						func(n int) int { return n + 100 },
					)

					for range seq {
					}
				}
			},
		)
	}
}

func Benchmark_map_as_slice(b *testing.B) {
	var size int = 100_000

	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = i
	}

	for i := 10; i <= size; i *= 10 {
		b.Run(
			fmt.Sprintf("Map_as_slice-%d", i),
			func(b *testing.B) {
				for j := 0; j < b.N; j++ {
					rnums := make([]int, i)

					for k := 0; k < i; k++ {
						num := nums[k]
						if num%2 != 0 {
							continue
						}

						rnums[k] = num*num + 100
					}

					for range rnums {
					}
				}
			},
		)
	}
}
