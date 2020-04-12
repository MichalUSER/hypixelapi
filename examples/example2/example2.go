package main

import (
	"fmt"

	"github.com/MichalUSER/hypixelapi"
)

func main() {
	a := api.Init{Key: "your-key"}
	name := "gamerboy80"
	leaderboardsJSON := a.GetLeaderboards()
	leaderboards := a.InterfaceToMap(leaderboardsJSON["leaderboards"])
	leaderboardsBedwarsArray := a.InterfaceArray(leaderboards["BEDWARS"])
	leaderboardsBedwarsLevel := a.InterfaceToMap(leaderboardsBedwarsArray[0])
	leaderboardsBedwarsLeaders := a.InterfaceArray(leaderboardsBedwarsLevel["leaders"])
	bedwarsLeaderName := a.GetPlayerName(leaderboardsBedwarsLeaders[0].(string))
	playerCount := a.GetPlayerCount()
	recentGamesJSON := a.GetRecentGames(a.GetPlayerUUID(name))
	recentGamesArray := a.InterfaceArray(recentGamesJSON["games"])
	mostRecentGame := a.InterfaceToMap(recentGamesArray[0])
	playerStatusJSON := a.GetStatus(a.GetPlayerUUID(name))
	playerStatus := a.InterfaceToMap(playerStatusJSON["session"])
	watchdogStats := a.GetWatchdogStats()
	totalBannedPlayers := a.FloatToInt(watchdogStats["watchdog_total"].(float64))

	fmt.Println("#1 Bedwars Player:", bedwarsLeaderName)
	fmt.Println("Online players:", playerCount)
	fmt.Println("Last recent game:", mostRecentGame["gameType"], mostRecentGame["mode"])
	fmt.Println("Player", name, "is online:", playerStatus["online"])
	fmt.Println("Watchdog has banned exactly", totalBannedPlayers, "players")
}
