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
	//Defining the routes
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", heyThere).Methods("GET")
	myRouter.HandleFunc("/users", Users).Methods("GET")
	myRouter.HandleFunc("/user/{name}/{email}", NewUser).Methods("POST")
	myRouter.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))

}

func main(){
	fmt.Println("First GO ORM application")
	InitialMigration()
	handleRequests()
}