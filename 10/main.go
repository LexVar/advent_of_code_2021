package main

import (
	"fmt"
	"strings"
)

func read_lines () ([]string) {
	lines := []string{}
	var line string

	for n := 1; n > 0; {
		n, _ = fmt.Scanf("%s\n", &line);
		if n > 0 {
			lines = append(lines, line)
		}
	}
	return lines
}

func part1(lines []string) (int) {
	score := 0
	stack := []rune{}
	matches := map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
		'<': '>',
	}
	scores := map[rune]int{
		']': 57,
		')': 3,
		'}': 1197,
		'>': 25137,
	}
	opens := "[({<"
	closes := "])}>"

	for _, line := range lines {
		for _, c := range line {

			if strings.Contains(opens, string(c)) {
				stack = append(stack, c)
			} else if strings.Contains(closes, string(c)) {
				if matches[stack[len(stack)-1]] == c {
					stack = stack[:len(stack)-1]
				} else {
					score += scores[c]
					break
				}
			}
		}
	}

	return score
}

func insert_sorted (nums []int, n int) ([]int) {
	i := 0

	for ; i < len(nums) && nums[i] < n; i++ { }

	return append(nums[:i], append([]int{n}, nums[i:]...)...)
}

func part2(lines []string) (int) {
	stack := []rune{}
	total := []int{}
	match := map[rune]rune{
		'[': ']',
		'(': ')',
		'{': '}',
		'<': '>',
	}
	scores := map[rune]int{
		']': 2,
		')': 1,
		'}': 3,
		'>': 4,
	}
	opens := "[({<"
	closes := "])}>"

	for _, line := range lines {
		corrupted := 0
		stack = []rune{}
		for _, c := range line {

			if strings.Contains(opens, string(c)) {
				stack = append(stack, c)
			} else if strings.Contains(closes, string(c)) {
				if match[stack[len(stack)-1]] == c {
					stack = stack[:len(stack)-1]
				} else {
					// Ignore corrupted line
					corrupted = 1
					break
				}
			}
		}
		// Check if stack is empty, if not, complete line
		if corrupted == 0 && len(stack) > 0 {
			stack2 := []rune{}
			score := 0
			for i := len(stack)-1; i >= 0; i-- {

				if strings.Contains(closes, string(stack[i])) {
					stack2 = append(stack2, stack[i])
				} else {
					if len(stack2) > 0 && match[stack[i]] == stack2[len(stack2)-1] {
						stack2 = stack2[:len(stack2)-1]
					} else {
						score = score*5 + scores[match[stack[i]]]
					}
				}
			}
			total = insert_sorted(total, score)
		}
	}

	return total[len(total)/2]
}


func main () {
	lines := read_lines()

	// score := part1(lines)
	score := part2(lines)
	fmt.Println(score)
}
