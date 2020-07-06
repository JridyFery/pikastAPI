package helpers

import (
	"bufio"
<<<<<<< HEAD
	//"bytes"
	"os"
=======
>>>>>>> 114014fb23a130de2be1f541bd6f5b86b431c935

	"github.com/JridyFery/pikastAPI/models"
//	"github.com/jinzhu/gorm"
	//"encoding/hex"
	//"os"
)

// PokemonResponseFormatter func
func PokemonResponseFormatter(result models.Pokemon, pokemon *models.PokemonResponse) error {
	pokemon.Model=result.Model
	pokemon.PokemonName = result.PokemonName
	pokemon.PokemonPrefab = result.PokemonPrefab
	pokemon.PokemonFBX = result.PokemonFBX
	pokemon.PokemonisPremium = result.PokemonisPremium
	pokemon.PokemonCost = result.PokemonCost
	pokemon.WithDiamonds = result.WithDiamonds
	pokemon.PokemonMaxPower=result.PokemonMaxPower
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
	//myString := string(bytes)
	pokemon.PokemonImg=bytes
	return nil
}
