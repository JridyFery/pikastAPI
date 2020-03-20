package models

import "github.com/jinzhu/gorm"

//Pokemon Struct
type Pokemon struct {
	gorm.Model
	PokemonName       string    `json:"pokemon_name"`
	PokemonPrefab     string    `json:"pokemon_prefab"`
	PokemonImg        string    `json:"pokemon_img"`
	PokemonisPremium  bool      `json:"pokemon_isfree"`
	PokemonPoints     int       `json:"pokemon_points"`
	PokemonPower      float64   `json:"pokemon_power"`
	PokemonAttacktype string    `json:"pokemon_attack"`
	PokemonATKSpeed   float64   `json:"pokemon_atkspeed"`
	PokemonMOVSpeed   float64   `json:"pokemon_movspeed"`
	PokemonHeight     float64   `json:"height"`
	PokemonWidth      float64   `json:"width"`
	Player            []*Player `gorm:"many2many:player_pokemon;"`
}

//PokemonRequest struct
type PokemonRequest struct {
	PokemonName       string  `json:"pokemon_name"`
	PokemonPrefab     string  `json:"pokemon_prefab"`
	PokemonImg        string  `json:"pokemon_img"`
	PokemonisPremium  bool    `json:"pokemon_isfree"`
	PokemonPoints     int     `json:"pokemon_points"`
	PokemonPower      float64 `json:"pokemon_power"`
	PokemonAttacktype string  `json:"pokemon_attack"`
	PokemonATKSpeed   float64 `json:"pokemon_atkspeed"`
	PokemonMOVSpeed   float64 `json:"pokemon_movspeed"`
	PokemonHeight     float64 `json:"height"`
	PokemonWidth      float64 `json:"width"`
}

//PokemonResponse struct
type PokemonResponse struct {
	PokemonName       string  `json:"pokemon_name"`
	PokemonPrefab     string  `json:"pokemon_prefab"`
	PokemonImg        string  `json:"pokemon_img"`
	PokemonisPremium  bool    `json:"pokemon_isfree"`
	PokemonPoints     int     `json:"pokemon_points"`
	PokemonPower      float64 `json:"pokemon_power"`
	PokemonAttacktype string  `json:"pokemon_attack"`
	PokemonATKSpeed   float64 `json:"pokemon_atkspeed"`
	PokemonMOVSpeed   float64 `json:"pokemon_movspeed"`
	PokemonHeight     float64 `json:"height"`
	PokemonWidth      float64 `json:"width"`
}
