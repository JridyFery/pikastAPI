package repository

import (
	//"crypto/sha1"
	"errors"
	//"strings"

	"github.com/jinzhu/gorm"
	models "github.com/pikastAR/pikastAPI/models"
)

// UserRepository ...
type UserRepository interface {
	CreateUser(u models.User) (models.User, error)
}

// UserRepo ...
type UserRepo struct {
	Db *gorm.DB
}

// CreateUser ...
func (r *UserRepo) CreateUser(u models.User) (models.User, error) {
	User := u
	var user models.User
	err := r.Db.Where(map[string]interface{}{"name": u.Name}).Find(&user).Error
	if err == nil {
		return user, errors.New("ERROR: name already used")
	}
	err = r.Db.Where(map[string]interface{}{"email": u.Email}).Find(&user).Error
	if err == nil {
		return user, errors.New("ERROR: mail already used")
	}
	err = r.Db.Create(&User).Error
	return User, err
}
