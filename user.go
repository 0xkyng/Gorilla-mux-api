package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:G8keeper@mysql@tcp(127.0.0.1:3306)/gorrila-mux-api?charset=utf8mb4&parseTime=True&loc=Local"


type User struct{
	gorm.Model
	FirstName  string `json:"firstname"`
	LastName   string  `json:"lastname"`
	Email      string   `json:"email"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&User{})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	// save the data decoded
	DB.Create(&user)
	// Pass the data to the browser
	json.NewEncoder(w).Encode(user)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

}