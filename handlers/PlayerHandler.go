package handlers

import (
	"crypto/sha1"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/pikastAR/pikastAPI/helpers"

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

// Login ...
func (h *PlayerHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	var keys []string
	var values []interface{}
	var responseWithToken models.ResponseWithToken
	var response models.Response
	count := 0
	for key, value := range params {
		keys = append(keys, key)
		val, err := strconv.Atoi(value[0])
		if err != nil {
			values = append(values, value[0])
		} else {
			values = append(values, uint(val))
		}
		count++
	}
	result, err := h.Repo.GetPlayerBy(keys, values)
	if err != nil {
		responseFormatter(404, "NOT FOUND", err.Error(), &response)
		responseWithToken.Response = response
		responseWithToken.Token = ""
		json.NewEncoder(w).Encode(responseWithToken)
		return
	}
	if count < 2 {
		responseFormatter(404, "NOT FOUND", err.Error(), &response)
		responseWithToken.Response = response
		responseWithToken.Token = ""
		json.NewEncoder(w).Encode(responseWithToken)
		return
	}
	var player models.PlayerResponse
	helpers.PlayerResponseFormatter(result, &player)
	var role string
	if player.Admin {
		role = "admin"
	} else {
		role = "player"
	}
	token, err := helpers.GenerateJWT(result.PlayerName, role)
	responseFormatter(200, "OK", player, &response)
	responseWithToken.Response = response
	responseWithToken.Token = token
	json.NewEncoder(w).Encode(responseWithToken)
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
		if key == "player_password" {
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

//UpdatePlayerPic func
func (h *PlayerHandler) UpdatePlayerPic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	dt := time.Now().UnixNano()
	playerID, err0 := strconv.Atoi(r.URL.Query()["id"][0])
	if err0 != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR 3", err0.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	r.ParseMultipartForm(10 << 20)
	//upload picture
	file, handler, err := r.FormFile("player_img")
	var fileType string
	if err != nil {
		responseFormatter(400, "BAD REQUEST", err.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	defer file.Close()
	fileType = handler.Header["Content-Type"][0]
	fileType = fileType[6:]
	pictureFile, err3 := ioutil.TempFile("assets/pictures", "pic_*_"+strconv.Itoa(int(dt))+"."+fileType)
	if err3 != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR 1", err3.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	defer pictureFile.Close()
	fileBytes, err4 := ioutil.ReadAll(file)
	if err4 != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR 2", err4.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	pictureFile.Write(fileBytes)
	pictureName := pictureFile.Name()[16:]
	err3 = h.Repo.UpdatePlayerPic(pictureName, uint(playerID))
	if err3 != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR 4", err3.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}

	responseFormatter(200, "OK", "PICTURE UPDATED", &response)
	json.NewEncoder(w).Encode(response)
}

// GetPlayers ...
func (h *PlayerHandler) GetPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	var responseWithCount models.ResponseWithCount
	role := r.URL.Query()["role"][0]
	offset, err0 := strconv.Atoi(r.URL.Query()["offset"][0])
	if err0 != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR", err0.Error(), &response)
		responseWithCount.Response = response
		responseWithCount.Count = 0
		json.NewEncoder(w).Encode(responseWithCount)
		return
	}
	limit, err := strconv.Atoi(r.URL.Query()["limit"][0])
	if err != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR", err.Error(), &response)
		responseWithCount.Response = response
		responseWithCount.Count = 0
		json.NewEncoder(w).Encode(responseWithCount)
		return
	}
	result, count, err1 := h.Repo.GetPlayers(role, offset, limit)
	if err1 != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR", err1.Error(), &response)
		responseWithCount.Response = response
		responseWithCount.Count = 0
		json.NewEncoder(w).Encode(responseWithCount)
		return
	}

	var player models.PlayerResponse
	var players []models.PlayerResponse
	for _, res := range result {
		helpers.PlayerResponseFormatter(res, &player)
		players = append(players, player)
	}
	responseFormatter(200, "OK", players, &response)
	responseWithCount.Response = response
	responseWithCount.Count = count
	json.NewEncoder(w).Encode(responseWithCount)
}

// GetPlayerBy ...
func (h *PlayerHandler) GetPlayerBy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	var keys []string
	var values []interface{}
	var response models.Response
	for key, value := range params {
		keys = append(keys, key)
		val, err := strconv.Atoi(value[0])
		if err != nil {
			values = append(values, value[0])
		} else {
			values = append(values, uint(val))
		}
	}
	result, err := h.Repo.GetPlayerBy(keys, values)
	if err != nil {
		responseFormatter(404, "NOT FOUND", err.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	var player models.PlayerResponse
	helpers.PlayerResponseFormatter(result, &player)
	responseFormatter(200, "OK", player, &response)
	json.NewEncoder(w).Encode(response)
}

//AddPokemonPlayer Association
func (h *PlayerHandler) AddPokemonPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	idPlayer, err := strconv.Atoi(r.URL.Query()["id_player"][0])
	if err != nil {
		responseFormatter(400, "BAD REQUEST", err.Error(), &response)
		return
	}

	idPokemon, err := strconv.Atoi(r.URL.Query()["id_pokemon"][0])
	if err != nil {
		responseFormatter(400, "BAD REQUEST", err.Error(), &response)
		return
	}

	err1 := h.Repo.AddPokemonPlayer(idPlayer, idPokemon)
	if err1 != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR", err1.Error(), &response)

		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(201, "CREATED", "POKEMON ADDED TO PLAYER SUCCESSEFULLY", &response)
	json.NewEncoder(w).Encode(response)
}

//GetPlayerPokemons Association
func (h *PlayerHandler) GetPlayerPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response models.Response
	var responsewithcount models.ResponseWithCount
	idPlayer, err := strconv.Atoi(r.URL.Query()["id_player"][0])
	if err != nil {
		responseFormatter(400, "BAD REQUEST", err.Error(), &response)
		return
	}
	result, count, err := h.Repo.GetplayerPokemons(uint(idPlayer))
	if err != nil {
		responseFormatter(404, "NOT FOUND", err.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200, "OK", result, &response)
	responsewithcount.Response = response
	responsewithcount.Count = count
	json.NewEncoder(w).Encode(responsewithcount)
}
