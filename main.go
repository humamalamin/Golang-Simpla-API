package main

import
(
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"gorilla/mux"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func returnAllArticles(w http.ResponseWriter,r *http.Request){
	articles := Articles{
		Article{Title: "Hello",Desc: "Article Description",Content: "Articles Content"},
		Article{Title: "Hello 2",Desc: "Article Description 2", Content: "Articles Content 2"},
	}
	fmt.Println("Endpoin Hit : returnAllArticles")
	
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests(){

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/",homePage)
	myRouter.HandleFunc("/all",returnAllArticles)
	log.Fatal(http.ListenAndServe(":8081",myRouter))
}

func main(){
	fmt.Println("Rest API V2.0 - Mux Routers")
	handleRequests()
}