package models

import (
	"github.com/jinzhu/gorm"
)

// User Struct
type User struct {
	gorm.Model
	Name				string 				`json:"name"`
	Email				string 				`json:"email"`
	Password			[]byte 				`json:"password"`
	Admin				bool 				`json:"admin"`
	SuperAdmin			bool 				`json:"superAdmin"`
	BirthDay 			int					`json:"birthDay"`
	BirthMonth 			string				`json:"birthMonth"`
	BirthYear 			int					`json:"birthYear"`
	Country 			string				`json:"countryOfResidence"`
}

// UserRequest Struct
type UserRequest struct {
	Name				string 				`json:"name"`
	Email				string 				`json:"email"`
	Password			string 				`json:"password"`
	Admin				bool 				`json:"admin"`
	SuperAdmin			bool 				`json:"superAdmin"`
	BirthDay 			int					`json:"birthDay"`
	BirthMonth 			string				`json:"birthMonth"`
	BirthYear 			int					`json:"birthYear"`
	Country 			string				`json:"countryOfResidence"`
}