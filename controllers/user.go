package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"../models"
	"../utils"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var db = utils.ConnectDB()

func Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp := FindOneUser(user.Name, user.Password)
	json.NewEncoder(w).Encode(resp)
}

func FindOneUser(name, password string) map[string]interface{} {
	user := &models.User{}

	if err := db.Where("Name = ?", name).First(user).Error; err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Email address not found"}
		return resp
	}

	// 3 hours
	expiresAt := time.Now().Add(time.Minute * 180).Unix()

	errf := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		var resp = map[string]interface{}{"status": false, "message": "Invalid login credentials. Please try again"}
		return resp
	}

	tk := &models.Token{
		UserID: user.ID,
		Name:   user.Name,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, error := token.SignedString([]byte("secret"))
	if error != nil {
		fmt.Println(error)
	}

	var resp = map[string]interface{}{"status": false, "message": "logged in"}
	resp["token"] = tokenString //Store the token in the response
	resp["user"] = map[string]interface{}{"id": user.ID, "name": user.Name, "email": user.Email}
	return resp
}

//CreateUser function -- create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	json.NewDecoder(r.Body).Decode(user)

	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		err := ErrorResponse{
			Err: "Password Encryption  failed",
		}
		json.NewEncoder(w).Encode(err)
	}

	user.Password = string(pass)

	createdUser := db.Create(user)
	var errMessage = createdUser.Error

	if createdUser.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdUser)
}

//FetchUser function
func FetchUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.Select("id, name ,email").Find(&users)

	// preload lets you get child structs easier
	// db.Preload("Person").Find(&users)

	json.NewEncoder(w).Encode(users)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	params := mux.Vars(r)
	var id = params["id"]
	db.First(&user, id)
	json.NewDecoder(r.Body).Decode(user)
	db.Save(&user)
	json.NewEncoder(w).Encode(&user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var user models.User
	db.First(&user, id)
	db.Delete(&user)
	json.NewEncoder(w).Encode("User deleted")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	// var onlyuser = params["onlyuser"]
	var user models.User

	var eatenMeals []models.EatenMeal

	db.Preload("People").Preload("Meals").Select("id, name ,email").First(&user, id)
	db.Preload("Meal").Where("user_id = ?", id).Find(&eatenMeals)

	var resp = map[string]interface{}{}
	resp["user"] = &user
	resp["eatenMeals"] = &eatenMeals

	json.NewEncoder(w).Encode(resp)
}
