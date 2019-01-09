// Reference: https://www.calhoun.io/creating-random-strings-in-go/
package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const charset = "0123456789"

var db *sql.DB

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func main() {
	for i := 0; i < 10; i++ {
		go generate()
	}

	var input string
	fmt.Scanln(&input)
}

func generate() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	for i := 0; i < 1000000000; i++ {
		random_str := String(12)
		_, err = db.Exec("INSERT INTO codes(code) VALUES(?)", random_str)
		if err != nil {
			continue
		}
		log.Println(random_str)
	}
}
