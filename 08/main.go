package main

import (
	"fmt"
)

func part1 () ([][]string, int) {
	digits := [][]string{}
	var dummy string
	count := 0

	for n := 1; n > 0; {
		d := make([]string, 4)
		n, _ = fmt.Scanf("%s %s %s %s %s %s %s %s %s %s | %s %s %s %s\n", &dummy, &dummy, &dummy, &dummy, &dummy, &dummy, &dummy, &dummy, &dummy, &dummy, &d[0], &d[1], &d[2], &d[3]);
		if n > 0 {
			digits = append(digits, d)
			for i := 0; i < 4; i++ {
				if len(d[i]) == 2 || len(d[i]) == 3 || len(d[i]) == 4 || len(d[i]) == 7 {
					count++
				}
			}
		}
	}
	fmt.Println(digits)
	return digits, count
}

func read_lines () ([][]string, [][]string) {
	digits := [][]string{}
	outputs := [][]string{}

	for n := 1; n > 0; {
		d := make([]string, 10)
		o := make([]string, 4)
		n, _ = fmt.Scanf("%s %s %s %s %s %s %s %s %s %s | %s %s %s %s\n", &d[0], &d[1], &d[2], &d[3], &d[4], &d[5], &d[6], &d[7], &d[8], &d[9], &o[0], &o[1], &o[2], &o[3]);
		if n > 0 {
			digits = append(digits, d)
			outputs = append(outputs, o)
		}
	}
	return outputs, digits
}

func letter_in_digit(s string, b byte) (bool) {
	found := false

	for i := 0; i < len(s); i++ {
		if s[i] == b {
			found = true
			break
		}
	}
	return found
}

func equal_segments(s1 string, s2 string) (bool) {
	equal := true

	if len(s1) != len(s2) {
		return false
	}
	for i := 0; i < len(s1) && equal; i++ {
		if found := letter_in_digit(s2, s1[i]); !found {
			equal = false

		}
	}
	return equal
}

func common_letters(s1 string, s2 string) (int) {
	count := 0

	for i := 0; i < len(s1); i++ {
		if found := letter_in_digit(s2, s1[i]); found {
			count++
		}
	}
	return count
}

func decode(digits []string) ([]string) {
	// numbers := make(map[string]int)
	numbers := make([]string, 10)

	for _, d := range digits {
		if len(d) == 2 {
			numbers[1] = d
		} else if len(d) == 3 {
			numbers[7] = d
		} else if len(d) == 4 {
			numbers[4] = d
		} else if len(d) == 7 {
			numbers[8] = d
		}
	}

	// Find common segments between 1 and 5,2,3 (all have 5 segments)
	// The number three has 2 common segments; 5,2 have 1 in common
	for _, d := range digits {
		if len(d) != 5 {
			continue
		}
		if c := common_letters(d, numbers[1]); c == 2 {
			numbers[3] = d
			break
		}
	}
	// 5: 5,2,3
		// Letras em comum:
		// 1: 1, 1, 2
		// Descrubo o 3
		// 4: 2, 3
		// Descubro 5 e 2

	// Find 5 and 2
	for _, d := range digits {
		if len(d) != 5 || d == numbers[3] {
			continue
		}
		c := common_letters(d, numbers[4])
		if c == 3 {
			numbers[5] = d
		} else if c == 2 {
			numbers[2] = d
		}
	}

	// Now compare numbers with 6 segments: 9,6,0
	// 6: 9,6,0
		// 1: 2,1,2
		// Descubro o 6
		// 4: 4, 3
		// Descubro 9 e 0
	for _, d := range digits {
		if len(d) != 6 {
			continue
		}
		if c := common_letters(d, numbers[1]); c == 1 {
			numbers[6] = d
			break
		}
	}
	for _, d := range digits {
		if len(d) != 6 || d == numbers[6] {
			continue
		}
		c := common_letters(d, numbers[4])
		if c == 4 {
			numbers[9] = d
		} else if c == 3 {
			numbers[0] = d
		}
	}
	return numbers
}

func part2(digits [][]string, outputs [][]string) (int) {
	res := 0

	for i, entry := range digits {
		numbers := decode(entry)
		sum := 0
		for _, out := range outputs[i] {
			for dig, n := range numbers {
				if equal_segments(n, out) {
					if sum == 0 {
						sum = dig
					} else {
						sum = sum * 10 + dig
					}
					break
				}
			}
		}
		res += sum
	}

	return res
}

func main () {
	// digits, n := part1 ()
	outputs, digits := read_lines ()

	numbers := part2 (digits, outputs)
	fmt.Println(numbers)
}
