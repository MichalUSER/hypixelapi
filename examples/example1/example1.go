package main

import (
	"fmt"

	"github.com/MichalUSER/hypixelapi"
)

func main() {
	a := api.Init{Key: "de09bc71-09be-4b8f-bd6f-2dd4a23d85ac"}
	stats := a.GetPlayerStats("gamerboy80")
	player := a.InterfaceToMap(stats["player"])
	playerStats := a.InterfaceToMap(player["stats"])
	bedwars := a.InterfaceToMap(playerStats["Bedwars"])
	guildInfo := a.FindGuild("NOLIFE")
	boosters := a.GetBoosters()
	boosterState := a.InterfaceToMap(boosters["boosterState"])
	friends := a.GetFriends(player["uuid"].(string))
	friendsRecords := a.InterfaceArray(friends["records"])
	firstFriend := a.InterfaceToMap(friendsRecords[0])
	gameCounts := a.GetGameCounts()
	guild := a.GetPlayerGuild(player["uuid"].(string))
	guildMembers := a.InterfaceArray(a.InterfaceToMap(guild["guild"])["members"])
	firstGuildMember := a.InterfaceToMap(guildMembers[0])
	keyInfo := a.GetKey(a.Key)
	keyRecord := a.InterfaceToMap(keyInfo["record"])

	fmt.Println("Total kills:", bedwars["kills_bedwars"])
	fmt.Println("Guild ID:", guildInfo["guild"])
	fmt.Println("First guild member rank:", firstGuildMember["rank"])
	fmt.Println("Booster state:\n    ", "decrementing:", boosterState["decrementing"])
	fmt.Println("First friend id:", firstFriend["_id"])
	fmt.Println("Players online:", gameCounts["playerCount"])
	fmt.Println("Total queries of current key:", keyRecord["totalQueries"])
}
