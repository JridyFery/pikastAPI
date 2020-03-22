package helpers

import "github.com/pikastAR/pikastAPI/models"

//PokemonRequestFormatter func
func PokemonRequestFormatter(request models.PokemonRequest, pokemon *models.Pokemon) {
	pokemon.PokemonName = request.PokemonName
	pokemon.PokemonPrefab = request.PokemonPrefab
	pokemon.PokemonImg = request.PokemonImg
	pokemon.PokemonisPremium = request.PokemonisPremium
	pokemon.PokemonPoints = request.PokemonPoints
	pokemon.PokemonPower = request.PokemonPower
	pokemon.PokemonAttacktype = request.PokemonAttacktype
	pokemon.PokemonATKSpeed = request.PokemonATKSpeed
	pokemon.PokemonMOVSpeed = request.PokemonMOVSpeed
	pokemon.PokemonHeight = request.PokemonHeight
	pokemon.PokemonWidth = request.PokemonWidth

}

//PokemonResponseFormatter func
func PokemonResponseFormatter(result models.Pokemon, pokemon *models.PokemonResponse) {
	pokemon.PokemonID = result.ID
	pokemon.PokemonName = result.PokemonName
	pokemon.PokemonPrefab = result.PokemonPrefab
	pokemon.PokemonImg = result.PokemonImg
	pokemon.PokemonisPremium = result.PokemonisPremium
	pokemon.PokemonPoints = result.PokemonPoints
	pokemon.PokemonPower = result.PokemonPower
	pokemon.PokemonAttacktype = result.PokemonAttacktype
	pokemon.PokemonATKSpeed = result.PokemonATKSpeed
	pokemon.PokemonATKSpeed = result.PokemonATKSpeed
	pokemon.PokemonMOVSpeed = result.PokemonMOVSpeed
	pokemon.PokemonHeight = result.PokemonHeight
	pokemon.PokemonWidth = result.PokemonWidth
}
