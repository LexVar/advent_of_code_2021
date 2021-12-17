package main

import (
	"fmt"
	"math"
)

var translate = map[string][]byte {
	"0": []byte{0, 0, 0, 0},
	"1": []byte{0, 0, 0, 1},
	"2": []byte{0, 0, 1, 0},
	"3": []byte{0, 0, 1, 1},
	"4": []byte{0, 1, 0, 0},
	"5": []byte{0, 1, 0, 1},
	"6": []byte{0, 1, 1, 0},
	"7": []byte{0, 1, 1, 1},
	"8": []byte{1, 0, 0, 0},
	"9": []byte{1, 0, 0, 1},
	"A": []byte{1, 0, 1, 0},
	"B": []byte{1, 0, 1, 1},
	"C": []byte{1, 1, 0, 0},
	"D": []byte{1, 1, 0, 1},
	"E": []byte{1, 1, 1, 0},
	"F": []byte{1, 1, 1, 1},
}

var sum int

func read_lines() ([]byte) {
	var transmission string

	bits := []byte{}

	fmt.Scanf("%s\n", &transmission);

	// Literal Value - code 4
	for _, c := range transmission {
		bits = append(bits, translate[string(c)]...)
	}

	return bits
}

func bits_to_decimal(bits []byte) int {
	num := 0

	for i := 0; i < len(bits); i++ {
		num += int(math.Pow(2, float64(len(bits)-i-1))) * int(bits[i])
	}

	return num
}

// Parse literal number packet, after version and type
func parse_literal_number(bits []byte) ([]byte, int, int) {
	number := 0
	parsed := 0

	field, bits := bits[0:5], bits[5:]
	parsed += 5
	for field[0] != 0 {
		field, bits = bits[0:5], bits[5:]
		parsed += 5

		// Calc number here
	}

	return bits, parsed, number
}

func parse_operator(bits[]byte) ([]byte, int) {
	var size, n, parsed, npackets int

	// length ID
	bits, id := parse_parameter(bits, 1)

	// number of subpackets
	if id == 1 {
		// Get number of subpackets to parse
		bits, size = parse_parameter(bits, 11)
		parsed += 12

		// Parse 1st subpacket
		bits, n = start_parser(bits)
		parsed += n

		// while number of subpackets are not parsed
		for npackets = 1; npackets < size; npackets++ {
			bits, n = start_parser(bits)
			parsed += n
		}

	// length of subpackets
	} else {
		// Get length of subpackets
		bits, size = parse_parameter(bits, 15)
		parsed += 16

		bits, n = start_parser(bits)

		// while length of subpackets are not parsed
		for parsed += n; parsed < size; parsed += n {
			bits, n = start_parser(bits)
		}
	}

	return bits, parsed
}

// Parse and return parameter
func parse_parameter(bits []byte, n int) ([]byte, int) {
		return bits[n:], bits_to_decimal(bits[0:n])
}

// start a packet parsing with version + id
func start_parser(bits []byte) ([]byte, int) {
	var version, id, parsed, n int

	// parse version
	bits, version = parse_parameter(bits, 3)
	sum += version

	// parse packet Id
	bits, id = parse_parameter(bits, 3)
	parsed += 6

	if len(bits) <= 0 {
	} else if id == 4 {
		bits, n, _ = parse_literal_number(bits)
		parsed += n
	} else {
		bits, n = parse_operator(bits)
		parsed += n
	}

	return bits, parsed
}

func part1(bits []byte) {

	for len(bits) > 7 {
		bits, _ = start_parser(bits)
	}

	return
}

func main () {
	bits := read_lines()
	sum = 0

	part1(bits)
	fmt.Println(sum)
}
