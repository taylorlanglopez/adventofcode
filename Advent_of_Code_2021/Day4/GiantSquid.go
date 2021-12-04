package main

//Solution 39902

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Tile struct {
	value  int
	marked bool
}

type Board struct {
	rows      int
	columns   int
	unmarkSum int
	data      [][]Tile
}

func (e Board) printMarked() {
	fmt.Println()
	fmt.Println("---BOARD---")
	for _, r := range e.data {
		fmt.Printf("[ ")
		for _, v := range r {
			fmt.Printf("%v ", v.marked)
		}
		fmt.Println("]")
	}
	fmt.Println("---BOARD---")
	fmt.Println()
}

func (e Board) CheckWin() bool {
	// Row has won
	for _, r := range e.data {
		isWin := true
		for _, v := range r {
			if v.marked == false {
				isWin = false
				break
			}
		}
		if isWin {
			return true
		}
	}

	// Column has won
	for i := range e.data {
		isWin := true
		for _, r := range e.data {
			if r[i].marked == false {
				isWin = false
				break
			}
		}
		if isWin {
			return true
		}
	}

	return false
}

func getBoard(idx int, list []Board) Board {
	return list[idx]
}

func (e *Board) findNumber(n int) bool {
	for _, r := range e.data {
		for i, v := range r {
			if v.value == n && r[i].marked == false {
				r[i].marked = true
				e.unmarkSum -= v.value
				return true
			}
		}
	}
	return false
}

func NewTile(val int, marked bool) Tile {
	p := new(Tile)
	p.value = val
	p.marked = marked
	return *p
}

func NewTileSet(s string, r, c int) (int, [][]Tile) {
	tiles := strings.Fields(s)

	sum := 0
	newTileSet := make([][]Tile, r, c)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			currVal := tiles[(r*i)+j]
			newValue, err := strconv.Atoi(currVal)
			checkError(err)
			sum += newValue
			newTileSet[i] = append(newTileSet[i], NewTile(newValue, false))
		}
	}

	return sum, newTileSet
}

func NewBoard(s string) *Board {
	p := new(Board)
	p.rows = 5
	p.columns = 5
	p.unmarkSum, p.data = NewTileSet(s, p.rows, p.columns)
	return p
}

func parseBoard(s string) Board {
	return *NewBoard(s)
}

func prettyPrintBoardFromSet(bNum int, list []Board) {
	b := list[bNum]
	fmt.Println()
	fmt.Println("---BOARD---")
	fmt.Println("Unmarked Sum ->", b.unmarkSum)
	sum := 0
	for _, r := range b.data {
		fmt.Printf("[ ")
		for _, v := range r {
			fmt.Printf("%v::%v ", v.value, v.marked)
			if v.marked == false {
				sum += v.value
			}
		}
		fmt.Println("]")
	}
	fmt.Println("Calced Unmarked Sum ->", sum)
	fmt.Println("---BOARD---")
}

func GiantSquid(fileName string) {
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
			boardSet = append(boardSet, parseBoard(tempString))
			tempString = ""
			continue
		}
		tempString += scanner.Text() + " "
	}

	for _, v := range moveset {
		move, err := strconv.Atoi(v)
		checkError(err)
		win, boardnum, winningMove := checkMove(move, &boardSet)
		if win {
			unmarkedSum := getBoard(boardnum, boardSet).unmarkSum
			// Do a function here that gives us the answer
			// and prints some stuff out
			fmt.Println("Board # ->", boardnum)
			fmt.Println("Last Move ->", winningMove)
			fmt.Println("Unmarked Sum ->", unmarkedSum)
			fmt.Println("Solution ->", unmarkedSum*winningMove)
			break
		}
	}

}

// Executive Logic Region
func checkMove(move int, boardSet *[]Board) (bool, int, int) {

	// Check every board!
	for i := range *boardSet {
		currBoard := &(*boardSet)[i]
		find := (*currBoard).findNumber(move)
		if find {
			won := currBoard.CheckWin()
			if won {
				return true, i, move
			}
		}
	}

	return false, 0, move
}
