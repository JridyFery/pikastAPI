package repository

import (
	//"crypto/sha1"
	"errors"
	//"strings"

	"github.com/jinzhu/gorm"
	models "github.com/pikastAR/pikastAPI/models"
)

// PlayerRepository ...
type PlayerRepository interface {
	CreatePlayer(u models.Player) (models.Player, error)
	GetPlayer(id uint) (models.Player, error)
}

// PlayerRepo ...
type PlayerRepo struct {
	Db *gorm.DB
}

// Createplayer ....
func (r *PlayerRepo) Createplayer(p models.Player) (models.Player, error) {
	Player := p
	var player models.Player
	err := r.Db.Where(map[string]interface{}{"name": p.PlayerName}).Find(&player).Error
	if err == nil {
		return player, errors.New("ERROR: name already used")
	}
	err = r.Db.Where(map[string]interface{}{"email": p.PlayerEmail}).Find(&player).Error
	if err == nil {
		return player, errors.New("ERROR: mail already used")
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

//GetAll players...
//func (r *PlayerRepo) GetAll(p models.Player) (models.Player, error) {}

//UpdatePlayer ...
//func (r *PlayerRepo) UpdatePlayer(p models.Player) (models.Player, error) {}

//DeletePlayer ...
//func (r *PlayerRepo) DeletePlayer(p models.Player) (models.Player, error) {}
