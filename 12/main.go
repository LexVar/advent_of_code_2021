package main

import (
	"fmt"
	"strings"
)

func read_lines () (map[string][]string) {
	var line string

	g := make(map[string][]string)

	for n := 1; n > 0; {
		n, _ = fmt.Scanf("%s\n", &line);
		if n > 0 {
			s := strings.Split(line, "-")
			g[s[0]] = append(g[s[0]], s[1])
			g[s[1]] = append(g[s[1]], s[0])
		}
	}
	return g
}

func is_uppercase(name string) bool {
	is_upper := true

	uppers := "ABCDEFGHIJKLMNOPQRSTUVWXZ"

	for _, c := range name {
		if ! strings.Contains(uppers, string(c)) {
			is_upper = false
		}
	}
	return is_upper
}

func contains(s []string, e string) bool {
    for _, n := range s {
        if n == e {
            return true
        }
    }
    return false
}

func count_item(s []string, e string) int {
	count := 0
    for _, n := range s {
        if n == e {
			count++
        }
    }
    return count
}

func part1(g map[string][]string, visited []string, node string) int {
	if node == "end" {
		fmt.Println(visited)
		return 1
	}
	count := 0

	for _, n := range g[node] {
		if is_uppercase(n) {
			count += part1(g, visited, n)
		} else {
			if ! contains(visited, n) {
				count += part1(g, append(visited, n), n)
			}
		}
	}

	return count
}

func part2(g map[string][]string, visited map[string]int, node string) int {
	if node == "end" {
		return 1
	}
	count := 0

	for _, n := range g[node] {
		// make copy of map
		vis := make(map[string]int)
		for k, v := range visited { vis[k] = v }

		if is_uppercase(n) {
			count += part2(g, vis, n)
		} else {
			c, node_visited := vis[n]
			if ! node_visited || c < 1{
				vis[n] = 1
				count += part2(g, vis, n)
			} else if n != "start" && c == 1 && vis["small"] <= 1 {
				vis[n]++
				vis["small"] = vis[n]
				count += part2(g, vis, n)
			}
		}
	}

	return count
}

func main () {
	g := read_lines()

	visited := make(map[string]int)
	visited["start"] = 1
	visited["small"] = 1

	// count := part1(g, []string{"start"}, "start")
	// fmt.Println(count)
	count := part2(g, visited, "start")

	fmt.Println(count)
}
