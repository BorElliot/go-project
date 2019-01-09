package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	project_url := "https://baidu.com?click=3"

	u, err := url.Parse(project_url)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.Scheme)

	fmt.Println(u.Host)

	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RawQuery)
}
