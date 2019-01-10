package main

import (
	"fmt"
	"regexp"
)

var re = regexp.MustCompile("(_|-)([a-zA-Z]+)")

func ToCamelCase(str string) string {
	camel := re.ReplaceAllString(str, "$2")

	return camel
}

func main() {
	str := "The_Stealth_Warrior"
	fmt.Println(ToCamelCase(str))
}
