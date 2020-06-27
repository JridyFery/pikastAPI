package helpers

import (
	"bufio"
	"bytes"
	"os"

	"github.com/JridyFery/pikastAPI/models"
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

	pokemonPicture, err := os.Open("assets/pokemons/" + result.PokemonImg) // For read access.
	if err != nil {
		return err
	}

	defer pokemonPicture.Close()

	reader := bufio.NewReader(pokemonPicture)
	buffer := bytes.NewBuffer(make([]byte, 0))

	var chunk []byte
	var eol bool

	for {
		if chunk, eol, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(chunk)
		if !eol {
			buffer.Reset()
		}
	}

	pokemon.PokemonImg = chunk

	return nil
}
