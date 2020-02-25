package models

import (
	"github.com/jinzhu/gorm"
)

//User struct declaration
type User struct {
	gorm.Model

	Name     string `gorm:"type:varchar(100);unique_index"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
	Person   []Person
}
