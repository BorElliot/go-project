package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Height int    `json:"height"`
	Email  string `json:"email"`
}

func main() {
	peopleOne := &Person{
		Name:   "Jack",
		Age:    23,
		Height: 180,
		Email:  "jack@hotmail.com",
	}

	peopleTwo := &Person{
		Name:   "Telen",
		Age:    32,
		Height: 176,
		Email:  "telen@hotmail.com",
	}

	people := []*Person{}
	people = append(people, peopleOne, peopleTwo)

	result, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string(result))
	})

	http.ListenAndServe(":10060", nil)
}
