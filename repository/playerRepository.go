package repository

import (
	"crypto/sha1"
	"errors"
	"net/http"
	"strings"

	"github.com/jinzhu/gorm"
	models "github.com/pikastAR/pikastAPI/models"
)

// PlayerRepository ...
type PlayerRepository interface {
	CreatePlayer(u models.Player) (models.Player, error)
	GetPlayer(id uint) (models.Player, error)
	GetPlayers(role string, offset int, limit int) ([]models.Player, int, error)
	GetPlayerBy(keys []string, values []interface{}) (models.Player, error)
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
