package helpers

import (
	"crypto/sha1"

	models "github.com/JridyFery/pikastAPI/models"
)

// PlayerResponseFormatter func
func PlayerResponseFormatter(result models.Player, player *models.PlayerResponse) {
	player.PlayerID = result.ID
	player.PlayerName = result.PlayerName
	player.PlayerEmail = result.PlayerEmail
	player.Admin = result.Admin
	player.PlayerTel = result.PlayerTel
	player.DateOfBirth.Day = result.BirthDay
	player.DateOfBirth.Month = result.BirthMonth
	player.DateOfBirth.Year = result.BirthYear
	player.Country = result.Country
	player.PlayerImg = result.PlayerImg
	player.PlayerRank = result.PlayerRank
	player.PlayerCoins = result.PlayerCoins
	player.PlayerLevelCount = result.PlayerLevelCount
	player.PlayerLevelProgress = result.PlayerLevelProgress

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
