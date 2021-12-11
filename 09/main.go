package main

import (
	"fmt"
)

func read_lines () ([][]int) {
	heightmap := [][]int{}
	var line string

	for n := 1; n > 0; {
		n, _ = fmt.Scanf("%s\n", &line);
		if n > 0 {
			l := make([]int, len(line))
			heightmap = append(heightmap, l)

			for i := 0; i < len(line); i++ {
				l[i] = int(line[i] - 48)
			}
		}
	}
	return heightmap
}

func is_low_point(heightmap [][]int, i int, j int) (bool) {
	low_point := true

	points := [4][2]int{{i,j-1}, {i-1,j}, {i+1,j}, {i,j+1}}

	for _, point := range points {
		if point[0] >= 0 && point[0] <= (len(heightmap)-1) && point[1] >= 0 && point[1] <= (len(heightmap[0])-1) {
			if heightmap[i][j] >= heightmap[point[0]][point[1]] {
				low_point = false
			}
		}
	}

	return low_point
}

func find_low_points(heightmap [][]int) ([][]int) {
	points := [][]int{}

	for i, row := range heightmap {
		for j := 0; j < len(row); j++ {
			if is_low_point(heightmap, i, j) {
				point := make([]int, 2)
				point[0] = i
				point[1] = j
				points = append(points, point)
			}
		}
	}
	return points
}

func calc_risk(heightmap [][]int, points [][]int) (int) {
	risk := 0
	for _, point := range points {
		risk += heightmap[point[0]][point[1]]+1
	}
	return risk
}

func calc_basins(heightmap [][]int, points [][]int) ([][]int) {
	basin_size := [][]int{}

	for i := 0; i < len(points); i++ {
		s := check_basin(heightmap, points[i][0], points[i][1], -1, -1)
		fmt.Println(s)
	}
	return basin_size
}

func check_basin(heightmap [][]int, i int, j int, origi int, origj int) (int) {
	count := 0
	basin := true
	
	if i < 0 || i >= len(heightmap) || j < 0 || j >= len(heightmap[0]) {
		return count
	}

	fmt.Println("potential", i,j)
	points := [4][2]int{{i,j-1}, {i-1,j}, {i+1,j}, {i,j+1}}

	for _, point := range points {
		if basin && point[0] >= 0 && point[0] < len(heightmap) && point[1] >= 0 && point[1] < len(heightmap[0]) {
			if point[0] == origi && point[1] == origj {
				continue
			}
			if heightmap[i][j] >= heightmap[point[0]][point[1]] {
				basin = false
				return 0
			}
		}
	}
	if basin {
		fmt.Println("basin", i,j)
		for _, point := range points {
			if point[0] == origi && point[1] == origj {
				continue
			} else {
				count += check_basin(heightmap, point[0], point[1], i, j)
			}
		}
	}
	return count+1
}

func main () {
	heightmap := read_lines ()
	points := find_low_points(heightmap)

	fmt.Println(points)
	basin_size := calc_basins(heightmap, points)
	fmt.Println(basin_size)
	// risk := calc_risk(heightmap, points)

}
