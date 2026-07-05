package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	configpkg "github.com/Hungnd562k/GoRestApiTutorial/config_pkg"
	dbhelper "github.com/Hungnd562k/GoRestApiTutorial/dataaccess"
	"github.com/gorilla/mux"
)

type Article struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	Deleted bool   `json:"deleted"`
}

var Articles []Article

func health(w http.ResponseWriter, r *http.Request) {
	isDbConnect := dbhelper.CheckDbConnection()
	if !isDbConnect {
		http.Error(w, "Could not connect to Database", http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "Healthy")
}

// Main home page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// Get all articles
func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	var result []Article
	for _, article := range Articles {
		if article.Deleted == false {
			result = append(result, article)
		}
	}
	json.NewEncoder(w).Encode(result)
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
		if article.Id == key && article.Deleted != true {
			json.NewEncoder(w).Encode(article)
		}
	}
}

// Delete resource article by Id
func deleteArticleById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteArticleById")
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Fprint(w, "Parameter is invalid")
	}
	for index, article := range Articles {
		if article.Id == key && article.Deleted != true {
			Articles[index].Deleted = true
			fmt.Fprintln(w, "Delete article with Id: ", key, "SUCCESSED")
			return
		}
	}

	fmt.Fprintln(w, "Delete article with Id: ", key, "FAILED")
}

// Create new article
func createNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewArticle")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error while reading request's payload", http.StatusBadGateway)
		return
	}
	var jsonBody Article
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		http.Error(w, "Error while parsing request's payload", http.StatusBadRequest)
	}
	Articles = append(Articles, jsonBody)
	fmt.Fprintf(w, "Create new article successfully!")
}

func updateArticleById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateArticleById")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error while reading request's payload", http.StatusBadGateway)
		return
	}
	var jsonBody Article
	err = json.Unmarshal(body, &jsonBody)
	if err != nil {
		http.Error(w, "Error while parsing request's payload", http.StatusBadRequest)
	}
	id := jsonBody.Id
	for index := range Articles {
		if Articles[index].Id != id {
			continue
		}
		Articles[index].Content = jsonBody.Content
		Articles[index].Desc = jsonBody.Desc
		Articles[index].Title = jsonBody.Title
	}
	fmt.Fprintf(w, "Update article successfully!")
}

// Handle incoming request
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/health", health).Methods("GET")
	myRouter.HandleFunc("/articles", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles/{id}", returnArticleById).Methods("GET")
	myRouter.HandleFunc("/articles/{id}", deleteArticleById).Methods("DELETE")
	myRouter.HandleFunc("/articles", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles", updateArticleById).Methods("PUT")
	address := fmt.Sprintf("%s:%s", string(configpkg.Host), strconv.Itoa(int(configpkg.Port)))
	log.Fatal(http.ListenAndServe(address, myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	Articles = []Article{
		{Id: 1, Title: "Đây là title 1", Desc: "Mô tả của title 1", Content: "Còn đây là nội dung của Title 1", Deleted: false},
		{Id: 2, Title: "Đây là title 2", Desc: "Mô tả của title 2", Content: "Còn đây là nội dung của Title 2", Deleted: false},
	}
	configpkg.Configpkg()
	fmt.Println("Host: ", configpkg.Host)
	fmt.Println("Port: ", configpkg.Port)
	fmt.Println("Username: ", configpkg.Username)
	fmt.Println("Password: ", configpkg.Password)
	handleRequests()
}
