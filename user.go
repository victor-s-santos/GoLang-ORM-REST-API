package main

import (
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/sqlite"
	"encoding/json"
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
	db.AutoMigrate(&User{})
}

func Users(w http.ResponseWriter, r *http.Request){
	//fmt.Fprintf(w, "Endpoint users")
	db, err = gorm.Open("sqlite3", "first_test.db")
	if err != nil{
		panic("Could not connect to the database")
	}
	defer db.Close()

	var users []User
	db.Find(&users)

	json.NewEncoder(w).Encode(users)
}

func NewUser(w http.ResponseWriter, r *http.Request){
	db, err = gorm.Open("sqlite3", "first_test.db")
	if err != nil{
		panic("Could not connect to the database")
	}
	defer db.Close()
	//fmt.Fprintf(w, "New User Endpoint")
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Delete User Endpoint")
}

func UpdateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Update User Endpoint")
}