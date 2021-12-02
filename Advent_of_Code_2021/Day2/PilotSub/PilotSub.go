package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type KeyValue struct {
	str string
	num int
}

func addPair(command string) KeyValue {
	var temp KeyValue
	temp.str = parseString(command)
	temp.num = parseNum(command)
	return temp
}

func parseNum(command string) int {
	str := strings.Fields(command)
	v, err := strconv.Atoi(str[1])

	if err != nil {
		panic(err)
	}

	return v
}

func parseString(command string) string {
	str := strings.Fields(command)
	return str[0]
}

func executeCommand(pos, depth *int, pair KeyValue) {
	dir := pair.str
	dep := pair.num

	switch dir {
	case "forward":
		*pos += dep
	case "up":
		*depth -= dep
	case "down":
		*depth += dep
	default:
		return
	}
}

func PilotSub(fileName string) {
	pos := 0
	depth := 0

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
		executeCommand(&pos, &depth, addPair(v))
	}

	fmt.Println("Pos, Depth ->", pos, depth, pos*depth)

}
