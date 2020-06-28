package helpers

import (
	"bufio"

	"github.com/JridyFery/pikastAPI/models"
	"os"
)

// PokemonResponseFormatter func
func PokemonResponseFormatter(result models.Pokemon, pokemon *models.PokemonResponse) error {
	pokemon.PokemonName = result.PokemonName
	pokemon.PokemonPrefab = result.PokemonPrefab
	pokemon.PokemonFBX = result.PokemonFBX
	pokemon.PokemonisPremium = result.PokemonisPremium
	pokemon.PokemonCost = result.PokemonCost
	pokemon.WithDiamonds = result.WithDiamonds
	pokemon.PokemonPower = result.PokemonPower
	pokemon.PokemonAttacktype = result.PokemonAttacktype
	pokemon.PokemonATKSpeed = result.PokemonATKSpeed
	pokemon.PokemonMOVSpeed = result.PokemonMOVSpeed
	pokemon.PokemonHeight = result.PokemonHeight
	pokemon.PokemonWidth = result.PokemonWidth

	pokemonPicture, err := os.Open("assets/pictures/pokemons/" + result.PokemonImg) // For read access.
	if err != nil {
		return err
	}

	defer pokemonPicture.Close()

	fileInfo, _ := pokemonPicture.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	// read file into bytes
	buffer := bufio.NewReader(pokemonPicture)
	_, err = buffer.Read(bytes)
	pokemon.PokemonImg=bytes
	return nil
}
