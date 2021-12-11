package main

import (
	"fmt"
	"io/ioutil"

	"github.com/taylorlanglopez/adventofcode/Advent_of_Code_2021/utils"
)

func SevenSegment(fileName string) {
	defer utils.TotalProgramTime()()
	file, err := ioutil.ReadFile(fileName)

	utils.CheckError(err)

	fmt.Println(file)

	fmt.Println("Hello World!")
}
