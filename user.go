package main

import (
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type User struct {
	gorm.Model
	Name string
	Email string
}

func InitialMigration(){
	db, err = gorm.Open("sqlite3", "first_test.db")
	if err != nil{
		fmt.Println(err.Error())
		panic("Error when trying to connect to database.")
	}
	defer db.Close()
	db.Automigrate(&User{})
}

func Users(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Endpoint users")
}

func NewUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "New User Endpoint")
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Delete User Endpoint")
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Update User Endpoint")
}