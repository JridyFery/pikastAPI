package models

import "github.com/jinzhu/gorm"

//Pokemon Struct
type Pokemon struct {
	gorm.Model
	PokemonName       string    `json:"pokemon_name"`
	PokemonPrefab     string    `json:"pokemon_prefab"`
	PokemonFBX        string    `json:"pokemon_fbx"`
	PokemonImg        string    `json:"pokemon_img"`
	PokemonisPremium  bool      `json:"pokemonis_premium"`
	PokemonCost       int       `json:"pokemon_cost"`
	WithDiamonds      bool      `json:"with_diamonds"`
	PokemonPower      float64   `json:"pokemon_power"`
	PokemonMaxPower      float64   `json:"pokemon_maxpower"`
	PokemonAttacktype string    `json:"pokemon_attack"`
	PokemonATKSpeed   float64   `json:"pokemon_atkspeed"`
	PokemonMOVSpeed   float64   `json:"pokemon_movspeed"`
	PokemonHeight     float64   `json:"height"`
	PokemonWidth      float64   `json:"width"`
	Player            []*Player `gorm:"many2many:player_pokemons;"`
}

//PokemonResponse Struct
type PokemonResponse struct {
	gorm.Model
	PokemonName       string    `json:"pokemon_name"`
	PokemonPrefab     string    `json:"pokemon_prefab"`
	PokemonFBX        string    `json:"pokemon_fbx"`
	PokemonImg        []byte    `json:"pokemon_img"`
	PokemonisPremium  bool      `json:"pokemonis_premium"`
	PokemonCost       int       `json:"pokemon_cost"`
	WithDiamonds      bool      `json:"with_diamonds"`
	PokemonPower      float64   `json:"pokemon_power"`
	PokemonMaxPower      float64   `json:"pokemon_maxpower"`
	PokemonAttacktype string    `json:"pokemon_attack"`
	PokemonATKSpeed   float64   `json:"pokemon_atkspeed"`
	PokemonMOVSpeed   float64   `json:"pokemon_movspeed"`
	PokemonHeight     float64   `json:"height"`
	PokemonWidth      float64   `json:"width"`
}