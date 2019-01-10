package main

import (
	"fmt"
	"regexp"
	"strings"
)

var re = regexp.MustCompile("(_|-)([a-zA-Z]+)")

func ToCamelCase(str string) string {
	str = strings.Title(str)
	fmt.Println(str)
	camel := re.ReplaceAllString(str, "$2")

	return camel
}

func main() {
	str := "The-Stealth-Warrior"
	fmt.Println(ToCamelCase(str))
}
