package repository

import (
	"errors"
	"strings"

	"github.com/jinzhu/gorm"
	models "github.com/pikastAR/pikastAPI/models"
)

//PokemonRepository ...
type PokemonRepository interface {
	CreatePokemon(p models.Pokemon) (models.Pokemon, error)
	GetPokemon(id uint) (models.Pokemon, error)
	GetPokemons(pokemonType string, offset int, limit int) ([]models.Pokemon, int, error)
	UpdatePokemon(m map[string]interface{}, id uint) error
	DeletePokemon(id uint) error
}

//PokemonRepo ...
type PokemonRepo struct {
	Db *gorm.DB
}

//CreatePokemon func
func (r *PokemonRepo) CreatePokemon(p models.Pokemon) (models.Pokemon, error) {
	Pokemon := p
	var pokemon models.Pokemon
	err := r.Db.Where(map[string]interface{}{"pokemon_name": p.PokemonName}).Find(&pokemon).Error
	if err == nil {
		return pokemon, errors.New("ERROR: name already used")
	}
	err = r.Db.Create(&Pokemon).Error
	return Pokemon, err
}

//GetPokemon ...
func (r *PokemonRepo) GetPokemon(id uint) (models.Pokemon, error) {
	var Pokemon models.Pokemon
	err := r.Db.First(&Pokemon, id).Error
	return Pokemon, err
}

// GetPokemons ...
func (r *PokemonRepo) GetPokemons(pokemonType string, offset int, limit int) ([]models.Pokemon, int, error) {
	var Pokemons []models.Pokemon
	var Pokemon models.Pokemon
	var count int
	var err error
	if strings.ToUpper(pokemonType) == "PREMIUM" {
		err = r.Db.Where("pokemonis_premium= ?", true).Offset(offset).Limit(limit).Find(&Pokemons).Error
		r.Db.Model(&Pokemon).Where("pokemonis_premium = ?", true).Count(&count)
	} else if strings.ToUpper(pokemonType) == "FREE" {
		err = r.Db.Where("pokemonis_premium= ?", false).Offset(offset).Limit(limit).Find(&Pokemons).Error
		r.Db.Model(&Pokemon).Where("pokemonis_premium = ?", false).Count(&count)
	}
	return Pokemons, count, err
}

//DeletePokemon ...
func (r *PokemonRepo) DeletePokemon(id uint) error {
	pokemon := models.Pokemon{}
	err := r.Db.First(&pokemon, id).Error
	if err != nil {
		return err
	}
	pokemon.ID = id
	err = r.Db.Delete(&pokemon).Error
	return err

}

//UpdatePokemon ...
func (r *PokemonRepo) UpdatePokemon(m map[string]interface{}, id uint) error {
	pokemon := models.Pokemon{}
	err := r.Db.Where("pokemon_name = ?", m["pokemon_name"]).Find(&pokemon).Error
	if err == nil {
		return errors.New("ERROR: name is already in use")
	}
	err = r.Db.First(&pokemon, id).Error
	if err != nil {
		return errors.New("ERROR: ID does not exist")
	}
	pokemon.ID = id
	err1 := r.Db.Model(&pokemon).Updates(m).Error
	return err1
}
