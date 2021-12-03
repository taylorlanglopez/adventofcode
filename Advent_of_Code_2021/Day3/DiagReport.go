package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Pair struct {
	zero int
	one  int
}

func DiagReport(fileName string) {
	strArr, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	master := strings.Replace(string(strArr), "\r\n", "\n", -1)
	list := strings.Split(master, "\n")
	list = list[:len(list)-1]

	bitWidth := len(list[0])
	columnTracker := make([]Pair, bitWidth)
	fmt.Println("Bitwidth -> ", bitWidth)
	fmt.Println("List[0] ->", list[0])

	for _, v := range list {
		for i := 0; i < bitWidth; i++ {
			if v[i] == '1' {
				columnTracker[i].one++
			} else {
				columnTracker[i].zero++
			}
		}
	}

	gamma := getGamma(columnTracker)
	epsilon := inverseBinString(gamma)
	fmt.Println(columnTracker)

	// function here that converts binary strings to decimal numbers
	fmt.Println(gamma, epsilon)
	fmt.Println(convertToNum(gamma) * convertToNum(epsilon))
}

func convertToNum(s string) int64 {
	strlen := len(s)
	var retVal int64 = 0

	for i, v := range s {
		oneOrZero, err := strconv.ParseInt(string(v), 10, 64)

		if err != nil {
			panic(err)
		}

		retVal = retVal | (oneOrZero << (int64(strlen - i - 1)))
	}

	return retVal
}

func getGamma(colTrack []Pair) string {
	var retVal string = ""
	for _, v := range colTrack {
		if v.one >= v.zero {
			retVal += "1"
		} else {
			retVal += "0"
		}
	}
	return retVal
}

func getEpsilon(colTrack []Pair) string {
	var retVal string = ""
	for _, v := range colTrack {
		if v.one >= v.zero {
			retVal += "0"
		} else {
			retVal += "1"
		}
	}
	return retVal
}

func inverseBinString(s string) string {
	var retVal string = ""
	for _, v := range s {
		if v == '1' {
			retVal += "0"
		} else {
			retVal += "1"
		}
	}
	return retVal
}
