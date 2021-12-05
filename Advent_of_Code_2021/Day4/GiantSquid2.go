package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func GiantSquid2(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer TotalProgramTime()()
	defer file.Close()

	scanner := bufio.NewScanner(file)
	//First line is our moveset
	scanner.Scan()
	moveset := strings.Split(scanner.Text(), ",")

	var boardSet []Board

	tempString := ""

	for scanner.Scan() {
		if scanner.Text() == "" && tempString != "" {
			boardSet = append(boardSet, parseBoard(tempString, false))
			tempString = ""
			continue
		}
		tempString += scanner.Text() + " "
	}

	var LB Board
	reverseAny(boardSet)
	for _, v := range moveset {
		move, err := strconv.Atoi(v)
		checkError(err)
		win, boardnum, winningMove := checkMove(move, &boardSet, true)
		if win {
			LB = boardSet[boardnum]
			// Do a function here that gives us the answer
			// and prints some stuff out
			fmt.Println("Board # ->", boardnum)
			fmt.Println("Last Move ->", winningMove)
		}
		for i := len(boardSet) - 1; i >= 0; i-- {
			if boardSet[i].solved {
				boardSet = append(boardSet[:i], boardSet[i+1:]...)
			}
		}
	}

	prettyPrintBoard(LB)
	fmt.Println(getUnmarkedSum(LB))
}

func getUnmarkedSum(t Board) int {
	sum := 0
	for _, r := range t.data {
		for _, v := range r {
			if v.marked == false {
				sum += v.value
			}
		}
	}
	return sum
}

func inverseBoard(boardNum int, list []Board) Board {
	retVal := list[boardNum]
	for _, r := range retVal.data {
		for i := range r {
			r[i].marked = !r[i].marked
		}
	}
	return retVal
}

func reverseAny(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
