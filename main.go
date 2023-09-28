package main

import (
	"fmt"
	"strings"
)

func vinay(input string) string {
	return strings.ToUpper(input)
}

func main() {
	lowercaseString := "vinay kapoor"
	uppercaseString := vinay(lowercaseString)
	fmt.Println(uppercaseString)
}
