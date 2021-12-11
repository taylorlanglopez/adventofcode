package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/taylorlanglopez/adventofcode/Advent_of_Code_2021/utils"
)

func SevenSegment(fileName string) {
	defer utils.TotalProgramTime()()
	file, err := os.Open(fileName)
	utils.CheckError(err)

	scanner := bufio.NewScanner(file)

	lastFourCodeStrings := make([]string, 0)

	// # => # of segments
	// 1 == 2
	// 4 == 4
	// 7 == 3
	// 8 == 7

	for scanner.Scan() {
		for _, i := range strings.Fields(strings.Split(scanner.Text(), "|")[1]) {
			lastFourCodeStrings = append(lastFourCodeStrings, i)
		}
	}

	countKeys := 0
	for _, v := range lastFourCodeStrings {
		if len(v) == 2 || len(v) == 4 || len(v) == 3 || len(v) == 7 {
			countKeys++
		}
	}

	fmt.Println(countKeys)
}
