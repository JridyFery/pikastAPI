package repository

import (
	"errors"

	"github.com/jinzhu/gorm"
	models "github.com/pikastAR/pikastAPI/models"
)

//PokemonRepository ...
type PokemonRepository interface {
	CreatePokemon(p models.Pokemon) (models.Pokemon, error)
	GetPokemon(id uint) (models.Pokemon, error)
	DeletePokemon(id uint) error
	//get pok by
	//delete pok
	//update pok
}

//PokemonRepo ...
type PokemonRepo struct {
	Db *gorm.DB
}

//CreatePokemon func
func (r *PokemonRepo) CreatePokemon(p models.Pokemon) (models.Pokemon, error) {
	Pokemon := p
	var pokemon models.Pokemon
	err := r.Db.Where(map[string]interface{}{"name": p.PokemonName}).Find(&pokemon).Error
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
