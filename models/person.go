package models

import (
	"github.com/jinzhu/gorm"
)

//User struct declaration
type Person struct {
	gorm.Model

	UserID int    `json:"UserID"`
	Name   string `json:"Name""`
	Age    int    `json:"Age"`
	Gender string `json:"Gender"`
}
