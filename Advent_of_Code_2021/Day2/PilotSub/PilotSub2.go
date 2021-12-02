package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func executeCommandWithAim(pos, depth, aim *int, pair KeyValue) {
	dir := pair.str
	dep := pair.num

	switch dir {
	case "forward":
		*pos += dep
		if *aim != 0 {
			*depth = *depth + (*aim)*(dep)
		}
	case "up":
		*aim -= dep
	case "down":
		*aim += dep
	default:
		return
	}
}

func PilotSub2(fileName string) {
	pos := 0
	depth := 0
	aim := 0

	strArr, err := ioutil.ReadFile(fileName)

	if err != nil {
		panic(err)
	}

	master := strings.Replace(string(strArr), "\n", "\r\n", -1)
	list := strings.Split(master, "\n")

	for _, v := range list {
		if v == "" {
			continue
		}
		executeCommandWithAim(&pos, &depth, &aim, addPair(v))
	}

	fmt.Println("Pos, Depth ->", pos, depth, pos*depth)

}
