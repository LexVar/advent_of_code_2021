package main

import (
	"fmt"
)

func read_lines () ([][]int) {
	grid := [][]int{}
	var line string

	for n := 1; n > 0; {
		n, _ = fmt.Scanf("%s\n", &line);
		if n > 0 {
			l := make([]int, len(line))
			grid = append(grid, l)

			for i := 0; i < len(line); i++ {
				l[i] = int(line[i] - 48)
			}
		}
	}
	return grid
}

func is_low_point(grid [][]int, i int, j int) (bool) {
	low_point := true

	points := [4][2]int{{i,j-1}, {i-1,j}, {i+1,j}, {i,j+1}}

	for _, point := range points {
		if point[0] >= 0 && point[0] <= (len(grid)-1) && point[1] >= 0 && point[1] <= (len(grid[0])-1) {
			if grid[i][j] >= grid[point[0]][point[1]] {
				low_point = false
			}
		}
	}

	return low_point
}

func find_low_points(grid [][]int) ([][]int) {
	points := [][]int{}

	for i, row := range grid {
		for j := 0; j < len(row); j++ {
			if is_low_point(grid, i, j) {
				point := make([]int, 2)
				point[0] = i
				point[1] = j
				points = append(points, point)
			}
		}
	}
	return points
}

func calc_risk(grid [][]int, points [][]int) (int) {
	risk := 0
	for _, point := range points {
		risk += grid[point[0]][point[1]]+1
	}
	return risk
}

func insert_sorted (nums []int, n int) ([]int) {
	i := 0

	for ; i < len(nums) && nums[i] > n; i++ { }

	return append(nums[:i], append([]int{n}, nums[i:]...)...)
}

func calc_basins(grid [][]int, points [][]int) (int) {
	basins := []int{}

	for i := 0; i < len(points); i++ {
		s := bfs(grid, points[i][0], points[i][1])
		basins = insert_sorted(basins, s)
	}

	fmt.Println(basins)
	return basins[0] * basins[1] * basins[2]
}

func contains(s [][]int, e [] int) bool {
    for _, a := range s {
        if a[0] == e[0] && a[1] == e[1] {
            return true
        }
    }
    return false
}

func bfs(grid [][]int, i int, j int) int {
	// Queue of points to visit
	queue := [][]int{}
	queue = append(queue, []int{i,j})

	// Set of visited points
	reached := [][]int{}
	reached = append(reached, []int{i,j})

	for len(queue) > 0 {
		// Pop next element from queue
		cur := []int{}
		cur, queue = queue[0], queue[1:]
		x := cur[0]
		y := cur[1]

		// Calc neighbours
		neighbors := [][]int{{x,y-1}, {x-1,y}, {x+1,y}, {x,y+1}}
		for _, p := range neighbors {

			// If it's a valid neighbor, visit it
			if p[0] >= 0 && p[0] < len(grid) && p[1] >= 0 && p[1] < len(grid[0]) {

				// Check if it was already visited and is part of the basin (< 9)
				if ! contains(reached, p) && grid[p[0]][p[1]] < 9 {
					reached = append(reached, p)
					queue = append(queue, p)
				}
			}
		}
	}
	return len(reached)
}

func main () {
	grid := read_lines ()
	points := find_low_points(grid)

	basin_size := calc_basins(grid, points)
	fmt.Println(basin_size)

	// risk := calc_risk(grid, points)

}
