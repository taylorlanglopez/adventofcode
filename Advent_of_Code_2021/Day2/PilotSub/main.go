package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("You're missing the input.txt argument and the function argument")
		return
	}

	fileName := os.Args[1]
	funcSelection, _ := strconv.Atoi(os.Args[2])

	if funcSelection == 1 {
		PilotSub(fileName)
	}

	if funcSelection == 2 {
		PilotSub2(fileName)
	}
}
