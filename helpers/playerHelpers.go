package helpers

import (
	"bufio"
	"crypto/sha1"
	"os"

	models "github.com/JridyFery/pikastAPI/models"
)

// PlayerResponseFormatter func
func PlayerResponseFormatter(result models.Player, player *models.PlayerResponse) error {
	player.PlayerID = result.ID
	player.PlayerName = result.PlayerName
	player.PlayerEmail = result.PlayerEmail
	player.Admin = result.Admin
	player.PlayerTel = result.PlayerTel
	player.DateOfBirth.Day = result.BirthDay
	player.DateOfBirth.Month = result.BirthMonth
	player.DateOfBirth.Year = result.BirthYear
	player.Country = result.Country
	player.PlayerRank = result.PlayerRank
	player.PlayerCoins = result.PlayerCoins
	player.PlayerLevelCount = result.PlayerLevelCount
	player.PlayerLevelProgress = result.PlayerLevelProgress

	playerPicture, err := os.Open("assets/pictures/players/" + result.PlayerImg) // For read access.
	if err != nil {
		return err
	}

	defer playerPicture.Close()

	fileInfo, _ := playerPicture.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	// read file into bytes
	buffer := bufio.NewReader(playerPicture)
	_, err = buffer.Read(bytes)
	//myString := string(bytes)
	player.PlayerImg = bytes
	return nil
}

// PlayerRequestFormatter func
func PlayerRequestFormatter(request models.PlayerRequest, player *models.Player) {
	if request.PlayerPassword != "" {
		crypt := sha1.New()
		crypt.Write([]byte(request.PlayerPassword))
		player.PlayerPassword = crypt.Sum(nil)
	}
	player.PlayerName = request.PlayerName
	player.PlayerEmail = request.PlayerEmail
	player.Admin = request.Admin
	player.PlayerTel = request.PlayerTel
	player.BirthDay = request.BirthDay
	player.BirthMonth = request.BirthMonth
	player.BirthYear = request.BirthYear
	player.Country = request.Country
	player.PlayerImg = request.PlayerImg
	player.PlayerRank = request.PlayerRank
	player.PlayerCoins = request.PlayerCoins
	player.PlayerLevelCount = request.PlayerLevelCount
	player.PlayerLevelProgress = request.PlayerLevelProgress
}
