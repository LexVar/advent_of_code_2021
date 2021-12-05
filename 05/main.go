package main

import (
	"fmt"
)

const BOARD_SIZE = 1000

func read_lines () ([][]int) {
	var lines [][]int

	for n := 1; n > 0; {
		l := make([]int, 4)
		n, _ = fmt.Scanf("%d,%d -> %d,%d\n", &l[0], &l[1], &l[2], &l[3]);
		if n > 0 {
			lines = append(lines, l)
		}
	}
	return lines
}

func draw_board (lines [][]int) ([][]int, int) {
	count := 0
	// Initialize 9x9 board
	board := make([][]int, BOARD_SIZE)
	for i := 0; i < BOARD_SIZE; i++ {
		board[i] = make([]int, BOARD_SIZE)
		for j := 0; j < BOARD_SIZE; j++ {
			board[i][j] = 0
		}
	}

	for _, line := range lines {
		// If Xs are equal, or Ys are equal
		if line[0] == line[2] {
			r := python_range(line[1],line[3])
			for _, y := range r {
				board[line[0]][y]++
				// Count on the spot
				if board[line[0]][y] == 2 {
					count++
				}
			}
		} else if line[1] == line[3] {
			r := python_range(line[2],line[0])
			for _, x := range r {
				board[x][line[1]]++
				// Count on the spot
				if board[x][line[1]] == 2 {
					count++
				}
			}
		} else {
			xr := python_range(line[0],line[2])
			yr := python_range(line[1],line[3])
			for i := 0; i < len(xr); i++ {
				board[xr[i]][yr[i]]++
				// Count on the spot
				if board[xr[i]][yr[i]] == 2 {
					count++
				}
			}
		}
	}

	return board, count
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

func main () {
	lines := read_lines ()

	_, count := draw_board(lines)

	fmt.Println(count)
}
