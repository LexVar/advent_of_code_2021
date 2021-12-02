package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
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
            return
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

	fmt.Println(counter)
}
