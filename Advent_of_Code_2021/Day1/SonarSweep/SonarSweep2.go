package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func SonarSweep2(fileName string) {
	var count = 0
	var master = []int{}

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Can't read file: ", os.Args[1])
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var text = scanner.Text()
		var num, err = strconv.Atoi(text)

		if err != nil {
			fmt.Printf("Couldn't convert \" %v \" to an integer\n", text)
			continue
		}

		master = append(master, num)
	}

	var windowA = master[0] + master[1] + master[2]

	for i := 0; i < len(master); i++ {
		var windowB = sum(master[i : i+3])

		if windowA < windowB {
			count++
		}

		windowA = windowB
	}

	fmt.Println("Number of window increases: ", count)
}
