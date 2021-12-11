package utils

import (
	"fmt"
	"log"
	"time"
)

func CheckError(e error) {
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
