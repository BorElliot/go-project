/*
 * FROM: https://www.ardanlabs.com/blog/2013/09/iterating-over-slices-in-go.html
 * slice这种指针的方式可以保证使用的是同样的数据
 * 如果只是简单的传入struct实例(即不适用&),那么range的时候就会copy一份
 */
package main

import "fmt"

type Dog struct {
	Name string
	Age  int
}

func main() {
	sameValue()
	copyValue()
}

func sameValue() {
	jackie := &Dog{
		Name: "Jackie",
		Age:  19,
	}
	fmt.Printf("Jackie Addr: %p\n", jackie)

	sammy := &Dog{
		Name: "Sammy",
		Age:  10,
	}
	fmt.Printf("Sammy Addr: %p\n\n", sammy)

	dogs := []*Dog{jackie, sammy}

	for _, dog := range dogs {
		fmt.Printf("Name: %s Age: %d\n", dog.Name, dog.Age)
		fmt.Printf("Addr: %p\n\n", dog)
	}
}

func copyValue() {
	jackie := Dog{
		Name: "Jackie",
		Age:  19,
	}
	fmt.Printf("Jackie Addr: %p\n", &jackie)

	sammy := Dog{
		Name: "Sammy",
		Age:  10,
	}
	fmt.Printf("Sammy Addr: %p\n", &sammy)

	// 此处会导致copy
	dogs := []Dog{jackie, sammy}
	fmt.Println("")

	// range的时候也会导致copy
	for _, dog := range dogs {
		fmt.Printf("Name: %s Age: %d\n", dog.Name, dog.Age)
		fmt.Printf("Addr: %p\n", &dog)
		fmt.Println("")
	}
}
