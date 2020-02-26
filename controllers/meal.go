package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"github.com/gorilla/mux"
)

func CreateMeal(w http.ResponseWriter, r *http.Request) {

	meal := &models.Meal{}
	json.NewDecoder(r.Body).Decode(meal)

	createdMeal := db.Create(meal)
	var errMessage = createdMeal.Error

	if createdMeal.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdMeal)
}

func UpdateMeal(w http.ResponseWriter, r *http.Request) {
	meal := &models.Meal{}
	params := mux.Vars(r)
	var id = params["id"]
	db.First(&meal, id)
	json.NewDecoder(r.Body).Decode(meal)
	db.Save(&meal)
	json.NewEncoder(w).Encode(&meal)
}

func DeleteMeal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var meal models.Meal
	db.First(&meal, id)
	db.Delete(&meal)
	json.NewEncoder(w).Encode("Meal deleted")
}

func GetMeal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var meal models.Meal
	db.Preload("Ingredients").First(&meal, id)
	json.NewEncoder(w).Encode(&meal)
}
