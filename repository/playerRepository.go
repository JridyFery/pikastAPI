package repository

import (
	"errors"
	"net/http"

	"github.com/jinzhu/gorm"
	models "github.com/pikastAR/pikastAPI/models"
)

// PlayerRepository ...
type PlayerRepository interface {
	CreatePlayer(u models.Player) (models.Player, error)
	GetPlayer(id uint) (models.Player, error)
	DeletePlayer(id uint) error
	UpdatePlayer(w http.ResponseWriter, r *http.Request)
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

//to explain !!! fasserhouli ya 3ala ya5raaa
//UpdatePlayer ...
func (r *PlayerRepo) UpdatePlayer(m map[string]interface{}, id uint) error {
	player := models.Player{}
	err := r.Db.Where("name = ? AND id != ?", m["name"], id).Find(&player).Error
	if err == nil {
		return errors.New("ERROR: name already used")
	}
	err = r.Db.Where("email = ? AND id != ?", m["email"], id).Find(&player).Error
	if err == nil {
		return errors.New("ERROR: mail already used")
	}
	err = r.Db.First(&player, id).Error
	if err != nil {
		return err
	}
	player.ID = id
	err1 := r.Db.Model(&player).Updates(m).Error
	return err1

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
