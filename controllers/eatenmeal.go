package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"github.com/gorilla/mux"
)

func CreateEatenMeal(w http.ResponseWriter, r *http.Request) {

	eatenMeal := &models.EatenMeal{}
	json.NewDecoder(r.Body).Decode(eatenMeal)

	createdEatenMeal := db.Create(eatenMeal)
	var errMessage = createdEatenMeal.Error

	if createdEatenMeal.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdEatenMeal)
}

func UpdateEatenMeal(w http.ResponseWriter, r *http.Request) {
	eatenMeal := &models.EatenMeal{}
	params := mux.Vars(r)
	var id = params["id"]
	db.First(&eatenMeal, id)
	json.NewDecoder(r.Body).Decode(eatenMeal)
	db.Save(&eatenMeal)
	json.NewEncoder(w).Encode(&eatenMeal)
}

func DeleteEatenMeal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var eatenMeal models.EatenMeal
	db.First(&eatenMeal, id)
	db.Delete(&eatenMeal)
	json.NewEncoder(w).Encode("EatenMeal deleted")
}

func GetEatenMeal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var eatenMeal models.EatenMeal
	db.Preload("Meal").First(&eatenMeal, id)
	json.NewEncoder(w).Encode(&eatenMeal)
}
