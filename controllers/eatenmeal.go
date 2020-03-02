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

	db.Preload("Meal").Find(&eatenMeal)
	json.NewEncoder(w).Encode(&eatenMeal)
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

	// b, err := json.MarshalIndent(mealRatings, "", "  ")
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }
	// fmt.Print(string(b))
	db.Where("eaten_meal_id = ?", id).Delete(&[]models.MealRating{})
	db.Delete(&eatenMeal)

	json.NewEncoder(w).Encode("EatenMeal deleted")
}

func GetEatenMeal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	var eatenMeal models.EatenMeal
	var people []models.Person

	db.Preload("Meal").Preload("MealRatings").First(&eatenMeal, id)
	db.Where("user_id = ?", eatenMeal.UserID).Find(&people)

	var resp = map[string]interface{}{}
	resp["eatenMeal"] = &eatenMeal
	resp["people"] = &people
	json.NewEncoder(w).Encode(resp)

}

func UpdateEatenMealAverageRating(id int) {
	var eatenMeal models.EatenMeal
	var mealRatings []models.MealRating

	db.Where("id = ?", id).Find(&eatenMeal)
	db.Where("eaten_meal_id = ? AND Ate = 1", id).Find(&mealRatings)

	var averageTotal float32 = 0
	var total int = 0

	for _, meal := range mealRatings {
		total += meal.Rating
	}

	averageTotal = (float32(total) / float32(len(mealRatings)))

	eatenMeal.AverageRating = averageTotal
	db.Save(&eatenMeal)
}
