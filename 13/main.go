package main

import (
	"fmt"
)

type fold struct {
	value int
	coor rune
}

func read_lines () ([][]int, []fold) {
	var n1, n2 int

	maxx := 0
	maxy := 0

	tmp := [][]int{}

	for n := 1; n > 0; {
		n, _ = fmt.Scanf("%d,%d\n", &n1, &n2);
		if n > 0 {
			if n1 > maxx {
				maxx = n1
			}
			if n2 > maxy {
				maxy = n2
			}
			tmp = append(tmp, []int{n1,n2})
		}
	}

	paper := make([][]int, maxy+1)
	for i := 0; i <= maxy; i++ {
		paper[i] = make([]int, maxx+1)
	}
	for _, dot := range tmp {
		paper[dot[1]][dot[0]] = 1
	}

	var folds []fold
	for n := 1; n > 0; {
		var f fold
		n, _ = fmt.Scanf("fold along %c=%d\n", &f.coor, &f.value);
		if n > 0 {
			folds = append(folds, f)
		}
	}


	return paper, folds
}

func part1(paper [][]int, folds []fold) ([][]int, int) {
	count := 0

	for _, f := range folds {
		if f.coor == 121 {
			for y := f.value+1; y < len(paper); y++ {
				for x := 0; x < len(paper[0]); x++ {
					if paper[y][x] == 1 {
						paper[len(paper)-y-1][x] = 1
					}
				}
			}
			paper = paper[:f.value]
		} else {
			for y := 0; y < len(paper); y++ {
				for x := f.value+1; x < len(paper[0]); x++ {
					if paper[y][x] == 1 {
						paper[y][len(paper[0])-x-1] = 1
					}
				}
			}
			for y := 0; y < len(paper); y++ {
				paper[y] = paper[y][:f.value]
			}
		}
		// printing
		for _, item := range paper {
			fmt.Println(item)
		}
	}
	for i := 0; i < len(paper); i++ {
		for j := 0; j < len(paper[0]); j++ {
			count += paper[i][j]
		}
	}

	// Count dots
	return paper, count
}

func main () {
	paper, folds := read_lines()

	for _, item := range paper {
		fmt.Println(item)
	}

	_, count := part1(paper, folds)

	fmt.Println(count)
}
