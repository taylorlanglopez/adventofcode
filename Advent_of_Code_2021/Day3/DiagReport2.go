package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func remove(s []string, j int) []string {
	s[j] = s[len(s)-1]
	return s[:len(s)-1]
}

func DiagReport2(fileName string) {
	strArr, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	master := strings.Replace(string(strArr), "\r\n", "\n", -1)
	list := strings.Split(master, "\n")
	list = list[:len(list)-1]

	bitWidth := len(list[0])
	columnTracker := make([]Pair, bitWidth)
	fmt.Println("Bitwidth ->", bitWidth)

	for _, v := range list {
		if v == "" {
			continue
		}
		for i := 0; i < bitWidth; i++ {
			if v[i] == '1' {
				columnTracker[i].one++
			} else {
				columnTracker[i].zero++
			}
		}
	}

	gamma := getGamma(columnTracker)
	epsilon := getEpsilon(columnTracker)
	fmt.Println("Gamma ->", gamma)
	fmt.Println("Epsilon ->", epsilon)

	oxy := parseList(gamma, list, true)
	co2 := parseList(epsilon, list, false)

	fmt.Println("Oxy ->", oxy)
	fmt.Println("Co2 ->", co2)
	fmt.Println("Result ->", convertToNum(oxy)*convertToNum(co2))
}

func parseList(s string, list []string, ge bool) string {

	for i := 0; i < len(s) && len(list) > 1; i++ { //12 times
		newList := make([]string, 0)
		for _, v := range list { // 12 * 1000 times
			digit := v[i]
			if s[i] == digit {
				newList = append(newList, v)
			}
		}
		list = newList

		if ge {
			s = recalcGamma(list)
		} else {
			s = inverseBinString(recalcGamma(list))
		}
	}
	return list[0]
}

func recalcGamma(list []string) string {
	bitWidth := len(list[0])
	columnTracker := make([]Pair, bitWidth)

	for _, v := range list {
		for i := 0; i < bitWidth; i++ {
			if v[i] == '1' {
				columnTracker[i].one++
			} else {
				columnTracker[i].zero++
			}
		}
	}

	return getGamma(columnTracker)
}
