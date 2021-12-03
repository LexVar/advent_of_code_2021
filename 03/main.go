package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
)

func main () {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	s := scanner.Scan()
	line := scanner.Text()
	size := len(line)
	count := 0

	array := make([]int, size)

	for s {
		for i:= 0; i < size; i++ {
			if line[i] == '1' {
				array[i] += 1
			}
		}
		count++

		// Read line
		s = scanner.Scan()
		line = scanner.Text()
        if err != nil {
            return
        }
	}

	gama := 0
	epsilon := 0
	for i:= size-1; i >= 0; i-- {
		if array[i] > count/2 {
			gama += int(math.Pow(2,float64(size-i-1)))
		} else {
			epsilon += int(math.Pow(2,float64(size-i-1)))
		}
	}
	fmt.Println(gama * epsilon)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
