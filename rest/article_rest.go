package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// Article - Our struct for all articles
type Article struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
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

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang?charset=utf8&loc=Asia%2FShanghai&parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var id int
	var title, content string
	var createdAt string
	article := Article{}
	err = db.QueryRow("select id, title, content, created_at from articles where id = ?", request_id).Scan(&id, &title, &content, &createdAt)
	if err != nil {
		if err == sql.ErrNoRows {
			errData := map[string]string{"errcode": "ER404", "errmsg": "not found"}
			json.NewEncoder(w).Encode(errData)
			return
		}
	}

	t1, _ := time.Parse(time.RFC3339, createdAt)
	createdAt = t1.Format("2006-01-02 15:04:05")

	article.Id = id
	article.Title = title
	article.Content = content
	article.CreatedAt = createdAt
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
