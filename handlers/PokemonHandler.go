package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/JridyFery/pikastAPI/helpers"
	models "github.com/JridyFery/pikastAPI/models"
	repository "github.com/JridyFery/pikastAPI/repository"
)

// PokemonHandler ...
type PokemonHandler struct {
	Repo repository.PokemonRepo
}

// CreatePokemon ...
func (h *PokemonHandler) CreatePokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Pokemon models.Pokemon
	var response models.Response
	err := json.NewDecoder(r.Body).Decode(&Pokemon)
	if err != nil {
		responseFormatter(400, "BAD REQUEST", err.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	result, err1 := h.Repo.CreatePokemon(Pokemon)
	if err1 != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR", err1.Error(), &response)

		json.NewEncoder(w).Encode(response)
		return
	}
	//Check if it's really necessary
	responseFormatter(201, "CREATED", result, &response)
	json.NewEncoder(w).Encode(response)
}

//GetPokemon ...
func (h *PokemonHandler) GetPokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	var PokemonResponse models.PokemonResponse
	id, err := strconv.Atoi(params[0])
	if err != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR", err.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	result, err1 := h.Repo.GetPokemon(uint(id))
	if err1 != nil {
		responseFormatter(404, "NOT FOUND", err1.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	helpers.PokemonResponseFormatter(result, &PokemonResponse)
	responseFormatter(200, "OK", PokemonResponse, &response)
	json.NewEncoder(w).Encode(response)

}

//DeletePokemon ...
func (h *PokemonHandler) DeletePokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
	id, err := strconv.Atoi(params[0])

	if err != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR", err.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	err1 := h.Repo.DeletePokemon(uint(id))
	if err1 != nil {
		responseFormatter(404, "NOT FOUND", err1.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200, "OK", "POKEMON DELETED", &response)
	json.NewEncoder(w).Encode(response)
}

//GetPokemons func
func (h *PokemonHandler) GetPokemons(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var pokemonResponse models.PokemonResponse
	var pokemonResponses []models.PokemonResponse
	var response models.Response
	var responseWithCount models.ResponseWithCount
	pokemonType := r.URL.Query()["pokemonType"][0]
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
	result, count, err1 := h.Repo.GetPokemons(pokemonType, offset, limit)
	if err1 != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR", err1.Error(), &response)
		responseWithCount.Response = response
		responseWithCount.Count = 0
		json.NewEncoder(w).Encode(responseWithCount)
		return
	}
	for _, pok := range result {
		helpers.PokemonResponseFormatter(pok, &pokemonResponse)
		pokemonResponses = append(pokemonResponses, pokemonResponse)
	}
	responseFormatter(200, "OK", pokemonResponses, &response)
	responseWithCount.Response = response
	responseWithCount.Count = count
	json.NewEncoder(w).Encode(responseWithCount)
}

//UpdatePokemon  ...
func (h *PokemonHandler) UpdatePokemon(w http.ResponseWriter, r *http.Request) {
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
	r.ParseMultipartForm(10 << 20)

	for key, value := range r.Form {

		if key != "id" {
			if value[0] == "true" {
				m[key] = true
			} else if value[0] == "false" {
				m[key] = false
			} else {
				val, err1 := strconv.Atoi(value[0])
				if err1 != nil {
					m[key] = value[0]
				} else {
					m[key] = val
				}
			}
		}
	}
	err2 := h.Repo.UpdatePokemon(m, uint(id))
	if err2 != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR", err2.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}

	responseFormatter(200, "OK", "PLAYER UPDATED", &response)
	json.NewEncoder(w).Encode(response)
}

//UpdatePokemonPic func
func (h *PokemonHandler) UpdatePokemonPic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json ")
	var response models.Response
	var requestImage models.PokemonRequestImage
	dt := time.Now().UnixNano()
	err := json.NewDecoder(r.Body).Decode(&requestImage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err != nil {
		responseFormatter(400, "BAD REQUEST", err.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	//defer file.Close()
	pictureFile, err3 := ioutil.TempFile("assets/pictures/pokemons", "pic_*_"+strconv.Itoa(int(dt))+".png")
	if err3 != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR 1", err3.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	defer pictureFile.Close()

	pictureFile.Write(requestImage.PokemonImg)
	pictureName := pictureFile.Name()[25:]
	err3 = h.Repo.UpdatePokemonPic(pictureName, uint(requestImage.PokemonID))
	if err3 != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR 4", err3.Error(), &response)
		json.NewEncoder(w).Encode(response)
		return
	}
	responseFormatter(200, "OK", "PICTURE UPDATED", &response)
	json.NewEncoder(w).Encode(response)
}
