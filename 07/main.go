package main

import (
	"fmt"
	"math"
)

const MAX_POSITION = 2000

func read_lines () ([]int, int) {
	positions := make([]int, MAX_POSITION)
	var pos int

	for i := 0; i < MAX_POSITION; i++ {
		positions[i] = 0
	}

	max := 0
	for n := 1; n > 0; {
		n, _ = fmt.Scanf("%d,", &pos);
		if n > 0 {
			positions[pos]++

			if positions[pos] > positions[max] {
				max = pos
			}
		}
	}
	return positions, max
}

func sum_range (n int) (int) {
	res := 0

	for i := 1; i <= n; i++ {
		res += i
	}

	return res
}

func count_fuel (positions []int, max int) (int) {
	fuel := math.MaxInt32

	for i := 0; i < len(positions); i++ {
		f := 0
		for j := 0; j < len(positions); j++ {
			n := int(math.Abs(float64(j-i)))
			sum := (n * (n+1)) / 2
			f += positions[j]*sum
			// f += positions[j]*sum_range(int(math.Abs(float64(j-i))))
		}
		if f < fuel {
			fuel = f
		}
	}

	return fuel
}

func main () {
	positions, max := read_lines ()

	fuel := count_fuel(positions, max)

	fmt.Println(fuel)
}
