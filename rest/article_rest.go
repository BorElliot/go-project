package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Article - Our struct for all articles
type Article struct {
	Id      int    `json:"Id"`
	Title   string `json:"Title"`
	Content string `json:"content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homepage")
}

// get articles list
func list(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	result, err := db.Query("SELECT id, title, content FROM articles")
	if err != nil {
		panic(err.Error())
	}

	article := Article{}
	articles := []Article{}

	for result.Next() {
		var id int
		var title, content string
		err = result.Scan(&id, &title, &content)
		if err != nil {
			panic(err.Error())
		}

		article.Id = id
		article.Title = title
		article.Content = content

		articles = append(articles, article)
	}

	json.NewEncoder(w).Encode(articles)
}

// get one artcle
func show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	request_id := vars["id"]

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	row, err := db.Query("SELECT id, title, content FROM articles WHERE id = ?", request_id)
	if err != nil {
		panic(err.Error())
	}

	article := Article{}
	for row.Next() {
		var id int
		var title, content string
		err = row.Scan(&id, &title, &content)
		if err != nil {
			panic(err.Error())
		}

		article.Id = id
		article.Title = title
		article.Content = content
	}
	json.NewEncoder(w).Encode(article)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", list)
	myRouter.HandleFunc("/articles/{id}", show)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
