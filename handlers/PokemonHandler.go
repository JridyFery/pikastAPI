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
