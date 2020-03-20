package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/pikastAR/pikastAPI/helpers"

	//"strconv"
	models "github.com/pikastAR/pikastAPI/models"
	"github.com/pikastAR/pikastAPI/repository"
	//"crypto/sha1"
	//"gopkg.in/gomail.v2"
	//"github.com/sethvargo/go-password/password"
)

// UserHandler ...
type UserHandler struct {
	Repo repository.UserRepository
}

func responseFormatter(code int, status string, data interface{}, response *models.Response) {
	response.Code = code
	response.Status = status
	response.Data = data
}

// CreateUser ...
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var User models.User
	var UserRequest models.UserRequest
	var response models.Response
	var responseWithToken models.ResponseWithToken
	err := json.NewDecoder(r.Body).Decode(&UserRequest)
	if err != nil {
		responseFormatter(400, "BAD REQUEST", err.Error(), &response)
		responseWithToken.Response = response
		responseWithToken.Token = ""

		json.NewEncoder(w).Encode(responseWithToken)
		return
	}
	helpers.UserRequestFormatter(UserRequest, &User)
	result, err1 := h.Repo.CreateUser(User)
	if err1 != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR", err1.Error(), &response)
		responseWithToken.Response = response
		responseWithToken.Token = ""

		json.NewEncoder(w).Encode(responseWithToken)
		return
	}
	var user models.UserResponse
	helpers.UserResponseFormatter(result, &user)
	token, err := helpers.GenerateJWT(result.Name, "user")
	responseFormatter(201, "CREATED", user, &response)
	responseWithToken.Response = response
	responseWithToken.Token = token

	json.NewEncoder(w).Encode(responseWithToken)
}

//DeleteUser ...
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {}

//UpdateUser ...
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {}

//FindUser ...
func (h *UserHandler) FindUser(w http.ResponseWriter, r *http.Request) {}

//FindAll Players
func (h *UserHandler) FindAll(w http.ResponseWriter, r *http.Request) {}
