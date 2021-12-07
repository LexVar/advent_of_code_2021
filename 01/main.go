package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func read_numbers() (int) {

	var num, sum int
	window := make([]int, 3)

	n, _ := fmt.Scanf("%d\n%d\n%d\n", &window[0], &window[1], &window[2]);

	counter, previous := 0, 0
	for _, m := range window {
		previous += m
	}
	for n > 0 {
		n, _ = fmt.Scanf("%d\n", &num);
		if n <= 0 {
			break
		}

		// Pop 1st element, add next element
		window = window[1:]
		window = append(window, num)
		// Sum numbers in window
		sum = window[0]+window[1]+window[2]

		fmt.Println(sum)

		if sum > previous {
			counter++
		}
		previous = sum
	}
	return counter
}

func part1 () (int) {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	index := 0
	counter := 0
	var previous int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Read integer
		x, err := strconv.Atoi(scanner.Text())
        if err != nil {
            return 0
        }
		if index > 0 && x > previous  {
			counter++
		}
		index++
		previous = x
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
	return counter
}

func main() {
	// counter := part1()
	counter := read_numbers()
	fmt.Println(counter)
}
