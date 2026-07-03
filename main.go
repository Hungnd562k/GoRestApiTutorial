package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var Articles []Article

// Main home page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// Get all articles
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

// Get articles by Id
func returnArticleById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnArticleById")
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprint(w, "Parameter is invalid")
	}
	// fmt.Println("Finding articles with Id = %s", key)
	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

// Handle incoming request
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles)
	myRouter.HandleFunc("/articles/{id}", returnArticleById)
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		{Id: 1, Title: "Đây là title 1", Desc: "Mô tả của title 1", Content: "Còn đây là nội dung của Title 1"},
		{Id: 2, Title: "Đây là title 2", Desc: "Mô tả của title 2", Content: "Còn đây là nội dung của Title 2"},
	}
	handleRequests()
}
