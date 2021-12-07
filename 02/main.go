package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main () {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var pos = [3]int{0, 0, 0}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		// Read integer
		line := scanner.Text()
        if err != nil {
            return
        }
		parameters := strings.SplitAfter(line , " ")
		command := parameters[0]
		count, err := strconv.Atoi(parameters[1])
        if err != nil {
            return
        }

		switch command {
			case "forward ":
				pos[0] += count
				pos[1] += pos[2]*count
			case "up ":
				pos[2] -= count
			case "down ":
				pos[2] += count
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println(pos[0] * pos[1])
}
