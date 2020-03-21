package handlers

import (
	"crypto/sha1"
	"encoding/json"
	"net/http"

	"github.com/pikastAR/pikastAPI/helpers"

	"strconv"

	models "github.com/pikastAR/pikastAPI/models"
	"github.com/pikastAR/pikastAPI/repository"
	//"crypto/sha1"
	//"gopkg.in/gomail.v2"
	//"github.com/sethvargo/go-password/password"
)

// PlayerHandler ...
type PlayerHandler struct {
	Repo repository.PlayerRepo
}

func responseFormatter(code int, status string, data interface{}, response *models.Response) {
	response.Code = code
	response.Status = status
	response.Data = data
}

// CreatePlayer ...
func (h *PlayerHandler) CreatePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Player models.Player
	var PlayerRequest models.PlayerRequest
	var response models.Response
	var responseWithToken models.ResponseWithToken
	err := json.NewDecoder(r.Body).Decode(&PlayerRequest)
	if err != nil {
		responseFormatter(400, "BAD REQUEST", err.Error(), &response)
		responseWithToken.Response = response
		responseWithToken.Token = ""

		json.NewEncoder(w).Encode(responseWithToken)
		return
	}

	helpers.PlayerRequestFormatter(PlayerRequest, &Player)
	result, err1 := h.Repo.Createplayer(Player)
	if err1 != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR", err1.Error(), &response)
		responseWithToken.Response = response
		responseWithToken.Token = ""

		json.NewEncoder(w).Encode(responseWithToken)
		return
	}
	var player models.PlayerResponse
	helpers.PlayerResponseFormatter(result, &player)
	token, err := helpers.GenerateJWT(result.PlayerName, "player")
	responseFormatter(201, "CREATED", player, &response)
	responseWithToken.Response = response
	responseWithToken.Token = token

	json.NewEncoder(w).Encode(responseWithToken)
}

//DeletePlayer ...
func (h *PlayerHandler) DeletePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])

	if err != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR", err.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err1 := h.Repo.DeletePlayer(uint(id))
	if err1 != nil {
		responseFormatter(404, "NOT FOUND", err1.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200, "OK", "USER DELETED", &response)
	json.NewEncoder(w).Encode(response)
}

//UpdatePlayer ...
func (h *PlayerHandler) UpdatePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := r.URL.Query()["id"]
	var response models.Response
	id, err1 := strconv.Atoi(params[0])
	if err1 != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR", err1.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	var m map[string]interface{}
	m = make(map[string]interface{})
	var password string
	r.ParseMultipartForm(10 << 20)

	for key, value := range r.Form {
		if key == "password" {
			crypt := sha1.New()
			password = value[0]
			crypt.Write([]byte(password))
			m[key] = crypt.Sum(nil)
		} else {
			if key != "id" {
				if value[0] == "true" {
					m[key] = true
				} else if value[0] == "false" {
					m[key] = false
				} else {
					val, err1 := strconv.Atoi(value[0])
					if err1 != nil || key == "player_tel" {
						m[key] = value[0]
					} else {
						m[key] = val
					}
				}
			}
		}
	}
	err2 := h.Repo.UpdatePlayer(m, uint(id))
	if err2 != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR", err2.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}

	responseFormatter(200, "OK", "PLAYER UPDATED", &response)
	json.NewEncoder(w).Encode(response)
}

//GetPlayer ...
func (h *PlayerHandler) GetPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])
	if err != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR", err.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	result, err1 := h.Repo.GetPlayer(uint(id))
	if err1 != nil {
		responseFormatter(404, "NOT FOUND", err1.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	var player models.PlayerResponse
	helpers.PlayerResponseFormatter(result, &player)
	responseFormatter(200, "OK", player, &response)
	json.NewEncoder(w).Encode(response)
}

//FindAll Players
//func (h *PlayerHandler) FindAll(w http.ResponseWriter, r *http.Request) {}
