package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func DiagReport2(fileName string) {
	strArr, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	master := strings.Replace(string(strArr), "\r\n", "\n", -1)
	list := strings.Split(master, "\n")
	list = list[:len(list)-1]

	gamma := calcGamma(list)
	epsilon := inverseBinString(gamma)
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
			s = calcGamma(list)
		} else {
			s = inverseBinString(calcGamma(list))
		}
	}
	return list[0]
}
