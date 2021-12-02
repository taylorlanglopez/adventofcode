package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func SonarSweep(fileName string) {
	var prev = 0
	var count = 0

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Can't read file: ", os.Args[1])
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var plato = false

	for scanner.Scan() {
		var text = scanner.Text()
		var num, err = strconv.Atoi(text)

		if err != nil {
			fmt.Printf("Couldn't convert \" %v \" to an integer\n", text)
			continue
		}

		if !plato {
			prev = num
			plato = true
			continue
		}

		if prev < num {
			count++
		}

		prev = num
	}

	fmt.Println("Number of descents: ", count)
}
