package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Page struct {
	Title    string
	Filename string
	Content  string
}

type Pages []Page

var rawJson = []byte(`[{"Title":"First Page","Filename":"page1.txt","Content":"This is the 1st page."},{"Title":"Second Page","Filename":"page2.txt","Content":"The 2nd Page is this."}]`)

func main() {
	// Decoding the JSON
	var pages Pages
	err := json.Unmarshal(rawJson, &pages)
	if err != nil {
		log.Fatal("Problem decoding JSON", err)
	}

	// Re-encode for demonstration purposes
	rejson, err := json.Marshal(pages)
	if err != nil {
		log.Fatal("Cannot encode to JSON", err)
	}
	fmt.Fprintf(os.Stdout, "%s", rejson)
}
