package main

import (
	"fmt"
	"regexp"
	"strings"
)

var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func ToSnakeCase(str string) string {
	snake := matchAllCap.ReplaceAllString(str, "${1}_${2}")
	fmt.Println(snake)
	return strings.ToLower(snake)
}

func main() {
	fmt.Println(ToSnakeCase("JapanCanadaAustraliaHelloWorld"))
}
