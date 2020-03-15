package helpers

import (
	"crypto/sha1"

	models "github.com/pikastAR/pikastAPI/models"
)

// UserResponseFormatter func
func UserResponseFormatter(result models.User, user *models.UserResponse) {
	user.ID = result.ID
	user.Name = result.Name
	user.Email = result.Email
	user.Roles = append(user.Roles, "user")
	if result.Admin {
		user.Roles = append(user.Roles, "admin")
	}
	if result.SuperAdmin {
		user.Roles = append(user.Roles, "super admin")
	}
	user.DateOfBirth.Day = result.BirthDay
	user.DateOfBirth.Month = result.BirthMonth
	user.DateOfBirth.Year = result.BirthYear
	user.Country = result.Country
}

// UserRequestFormatter func
func UserRequestFormatter(request models.UserRequest, user *models.User) {
	if request.Password != "" {
		crypt := sha1.New()
		crypt.Write([]byte(request.Password))
		user.Password = crypt.Sum(nil)
	}
	user.Name = request.Name
	user.Email = request.Email
	user.Admin = request.Admin
	user.SuperAdmin = request.SuperAdmin
	user.BirthDay = request.BirthDay
	user.BirthMonth = request.BirthMonth
	user.BirthYear = request.BirthYear
	user.Country = request.Country
}
