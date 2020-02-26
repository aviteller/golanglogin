package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"github.com/gorilla/mux"
)

func CreateIngredient(w http.ResponseWriter, r *http.Request) {

	ingredient := &models.Ingredient{}
	json.NewDecoder(r.Body).Decode(ingredient)

	createdIngredient := db.Create(ingredient)
	var errMessage = createdIngredient.Error

	var meal models.Meal
	db.First(&meal, ingredient.MealID)

	// currentMealCalories := meal.Calories

	meal.Calories += ingredient.Calories
	db.Save(&meal)

	// update meal calories

	if createdIngredient.Error != nil {
		fmt.Println(errMessage)
	}
	json.NewEncoder(w).Encode(createdIngredient)
}

func UpdateIngredient(w http.ResponseWriter, r *http.Request) {
	ingredient := &models.Ingredient{}
	params := mux.Vars(r)
	var id = params["id"]
	db.First(&ingredient, id)
	json.NewDecoder(r.Body).Decode(ingredient)
	db.Save(&ingredient)
	json.NewEncoder(w).Encode(&ingredient)
}

func DeleteIngredient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var ingredient models.Ingredient
	db.First(&ingredient, id)

	var meal models.Meal
	db.First(&meal, ingredient.MealID)

	meal.Calories -= ingredient.Calories
	db.Save(&meal)

	db.Delete(&ingredient)
	json.NewEncoder(w).Encode("Ingredient deleted")
}

func GetIngredient(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var ingredient models.Ingredient
	db.First(&ingredient, id)
	json.NewEncoder(w).Encode(&ingredient)
}
