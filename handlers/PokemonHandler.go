package handlers

import (
	"net/http"

	"encoding/json"
	"strconv"

	helpers "github.com/pikastAR/pikastAPI/helpers"
	models "github.com/pikastAR/pikastAPI/models"
	repository "github.com/pikastAR/pikastAPI/repository"
)

// PokemonHandler ...
type PokemonHandler struct {
	Repo repository.PokemonRepo
}

// CreatePokemon ...
func (h *PokemonHandler) CreatePokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Pokemon models.Pokemon
	var PokemonRequest models.PokemonRequest
	var response models.Response
	var responseWithToken models.ResponseWithToken
	err := json.NewDecoder(r.Body).Decode(&PokemonRequest)
	if err != nil {
		responseFormatter(400, "BAD REQUEST", err.Error(), &response)
		responseWithToken.Response = response
		responseWithToken.Token = ""

		json.NewEncoder(w).Encode(responseWithToken)
		return
	}

	helpers.PokemonRequestFormatter(PokemonRequest, &Pokemon)
	result, err1 := h.Repo.CreatePokemon(Pokemon)
	if err1 != nil {
		responseFormatter(500, "INTERNAL SERVER ERROR", err1.Error(), &response)
		responseWithToken.Response = response
		responseWithToken.Token = ""

		json.NewEncoder(w).Encode(responseWithToken)
		return
	}
	var pokemon models.PokemonResponse
	helpers.PokemonResponseFormatter(result, &pokemon)
	//Check if it's really necessary
	token, err := helpers.GenerateJWT(result.PokemonName, "pokemon")
	responseFormatter(201, "CREATED", pokemon, &response)
	responseWithToken.Response = response
	responseWithToken.Token = token

	json.NewEncoder(w).Encode(responseWithToken)
}

//GetPokemon ...
func (h *PokemonHandler) GetPokemon(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()["id"]
	var response models.Response
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
	var pokemon models.PokemonResponse
	helpers.PokemonResponseFormatter(result, &pokemon)
	responseFormatter(200, "OK", pokemon, &response)
	json.NewEncoder(w).Encode(response)
}
