package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func heyThere(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hey There")
}

func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", heyThere).Methods("GET")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}

func main(){
	fmt.Println("First GO ORM application")

	handleRequests()
}