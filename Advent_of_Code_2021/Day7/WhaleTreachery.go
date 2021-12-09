package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func WhaleTreachery(fileName string) {
	defer TotalProgramTime()()
	file, err := ioutil.ReadFile(fileName)
	checkError(err)

	fileAsString := string(file)
	fileAsString = strings.Replace(fileAsString, "\r\n", "", -1)
	numberListString := strings.Split(fileAsString, ",")
	numberList := make([]int, 0)
	// We have all the horizontal positions in numberList
	for _, v := range numberListString {
		x, e := strconv.Atoi(v)
		checkError(e)
		numberList = append(numberList, x)
	}
	sort.Ints(numberList)

	targetInt := numberList[len(numberList)/2]

	var fuelCost int64 = 0
	for _, v := range numberList {
		fuelCost += int64(math.Abs(float64(targetInt - v)))
	}

	fmt.Println(fuelCost)
}
