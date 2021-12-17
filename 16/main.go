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
	var bit_number []byte
	parsed := 0

	field, bits := bits[0:5], bits[5:]
	bit_number = append(bit_number, field[1:]...)
	parsed += 5

	for field[0] != 0 {
		field, bits = bits[0:5], bits[5:]
		parsed += 5
		bit_number = append(bit_number, field[1:]...)
	}
	return bits, parsed, bits_to_decimal(bit_number)
}

func calc_operation(id int, args []int) (int) {
	res := 0

	switch id {
	case 0:
		for _, n := range args { res += n }
	case 1:
		res = args[0]
		for i := 1; i < len(args); i++ { res = res * args[i] }
	case 2:
		res = math.MaxInt32
		for _, n := range args {
			if n < res {
				res = n
			}
		}
	case 3:
		for _, n := range args {
			if n > res {
				res = n
			}
		}
	case 5:
		if args[0] > args[1] {
			res = 1
		} else {
			res = 0
		}
	case 6:
		if args[0] < args[1] {
			res = 1
		} else {
			res = 0
		}
	case 7:
		if args[1] == args[0] {
			res = 1
		} else {
			res = 0
		}
	}
	return res
}

func parse_operator(bits[]byte) ([]byte, int, []int) {
	var size, n, parsed, npackets, res int
	var args []int

	// length ID
	bits, id := parse_parameter(bits, 1)

	// number of subpackets
	if id == 1 {
		// Get number of subpackets to parse
		bits, size = parse_parameter(bits, 11)
		parsed += 12

		// Parse 1st subpacket
		bits, n, res = start_parser(bits)
		parsed += n
		args = append(args, res)

		// while number of subpackets are not parsed
		for npackets = 1; npackets < size; npackets++ {
			bits, n, res = start_parser(bits)
			args = append(args, res)
			parsed += n
		}

	// length of subpackets
	} else {
		// Get length of subpackets
		bits, size = parse_parameter(bits, 15)
		parsed += 16

		bits, n, res = start_parser(bits)
		args = append(args, res)

		packetlen := n
		// while length of subpackets are not parsed
		for parsed += n; packetlen < size; parsed += n {
			bits, n, res = start_parser(bits)
			packetlen += n
			args = append(args, res)
		}
	}

	return bits, parsed, args
}

// Parse and return parameter
func parse_parameter(bits []byte, n int) ([]byte, int) {
		return bits[n:], bits_to_decimal(bits[0:n])
}

// start a packet parsing with version + id
func start_parser(bits []byte) ([]byte, int, int) {
	var version, id, parsed, n, res int
	var args []int

	// parse version
	bits, version = parse_parameter(bits, 3)
	sum += version

	// parse packet Id
	bits, id = parse_parameter(bits, 3)
	parsed += 6

	if id == 4 {
		// Parse a literal number packet
		bits, n, res = parse_literal_number(bits)
		parsed += n
	} else {
		// Parse an operator packet
		bits, n, args = parse_operator(bits)
		parsed += n

		// Compute operation with arguments
		res = calc_operation(id, args)
	}

	return bits, parsed, res
}

func main () {
	bits := read_lines()
	sum = 0
	res := 0

	_, _, res = start_parser(bits)

	fmt.Println(res)
	// fmt.Println(sum)
}
