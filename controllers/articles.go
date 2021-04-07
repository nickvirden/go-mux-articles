package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nickvirden/go-mux-articles/models"

	"github.com/gorilla/mux"
)

var Articles []models.Article

func InitializeArticles() {
	Articles = []models.Article{
		models.Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		models.Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
}

func HomePage(http_response_writer http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(http_response_writer, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: HomePage")
}

func ReturnAllArticles(http_response_writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Endpoint Hit: ReturnAllArticles")
	json.NewEncoder(http_response_writer).Encode(Articles)
}

func ReturnSingleArticle(http_response_writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(http_response_writer).Encode(article)
		}
	}
}

func CreateNewArticle(http_response_writer http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)

	var article models.Article

	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)

	json.NewEncoder(http_response_writer).Encode(article)
}

func DeleteArticle(http_response_writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}
