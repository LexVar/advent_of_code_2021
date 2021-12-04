package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func read_bingo (filename string) ([]int, [][][]int) {

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	// Get list of numbers
	numbers := strings.Split(scanner.Text(), ",")

	// Read empty line
	scanner.Scan()

	var boards [][][]string
	// boards := make([][][]string, 1, 200)
	board := make([][]string, 5)

	i := 0

	// Get boards
	for scanner.Scan() {
		// Read board line
		board[i] = strings.Fields(scanner.Text())
		i++
		if i == 5 {
			i = 0
			b := make([][]string, 5)
			copy(b, board)
			boards = append(boards, b)

			// Read empty line
			scanner.Scan()
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	nums := make([]int, len(numbers))

	for i, number := range numbers {
		nums[i], err = strconv.Atoi(number)
		if err != nil {
			fmt.Println(err)
		}
	}
	
	new_board := make([][][]int, len(boards))
	for i := 0; i < len(boards); i++ {
		new_board[i] = make([][]int, 5)
		for j := 0; j < 5; j++ {
			new_board[i][j] = make([]int, 5)
			for k := 0; k < 5; k++ {
				num, err := strconv.Atoi(boards[i][j][k])
				new_board[i][j][k] =  num
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	return nums, new_board
}

// Check if number is in bingo board
func check_bingo_board (number int, board [][]int) (int, int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] == number {
				return i,j
			}
		}
	}
	return -1,-1
}

// Check all bingo boards
func bingo_last (numbers []int, boards [][][]int) ([][]int, int) {
	finished := make([]int, len(boards))
	finished_num := 0
	for i := 0; i < len(boards); i++ {
		finished[i] = 0
	}

	bingo := make([][][]int, len(boards))
	// Initialize bingo slice with zeros
	for i := 0; i < len(bingo); i++ {
		bingo[i] = [][]int{{0, 0, 0, 0, 0},
						{0, 0, 0, 0, 0},
					}
	}

	for _, number := range numbers {
		for index, board := range boards {
			i, j := check_bingo_board(number, board)

			// If number is found in the card, increase line and collumn counters
			if i >= 0 && j >= 0 {
				bingo[index][0][i]++
				bingo[index][1][j]++
				// Check if any collumn or line has all 5 numbers
				if bingo[index][0][i] == 5 || bingo[index][1][j] == 5 {
					// Add finished board
					if finished[index] != 1 {
						finished[index] = 1
						finished_num++
					} 

					if finished_num == len(boards) {
						return board, number
					}
				}
			}
		}
	}
	return nil, -1
}

// Check all bingo boards
func bingo_winner (numbers []int, boards [][][]int) ([][]int, int) {
	bingo := make([][][]int, len(boards))
	// Initialize bingo slice with zeros
	for i := 0; i < len(bingo); i++ {
		bingo[i] = [][]int{{0, 0, 0, 0, 0},
						{0, 0, 0, 0, 0},
					}
	}

	for _, number := range numbers {
		for index, board := range boards {
			i, j := check_bingo_board(number, board)

			// If number is found in the card, increase line and collumn counters
			if i >= 0 && j >= 0 {
				bingo[index][0][i]++
				bingo[index][1][j]++
				// Check if any collumn or line has all 5 numbers
				if bingo[index][0][i] == 5 || bingo[index][1][j] == 5 {
					return board, number
				}
			}
		}
	}
	return nil, -1
}


func find_unmarked_sum (board [][]int, numbers []int, num int) (int) {
	sum := 0
	for _,number := range numbers {
		i, j := check_bingo_board(number, board)
		if i >= 0 && j >= 0 {
			board[i][j] = -1
		}
		if num == number {
			break
		}
	}
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if board[i][j] != -1 {
				sum += board[i][j]
			}
		}
	}
	return sum
}

func main () {
	numbers, boards := read_bingo ("input")

	board, it := bingo_last(numbers, boards)
	// board, it := bingo_first(numbers, boards)
	fmt.Println(board)
	sum := find_unmarked_sum(board, numbers, it)
	fmt.Println(sum , it)
	fmt.Println(sum * it)
}
