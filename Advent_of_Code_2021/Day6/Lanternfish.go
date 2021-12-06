package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func Lanternfish(fileName string) {
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

	for days := 0; days < 80; days++ {
		for i := range numbers {
			numbers[i]--
			if numbers[i] == -1 {
				addNumber(&numbers)
				numbers[i] = 6
			}
		}
	}
	fmt.Println(len(numbers))
}

func addNumber(numbers *[]int) {
	*numbers = append(*numbers, 8)
}
