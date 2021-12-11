package main

import (
	"fmt"
	// "time"
)

func read_lines () ([][]int) {
	var line string
	grid := [][]int{}

	for n := 1; n > 0; {
		n, _ = fmt.Scanf("%s\n", &line);
		if n > 0 {
			l := make([]int, len(line))

			for i := 0; i < len(line); i++ {
				l[i] = int(line[i] - 48)
			}
			grid = append(grid, l)
		}
	}
	return grid
}

func flash(grid [][]int, i int, j int) ([][]int) {
	grid[i][j] = 0

	points := [8][2]int{{i,j-1}, {i-1,j}, {i+1,j}, {i,j+1}, {i-1,j-1}, {i-1,j+1}, {i+1,j+1}, {i+1,j-1}}

	for _, p := range points {
		if p[0] >= 0 && p[0] < 10 && p[1] >= 0 && p[1] < 10 {
			grid[p[0]][p[1]]++
		}
	}
	return grid
}

func update_grid(grid [][]int)([][]int, int, [][]int) {
	counter := 0
	flashed := [][]int{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] > 9 {
				grid = flash(grid, i, j)
				flashed = append(flashed, []int{i,j})
				counter++
			}
		}
	}

	return grid, counter, flashed
}

func part1(grid [][]int, steps int) (int) {
	counter := 0

	for s := 0; s < steps; s++ {
		flashed := [][]int{}
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid); j++ {
				grid[i][j]++
				if grid[i][j] > 9 {
					grid = flash(grid, i, j)
					flashed = append(flashed, []int{i,j})
					counter++
				}
			}
		}
		g, c, f := update_grid(grid)
		for c > 0 { 
			grid = g
			flashed = append(flashed, f...)
			// fmt.Println(c)
			counter += c
			// time.Sleep(time.Second)
			g, c, f = update_grid(grid)
		}
		for _, p := range flashed {
			grid[p[0]][p[1]] = 0

		}
		all_flashed := true
		for i := 0; i < len(grid) && all_flashed; i++ {
			for j := 0; j < len(grid) && all_flashed; j++ {
				if grid[i][j] != 0 {
					all_flashed = false
					break
				}
			}
		}
		if all_flashed {
			fmt.Println("All flashed at step ", s)
			return s+1
		}

		fmt.Println("Step ", s, " done")
	}

	return -1
}

func main () {
	grid := read_lines()

	counter := part1(grid, 1000200)

	fmt.Println(counter)
}
