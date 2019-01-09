package main

import (
	"fmt"
	"strings"
)

func ToCamelCase(str string) string {
	temp := strings.Split(str, "-")
	for i, r := range temp {
		if i > 0 {
			temp[i] = strings.Title(r)
		}
	}

	return strings.Join(temp, "")
}

func main() {
	str := "the-stealth-warrior"
	fmt.Println(ToCamelCase(str))
}
