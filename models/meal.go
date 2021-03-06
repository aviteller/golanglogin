package models

import (
	"github.com/jinzhu/gorm"
)

type MealType int

const (
	Breakfast MealType = 0
	Lunch     MealType = 1
	Supper    MealType = 2
	Snack     MealType = 3
)

//Meal struct declaration
type Meal struct {
	gorm.Model
	UserID      int    `json:"UserID"`
	Name        string `json:"Name"`
	MealType    int    `json:"Mealtype"`
	Calories    int    `json:"Calories"`
	Ingredients []Ingredient
}

type Ingredient struct {
	gorm.Model
	MealID   int    `json:"MealID"`
	Name     string `json:"Name"`
	Calories int    `json:"Calories"`
}

type MealRating struct {
	gorm.Model
	EatenMealID int  `json:"EatenMealID"`
	PersonID    int  `json:"PersonID"`
	Ate         bool `json:"Ate"`
	Rating      int  `json:"Rating"`
	EatenMeal   EatenMeal
}

type EatenMeal struct {
	gorm.Model
	UserID        int     `json:"UserID"`
	MealID        int     `json:"MealID"`
	Date          string  `json:"Date"`
	AverageRating float32 `json:"AverageRating"`
	MealRatings   []MealRating
	Meal          Meal
}
