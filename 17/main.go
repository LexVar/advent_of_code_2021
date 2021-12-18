package main

import (
	"fmt"
	"math"
)

func read_lines() (x [2]int, y [2]int) {

	fmt.Scanf("target area: x=%d..%d, y=%d..%d\n", &x[0], &x[1], &y[0], &y[1]);

	return x, y
}

// Emulates python's range function
// Create list of consecutive numbers between two integers
func python_range(n1 int, n2 int) ([]int) {
	hi, lo := 0, 0
	if (n1 < n2) {
		lo = n1
		hi = n2
	} else {
		hi = n1
		lo = n2
	}
	s := make([]int, hi-lo+1)

	if (n1 < n2) {
		for i := range s {
			s[i] = i + lo
		}
	} else {
		for i := range s {
			s[i] = hi - i
		}
	}

	return s
}

// Is the current point inside the area?
func in_area(x [2]int, y [2]int, cur [2]int) bool {
	return cur[0] >= x[0] && cur[0] <= x[1] && cur[1] >= y[0] && cur[1] <= y[1]
}

func overshot(x [2]int, y [2]int, cur [2]int) bool {
	// Calculate if the current point has overshot the area
	return cur[1] < y[0] || math.Abs(float64(cur[0])) > math.Abs(float64(x[1]))
}

// Calculates the maximum height achieved of a point inside the area
func maximum_height(x [2]int, y [2]int, cur [2]int) (int, bool) {

	height := 0
	valid := false
	
	// Calculate max height achieved n(n+1)/2
	height = cur[1]*(cur[1]+1)/2

	// find if the point is valid
	// while it does not overshot the area
	// optimization
	for i, j := cur[0]-1, cur[1]-1; ! overshot(x, y, cur); j-- {

		// if inside the area, it is a valid point
		if in_area(x, y, cur)  {
			valid = true
			break
		}

		// Add velocity
		cur[0] += i
		cur[1] += j

		// Apply drag to x
		if i > 0 { 
			i--
		} else if i < 0 {
			i++
		}
	}

	return height, valid
}

func part1(x [2]int, y [2]int) (max int) {

	bound := int(math.Abs(float64(y[0]))-1)
	max = bound*(bound+1)/2

	return
}

func part2(x [2]int, y [2]int) (count int) {

	// give list of possible points
	// x: between start(0) and outerbound of area
	// y: between lowerbound of area and maximum y velocity
	xn := python_range(0, x[1])
	yn := python_range(-y[0], y[0])

	for _, i := range xn {
		for _, j := range yn {
			// find valid points and count them
			if _, valid := maximum_height(x, y, [2]int{i,j}); valid  {
				count++
			}
		}
	}
	return
}

func main () {
	x, y := read_lines()
	fmt.Println(x, y)

	max := part1(x, y)
	fmt.Println(max)

	count := part2(x, y)
	fmt.Println(count)
}
