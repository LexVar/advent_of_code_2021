package main

import (
	"fmt"
)

const BOARD_SIZE = 1000

func read_lines () ([]int) {
	var school []int
	var fish int

	for n := 1; n > 0; {
		n, _ = fmt.Scanf("%d ", &fish);
		if n > 0 {
			school = append(school, fish)
		}
	}
	return school
}

func count_fish (school []int, days int) ([]int, int) {
	// Newborns for each iteration
	var newborns []int
	for i := 0; i < days; i++ {
		for j := 0; j < len(school); j++ {
			if school[j] == 0 {
				school[j] = 6
				newborns = append(newborns, 8)
			} else {
				school[j]--
			}
		}
		// Append all newborns at the end of the day
		school = append(school, newborns...)
		newborns = make([]int, 0)

		// fmt.Println(school)
	}

	return school, len(school)
}

func main () {
	school := read_lines ()

	school, count := count_fish(school, 80)

	fmt.Println(count)
}
