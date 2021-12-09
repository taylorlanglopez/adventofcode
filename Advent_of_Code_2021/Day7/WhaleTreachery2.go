package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

func avg(s []int) int {
	avg := 0
	for _, v := range s {
		avg += v
	}
	return avg / len(s)
}

func WhaleTreachery2(fileName string) {
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

	targetInt := avg(numberList)

	var fuelCost int64 = 0
	for _, v := range numberList {
		dtt := int64(math.Abs(float64(targetInt - v)))
		fuelCost += (dtt * (dtt + 1) / 2)
	}

	fmt.Println(fuelCost)
}
