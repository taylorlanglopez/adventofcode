package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func checkError(e error) {
	if e != nil {
		fmt.Println("Error ->", e)
		log.Fatal(e)
	}
}

func TotalProgramTime() func() {
	start := time.Now()
	return func() {
		end := time.Now()
		fmt.Printf("Total Time: %d ms\n",
			end.Sub(start).Milliseconds())
	}
}

func main() {

	if len(os.Args) < 3 {
		fmt.Println("You're missing the input.txt argument and the function argument")
		return
	}

	fileName := os.Args[1]
	funcSelection, _ := strconv.Atoi(os.Args[2])

	if funcSelection == 1 {
		HydroThermal(fileName)
	}

	if funcSelection == 2 {
		HydroThermal2(fileName)
	}

}
