package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Line struct {
	x1 int
	y1 int
	x2 int
	y2 int
}

func ParseLine(s string, e string, max_num *int) Line {
	startCoords := strings.Split(s, ",")
	endCoords := strings.Split(e, ",")
	var p Line
	p.x1, _ = strconv.Atoi(startCoords[0])
	p.y1, _ = strconv.Atoi(startCoords[1])
	p.x2, _ = strconv.Atoi(endCoords[0])
	p.y2, _ = strconv.Atoi(endCoords[1])
	findMax(p.x1, p.y1, p.x2, p.y2, max_num)
	return p
}

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func findMax(a, b, c, d int, max_num *int) {
	x := []int{a, b, c, d}
	for i := range x {
		for j := range x {
			if max(x[i], x[j]) > *max_num {
				*max_num = max(x[i], x[j])
			}
		}
	}
}

func HydroThermal(fileName string) {
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

	makeLines(lines, &gridTrack)
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

func makeLines(lines []Line, grid *[][]int) {
	for _, v := range lines {
		if v.x1 == v.x2 {
			y1 := v.y1
			y2 := v.y2
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for i := y1; i <= y2; i++ {
				(*grid)[v.x1][i]++
			}
		}

		if v.y1 == v.y2 {
			x1 := v.x1
			x2 := v.x2
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for i := x1; i <= x2; i++ {
				(*grid)[i][v.y1]++
			}
		}
	}
}
