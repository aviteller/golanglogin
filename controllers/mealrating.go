package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"github.com/gorilla/mux"
)

func CreateMealRating(w http.ResponseWriter, r *http.Request) {

	mealRating := &models.MealRating{}
	json.NewDecoder(r.Body).Decode(mealRating)

	createdMealRating := db.Create(mealRating)
	var errMessage = createdMealRating.Error

	if createdMealRating.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdMealRating)
}

func UpdateMealRating(w http.ResponseWriter, r *http.Request) {
	mealRating := &models.MealRating{}
	params := mux.Vars(r)
	var id = params["id"]
	db.First(&mealRating, id)
	json.NewDecoder(r.Body).Decode(mealRating)
	db.Save(&mealRating)
	json.NewEncoder(w).Encode(&mealRating)
}

func DeleteMealRating(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var mealRating models.MealRating
	db.First(&mealRating, id)
	db.Delete(&mealRating)
	json.NewEncoder(w).Encode("MealRating deleted")
}

func GetMealRating(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var mealRating models.MealRating

	db.First(&mealRating, id)

	json.NewEncoder(w).Encode(&mealRating)

}
