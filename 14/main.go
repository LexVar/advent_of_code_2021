package main

import (
	"fmt"
)

func read_lines () (string, map[string]string) {
	var template, pair, out string
	rules := make(map[string]string)

	fmt.Scanf("%s\n\n", &template);

	for n := 1; n > 0; {
		n, _ = fmt.Scanf("%s -> %s\n", &pair, &out);
		if n > 0 {
			rules[pair] = out
		}
	}

	return template, rules
}

func part1(template string, rules map[string]string, iterations int) (res string) {
	for it := 0; it < iterations; it++ {
		fmt.Println(template)
		res = string(template[0])

		for i := 0; i < len(template)-1; i++ {
			res += rules[template[i:i+2]] + string(template[i+1])
		}
		template = res
	}
	return
}

func insert_sorted (nums []int, n int) ([]int) {
	i := 0

	for ; i < len(nums) && nums[i] > n; i++ { }

	return append(nums[:i], append([]int{n}, nums[i:]...)...)
}

func elements(template string) ([]int) {
	elements := make(map[string]int)

	for _, letter := range template {
		if _, exists := elements[string(letter)]; !exists {
			elements[string(letter)] = 0
		} else {
			elements[string(letter)]++
		}
	}

	var counts []int
	for _, value := range elements {
		counts = insert_sorted(counts, value)
	}
	return counts
}

func part2(template string, rules map[string]string, iterations int) (res map[string]int, elements map[string]int) {
	fmt.Println(template)
	res = make(map[string]int)
	elements = make(map[string]int)

	elements[string(template[0])]++
	for i := 0; i < len(template)-1; i++ {
		res[template[i:i+2]]++
		elements[string(template[i+1])]++
	}

	for it := 0; it < iterations; it++ {
		new_res := make(map[string]int)

		for key := range res {
			insert := rules[key]
			elements[insert] += res[key]
			new_res[string(key[0])+insert] += res[key]
			new_res[insert+string(key[1])] += res[key]
		}
		// Make copy for new map
		res = make(map[string]int)
		for k, v := range new_res {
			res[k] = v
		}
	}

	return res, elements
}

func count_elements(elements map[string]int) int {
	var tmp []int

	for _, v := range elements {
		tmp = insert_sorted(tmp, v)
	}

	return tmp[0] - tmp[len(tmp)-1]
}

func main () {
	template, rules := read_lines()

	res, elements := part2(template, rules, 40)

	// res := part1(template, rules, 40)
	// count := elements(res)
	
	fmt.Println(res)
	fmt.Println(count_elements(elements))
	// fmt.Println(count[0] - count[len(count)-1])
}
