package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:G8keeper@mysql@tcp(127.0.0.1:3306)/gorrila-mux-api?charset=utf8mb4&parseTime=True&loc=Local"


type User struct{
	gorm.Model  // Converting the struct into ORM Model
	FirstName  string `json:"firstname"`
	LastName   string  `json:"lastname"`
	Email      string   `json:"email"`
}

func InitialMigration() {
	// 1. Connect to a database
	// 2. Auto migrate data frm the DB to the struct
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{}) // 1. Connecting to database
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	// 2. Enabling auto migration
	DB.AutoMigrate(&User{})
}

// GetUsers gets the list of users available
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	//Fetching data from Db
	DB.Find(&users) 
	// Encode fetched data & send it back to the server
	json.NewEncoder(w).Encode(users)

}

// GetUser gets a particular user
func GetUser(w http.ResponseWriter, r *http.Request) {
	// Set content type to application json
	w.Header().Set("Content-Type", "application/json")
	// Get the user id from the parameter
	params := mux.Vars(r)
	var user User
	// Find the data using the id from params
	DB.First(&user, params["id"])
	// Encode fetched data & send it back to the server
	json.NewEncoder(w).Encode(user)
	
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	 // Set content type to application json
	w.Header().Set("content-Type", "application/json")
	var user User
	// Use json model to decode the data gotten from the request body
	// To the user
	json.NewDecoder(r.Body).Decode(&user) 
	// save the data decoded
	DB.Create(&user)
	// Pass the saved data back to the browser
	json.NewEncoder(w).Encode(user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)
	json.NewEncoder(w).Encode(&user)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The User is deleted successfully")

}