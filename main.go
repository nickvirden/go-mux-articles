package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nickvirden/go-mux-articles/controllers"

	"github.com/gorilla/mux"
)

func main() {

	controllers.InitializeArticles()

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/home", controllers.HomePage)
	router.HandleFunc("/articles", controllers.ReturnAllArticles)
	router.HandleFunc("/article", controllers.CreateNewArticle).Methods("POST")
	router.HandleFunc("/article/{id}", controllers.ReturnSingleArticle).Methods("GET")
	router.HandleFunc("/article/{id}", controllers.DeleteArticle).Methods("DELETE")

	fmt.Println("Running server on port 8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
