package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func Lanternfish2(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer TotalProgramTime()()
	defer file.Close()

	daysToBirth, _ := ioutil.ReadAll(file)

	listOfFish := strings.Split(string(daysToBirth), ",")
	var numbers []int
	for _, v := range listOfFish {
		i, _ := strconv.Atoi(v)
		numbers = append(numbers, i)
	}

	histogram := make([]int64, 9, 9)

	for _, v := range numbers {
		histogram[v]++
	}

	for i := 0; i < 256; i++ {
		zeroes := ageHistogram(&histogram)
		histogram[6] += zeroes
		histogram[8] += zeroes
	}
	fmt.Println(sum(histogram))
}

func ageHistogram(histogram *[]int64) int64 {
	var zeroes int64 = -999
	var save int64 = -999
	for i := len(*histogram) - 1; i > 0; i-- {
		if zeroes != -999 {
			save = (*histogram)[i-1]
			(*histogram)[i-1] = zeroes
			zeroes = save
		} else {
			zeroes = (*histogram)[i-1]
			(*histogram)[i-1] = (*histogram)[i]
			(*histogram)[i] = 0
		}
	}
	return zeroes
}

func sum(histogram []int64) int64 {
	var sum int64
	for i := range histogram {
		sum += int64(histogram[i])
	}
	return sum
}
