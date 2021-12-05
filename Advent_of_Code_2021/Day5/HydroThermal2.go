package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func HydroThermal2(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer TotalProgramTime()()
	defer file.Close()

	var lines []Line
	max_num := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			twoPairs := strings.Split(scanner.Text(), " -> ")
			lines = append(lines, ParseLine(twoPairs[0], twoPairs[1], &max_num))
		}
	}

	gridTrack := make([][]int, max_num+1, max_num+1)
	for i := range gridTrack {
		gridTrack[i] = make([]int, max_num+1, max_num+1)
	}

	makeLinesWithDiagonal(lines, &gridTrack)
	count := 0
	for _, v := range gridTrack {
		for _, e := range v {
			if e >= 2 {
				count++
			}
		}
	}
	fmt.Println(count)
}

func makeLinesWithDiagonal(lines []Line, grid *[][]int) {
	for _, v := range lines {
		// 1,1 -> 3,3 or 3,3 -> 1, 1
		if v.y1 == v.x1 && v.x2 == v.y2 {
			fmt.Println(v)
			x1 := v.x1
			x2 := v.x2
			y1 := v.y1
			y2 := v.y2
			swapIfBigger(&x1, &x2)
			swapIfBigger(&y1, &y2)
			for i := x1; i <= x2; i++ {
				(*grid)[i][i]++
			}
			// Inverse points 7,9 -> 9, 7
			// or 9, 7 -> 7, 9
		} else if v.x1 == v.y2 && v.x2 == v.y1 {
			max_x := max(v.x1, v.x2)
			max_y := max(v.y1, v.y2)
			for i, j := v.x1, v.y1; ; {
				(*grid)[i][j]++
				if i == v.x2 || j == v.y2 {
					break
				}
				if v.x1 == max_x {
					iterDown(&i, v.x2)
				} else if v.x2 == max_x {
					iterUp(&i, v.x2)
				}
				if v.y1 == max_y {
					iterDown(&j, v.y2)
				} else if v.y2 == max_y {
					iterUp(&j, v.y2)
				}
			}
			// Horizontal line 1, 5 -> 1, 9
		} else if v.x1 == v.x2 {
			y1 := v.y1
			y2 := v.y2
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for i := y1; i <= y2; i++ {
				(*grid)[v.x1][i]++
			}
			// Vertical Line 2, 3 -> 9, 3
		} else if v.y1 == v.y2 {
			x1 := v.x1
			x2 := v.x2
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for i := x1; i <= x2; i++ {
				(*grid)[i][v.y1]++
			}
			// Diagonal with no matching coordinates
			// 1, 4 -> 2, 5
		} else {
			max_x := max(v.x1, v.x2)
			max_y := max(v.y1, v.y2)
			for i, j := v.x1, v.y1; ; {
				(*grid)[i][j]++
				if i == v.x2 || j == v.y2 {
					break
				}
				if v.x1 == max_x {
					iterDown(&i, v.x2)
				} else if v.x2 == max_x {
					iterUp(&i, v.x2)
				}
				if v.y1 == max_y {
					iterDown(&j, v.y2)
				} else if v.y2 == max_y {
					iterUp(&j, v.y2)
				}
			}
		}
	}
}

func swapIfBigger(a, b *int) {
	if *a > *b {
		*a, *b = *b, *a
	}
}

func iterDown(i *int, dest int) {
	if *i >= dest {
		*i--
	}
}

func iterUp(i *int, dest int) {
	if *i <= dest {
		*i++
	}
}
