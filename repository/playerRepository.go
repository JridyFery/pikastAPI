package repository

import (
	"crypto/sha1"
	"errors"
	"strings"

	models "github.com/JridyFery/pikastAPI/models"
	"github.com/jinzhu/gorm"
)

// PlayerRepository ...
type PlayerRepository interface {
	CreatePlayer(u models.Player) (models.Player, error)
	GetPlayer(id uint) (models.Player, error)
	GetPlayers(role string, offset int, limit int) ([]models.Player, int, error)
	GetPlayerBy(keys []string, values []interface{}) (models.Player, error)
	DeletePlayer(id uint) error
	UpdatePlayer(m map[string]interface{}, id uint) error
	UpdatePlayerPic(img string, id uint) error
	AddPokemonPlayer(idPlayer int, idPokemon int) error
	GetplayerPokemons(playerID uint) ([]models.Pokemon, int, error)
}

// PlayerRepo ...
type PlayerRepo struct {
	Db *gorm.DB
}

// Createplayer ....
func (r *PlayerRepo) Createplayer(p models.Player) (models.Player, error) {
	Player := p
	var player models.Player
	err := r.Db.Where(map[string]interface{}{"player_name": p.PlayerName}).Find(&player).Error
	if err == nil {
		return player, errors.New("ERROR: name already used")
	}
	err = r.Db.Where(map[string]interface{}{"player_email": p.PlayerEmail}).Find(&player).Error
	if err == nil {
		return player, errors.New("ERROR: email is already used")
	}
	err = r.Db.Create(&Player).Error
	return Player, err
}

//GetPlayer ...
func (r *PlayerRepo) GetPlayer(id uint) (models.Player, error) {
	var Player models.Player
	err := r.Db.First(&Player, id).Error
	return Player, err
}

// GetPlayers ...
func (r *PlayerRepo) GetPlayers(role string, offset int, limit int) ([]models.Player, int, error) {
	var Players []models.Player
	var Player models.Player
	var count int
	var err error
	if strings.ToUpper(role) == "USER" {
		err = r.Db.Where("admin = ?", false).Offset(offset).Limit(limit).Find(&Players).Error
		r.Db.Model(&Player).Where("admin = ?", false).Count(&count)
	} else if strings.ToUpper(role) == "ADMIN" {
		err = r.Db.Where("admin = ? ", true).Offset(offset).Limit(limit).Find(&Players).Error
		r.Db.Model(&Player).Where("admin = ? ", true).Count(&count)
	}
	return Players, count, err
}

// GetPlayerBy ...
func (r *PlayerRepo) GetPlayerBy(keys []string, values []interface{}) (models.Player, error) {
	var Player models.Player
	var m map[string]interface{}
	var password string
	m = make(map[string]interface{})

	for index, value := range keys {
		if value == "player_password" {
			crypt := sha1.New()
			password = values[index].(string)
			crypt.Write([]byte(password))
			m[value] = crypt.Sum(nil)
		} else {
			m[value] = values[index]
		}
	}
	err := r.Db.Where(m).Find(&Player).Error
	return Player, err
}

//UpdatePlayer ...
func (r *PlayerRepo) UpdatePlayer(m map[string]interface{}, id uint) error {
	player := models.Player{}
	err := r.Db.Where("name = ? AND id != ?", m["name"], id).Find(&player).Error
	if err == nil {
		return errors.New("ERROR: name already used")
	}
	err = r.Db.Where("email = ? AND id != ?", m["email"], id).Find(&player).Error
	if err == nil {
		return errors.New("ERROR: MAIL ALREADY USED")
	}
	err = r.Db.First(&player, id).Error
	if err != nil {
		return err
	}
	player.ID = id
	err1 := r.Db.Model(&player).Updates(m).Error
	return err1

}

//UpdatePlayerPic function
func (r *PlayerRepo) UpdatePlayerPic(img string, id uint) error {
	Player := models.Player{}
	m := make(map[string]interface{})
	err := r.Db.First(&Player, id).Error
	if err != nil {
		return err
	}
	m["player_img"] = img
	Player.ID = id
	err = r.Db.Model(&Player).Update(m).Error
	return err
}

//DeletePlayer ...
func (r *PlayerRepo) DeletePlayer(id uint) error {
	player := models.Player{}
	err := r.Db.First(&player, id).Error
	if err != nil {
		return err
	}
	player.ID = id
	err = r.Db.Delete(&player).Error
	return err

}

//AddPokemonPlayer ...
func (r *PlayerRepo) AddPokemonPlayer(idPlayer int, idPokemon int) error {
	PlayerPokemon := models.PlayerPokemon{}
	PlayerPokemon.PlayerID = uint(idPlayer)
	PlayerPokemon.PokemonID = uint(idPokemon)

	//m := make(map[string]interface{})
	err := r.Db.Where("pokemon_id = ? AND player_id = ?", idPokemon, idPlayer).Find(&PlayerPokemon).Error
	if err == nil {
		return errors.New("ERROR: POKEMON IS ALREADY OWNED")
	}
	//Supposed that we'll simply purchase pokemons using coins
	Pokemon := models.Pokemon{}
	Player := models.Player{}
	//Get player with his current id
	err = r.Db.First(&Player, idPlayer).Error
	if err != nil {
		return err
	}
	//Get pokemon with his current id
	err = r.Db.First(&Pokemon, idPokemon).Error
	if err != nil {
		return err
	}
	err1 := r.Db.Create(&PlayerPokemon).Error
	//Checking buying Method
	if Pokemon.PokemonisPremium {
		if Pokemon.WithDiamonds {
			//Checking buying possibility
			if Player.PlayerDiamonds < Pokemon.PokemonCost {
				return (errors.New("NOT ENOUGH DIAMONDS"))
			}
			//Update player coins in DB
			Player.PlayerDiamonds -= Pokemon.PokemonCost
		} else {
			if Player.PlayerCoins < Pokemon.PokemonCost {
				return (errors.New("NOT ENOUGH MONEY"))
			}
			//Update player coins in DB
			Player.PlayerCoins -= Pokemon.PokemonCost
		}
		r.Db.Save(&Player)
	}
	return err1
}

//GetplayerPokemons func
func (r *PlayerRepo) GetplayerPokemons(playerID uint) ([]models.Pokemon, int, error) {
	var Pokemons []models.Pokemon
	Player := models.Player{}
	Player.ID = playerID
	var count int
	var err error
	err = r.Db.Model(&Player).Related(&Pokemons, "Pokemon").Error
	count = len(Pokemons)
	return Pokemons, count, err
}
