package models

import (
	"github.com/jinzhu/gorm"
)

// Player Struct
type Player struct {
	gorm.Model
	PlayerID       string `json:"player_id"`
	PlayerName     string `json:"player_name"`
	PlayerTel      string `json:"player_tel"`
	PlayerEmail    string `json:"player_email"`
	PlayerPassword []byte `json:"password"`
	Admin          bool   `json:"admin"`
	BirthDay       int    `json:"birthDay"`
	BirthMonth     string `json:"birthMonth"`
	BirthYear      int    `json:"birthYear"`
	Country        string `json:"countryOfResidence"`
	PlayerImg      string `json:"player_img"`
}

// PlayerRequest Struct
type PlayerRequest struct {
	PlayerID       string `json:"player_id"`
	PlayerName     string `json:"player_name"`
	PlayerTel      string `json:"player_tel"`
	PlayerEmail    string `json:"player_email"`
	PlayerPassword []byte `json:"password"`
	Admin          bool   `json:"admin"`
	BirthDay       int    `json:"birthDay"`
	BirthMonth     string `json:"birthMonth"`
	BirthYear      int    `json:"birthYear"`
	Country        string `json:"countryOfResidence"`
	PlayerImg      string `json:"player_img"`
}

//PlayerResponse Struct
type PlayerResponse struct {
	PlayerID       string `json:"player_id"`
	PlayerName     string `json:"player_name"`
	PlayerTel      string `json:"player_tel"`
	PlayerEmail    string `json:"player_email"`
	PlayerPassword []byte `json:"password"`
	Admin          bool   `json:"admin"`
	DateOfBirth    Date   `json:"dateOfBirth"`
	Country        string `json:"countryOfResidence"`
	PlayerImg      string `json:"player_img"`
}
