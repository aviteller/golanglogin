package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"github.com/gorilla/mux"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {

	person := &models.Person{}
	json.NewDecoder(r.Body).Decode(person)

	createdPerson := db.Create(person)
	var errMessage = createdPerson.Error

	if createdPerson.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdPerson)
}

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	person := &models.Person{}
	params := mux.Vars(r)
	var id = params["id"]
	db.First(&person, id)
	json.NewDecoder(r.Body).Decode(person)
	db.Save(&person)
	json.NewEncoder(w).Encode(&person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var person models.Person
	db.First(&person, id)
	db.Delete(&person)
	json.NewEncoder(w).Encode("Person deleted")
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var person models.Person
	// var mealsIds []int

	db.Preload("MealRatings", "Ate = true").Preload("MealRatings.EatenMeal").Preload("MealRatings.EatenMeal.Meal").First(&person, id)

	json.NewEncoder(w).Encode(&person)
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var userID = params["user_id"]
	var people []models.Person

	db.Where("user_id = ?", userID).Find(&people)

	json.NewEncoder(w).Encode(&people)
}
