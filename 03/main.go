package main

import (
	"bufio"
	"fmt"
	"os"
	"math"
)

func part1() {
	file, err := os.Open("input1")
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
	fmt.Println(array)

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

func read_lines() ([]string){
	var lines []string
	var line string

	for n := 1; n > 0; {
		n, _ = fmt.Scanf("%s\n", &line);

		if n > 0 {
			lines = append(lines, line)
		}
	}

	return lines
}

func count_bits(lines []string, index int, indexes []int) (int, int){
	ones, zeros := 0, 0

	for _, i := range indexes {
		if lines[i][index] == '1' {
			ones++
		} else {
			zeros++
		}
	}
	return ones, zeros
}

func get_lines(lines []string, indexes []int, index int, choice byte) ([]int){
	res := []int{}

	for _, i := range indexes {
		if lines[i][index] == choice {
			res = append(res, i)
		}
	}
	return res
}

func int_to_bin(s string) (int) {
	size := len(s)
	n := 0
	for i:= size-1; i >= 0; i-- {
		if s[i] == '1' {
			n += int(math.Pow(2,float64(size-i-1)))
		}
	}
	return n
}

func part2(lines []string) (int, int){
	indexes_ox := []int{}
	indexes_co := []int{}

	// Create indexes list 0,1,2,3,..
	for i := 0; i < len(lines); i++ {
		indexes_ox = append(indexes_ox, i)
		indexes_co = append(indexes_co, i)
	}

	for i := 0; i < len(lines[0]); i++ {
		ones, zeros := count_bits(lines, i, indexes_ox)

		if ones >= zeros {
			indexes_ox = get_lines(lines, indexes_ox, i, '1')
		} else {
			indexes_ox = get_lines(lines, indexes_ox, i, '0')
		}

		if len(indexes_ox) == 1 {
			fmt.Println(lines[indexes_ox[0]])
			break
		}
	}
	for i := 0; i < len(lines[0]); i++ {
		ones, zeros := count_bits(lines, i, indexes_co)
		fmt.Println(ones, zeros)

		if ones >= zeros {
			indexes_co = get_lines(lines, indexes_co, i, '0')
		} else {
			indexes_co = get_lines(lines, indexes_co, i, '1')
		}

		if len(indexes_co) == 1 {
			fmt.Println(lines[indexes_co[0]])
			break
		}
	}
	return int_to_bin(lines[indexes_ox[0]]), int_to_bin(lines[indexes_co[0]])
}

func main () {
	// part1()
	lines := read_lines()
	fmt.Println(lines)

	oxygen, co2 := part2(lines)
	fmt.Println(oxygen, co2)
	fmt.Println(oxygen * co2)
}
