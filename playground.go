package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "Johd Doe"
	arr1 := strings.Split(name, " ")
	fNameArr := strings.Split(arr1[0], "")
	lNameArr := strings.Split(arr1[1], "")
	fmt.Sprintln(fNameArr[0], lNameArr[0])
}
