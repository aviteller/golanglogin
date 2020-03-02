package models

import (
	"github.com/jinzhu/gorm"
)

//Person struct declaration
type Person struct {
	gorm.Model

	UserID      int    `json:"UserID"`
	Name        string `json:"Name"`
	Gender      string `json:"Gender"`
	DateOfBirth string `json:"DateOfBirth"`
	MealRatings []MealRating
}
