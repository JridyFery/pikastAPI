package models

import (
	"github.com/jinzhu/gorm"
)

// Player Struct
type Player struct {
	gorm.Model
	PlayerName          string     `json:"player_name"`
	PlayerTel           string     `json:"player_tel"`
	PlayerEmail         string     `json:"player_email"`
	PlayerPassword      []byte     `json:"password"`
	Admin               bool       `json:"admin"`
	BirthDay            int        `json:"birthDay"`
	BirthMonth          string     `json:"birthMonth"`
	BirthYear           int        `json:"birthYear"`
	Country             string     `json:"countryOfResidence"`
	PlayerImg           string     `json:"player_img"`
	PlayerCoins         int        `json:"player_coins"`
	PlayerRank          string     `json:"player_rank"`
	PlayerLevelCount    int        `json:"player_levelCount"`
	PlayerLevelProgress int        `json:"player_levelProgress"`
	Pokemon             []*Pokemon `gorm:"many2many:player_pokemons;"`
}

//PlayerPokemon ..
type PlayerPokemon struct {
	PokemonID uint `json:"PokemonID"`
	PlayerID  uint `json:"PlayerID"`
}

// PlayerRequest Struct
type PlayerRequest struct {
	PlayerName          string `json:"player_name"`
	PlayerTel           string `json:"player_tel"`
	PlayerEmail         string `json:"player_email"`
	PlayerPassword      string `json:"password"`
	Admin               bool   `json:"admin"`
	BirthDay            int    `json:"birthDay"`
	BirthMonth          string `json:"birthMonth"`
	BirthYear           int    `json:"birthYear"`
	Country             string `json:"countryOfResidence"`
	PlayerImg           string `json:"player_img"`
	PlayerCoins         int    `json:"player_coins"`
	PlayerRank          string `json:"player_rank"`
	PlayerLevelCount    int    `json:"player_levelCount"`
	PlayerLevelProgress int    `json:"player_levelProgress"`
}

//PlayerResponse Struct
type PlayerResponse struct {
	PlayerID            uint   `json:"player_id"`
	PlayerName          string `json:"player_name"`
	PlayerTel           string `json:"player_tel"`
	PlayerEmail         string `json:"player_email"`
	Admin               bool   `json:"admin"`
	DateOfBirth         Date   `json:"dateofbirth"`
	Country             string `json:"countryOfResidence"`
	PlayerImg           string `json:"player_img"`
	PlayerCoins         int    `json:"player_coins"`
	PlayerRank          string `json:"player_rank"`
	PlayerLevelCount    int    `json:"player_levelCount"`
	PlayerLevelProgress int    `json:"player_levelProgress"`
}
