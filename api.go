package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
)

// Init is a struct which contains player key.
type Init struct {
	Key string
}

// FloatToInt takes float and convert it to int
func (i Init) FloatToInt(number float64) int64 {
	floatNumber := math.Round(number)
	return int64(floatNumber)
}

// InterfaceToMap is a function which converts interface to map using type assertion.
func (i Init) InterfaceToMap(in interface{}) map[string]interface{} {
	return in.(map[string]interface{})
}

// InterfaceArray is a function which returns []interface{}.
func (i Init) InterfaceArray(in interface{}) []interface{} {
	return in.([]interface{})
}

func (i Init) getResponse(url string) []byte {
	resp, err1 := http.Get(url)
	if err1 != nil {
		fmt.Println(err1.Error())
	}
	defer resp.Body.Close()
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	return body
}

// GetPlayerName finds player name by uuid.
func (i Init) GetPlayerName(uuid string) string {
	playerStats := i.InterfaceToMap(i.GetPlayerStatsUUID(uuid)["player"])
	return playerStats["displayname"].(string)
}

// GetPlayerUUID finds player uuid by name.
func (i Init) GetPlayerUUID(name string) string {
	playerStats := i.InterfaceToMap(i.GetPlayerStats(name)["player"])
	return playerStats["uuid"].(string)
}

// GetPlayerStats returns a player's data, such as game stats.
func (i Init) GetPlayerStats(name string) map[string]interface{} {
	body := i.getResponse("https://api.hypixel.net/player?key=" + i.Key + "&name=" + name)
	var f interface{}
	err3 := json.Unmarshal(body, &f)
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	return f.(map[string]interface{})
}

// GetPlayerStatsUUID returns a player's data, such as game stats by uuid.
func (i Init) GetPlayerStatsUUID(uuid string) map[string]interface{} {
	body := i.getResponse("https://api.hypixel.net/player?key=" + i.Key + "&uuid=" + uuid)
	var f interface{}
	err3 := json.Unmarshal(body, &f)
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	return f.(map[string]interface{})
}

// GetBoosters returns list of boosters.
func (i Init) GetBoosters() map[string]interface{} {
	body := i.getResponse("https://api.hypixel.net/boosters?key=" + i.Key)
	var f interface{}
	err3 := json.Unmarshal(body, &f)
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	return f.(map[string]interface{})
}

// FindGuild returns the id of the requested guild.
func (i Init) FindGuild(name string) map[string]interface{} {
	body := i.getResponse("https://api.hypixel.net/findGuild?key=" + i.Key + "&byName=" + name)
	var f interface{}
	err3 := json.Unmarshal(body, &f)
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	return f.(map[string]interface{})
}

// GetFriends returns friendships for given player.
func (i Init) GetFriends(uuid string) map[string]interface{} {
	body := i.getResponse("https://api.hypixel.net/friends?key=" + i.Key + "&uuid=" + uuid)
	var f interface{}
	err3 := json.Unmarshal(body, &f)
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	return f.(map[string]interface{})
}

// GetGameCounts returns the current player count along with the player count of each public game + mode on the server.
func (i Init) GetGameCounts() map[string]interface{} {
	body := i.getResponse("https://api.hypixel.net/gameCounts?key=" + i.Key)
	var f interface{}
	err3 := json.Unmarshal(body, &f)
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	return f.(map[string]interface{})
}

// GetGuild returns information about given guild name.
func (i Init) GetGuild(name string) map[string]interface{} {
	body := i.getResponse("https://api.hypixel.net/guild?key=" + i.Key + "&name=" + name)
	var f interface{}
	err3 := json.Unmarshal(body, &f)
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	return f.(map[string]interface{})
}

// GetPlayerGuild returns information about given player uuid guild.
func (i Init) GetPlayerGuild(uuid string) map[string]interface{} {
	body := i.getResponse("https://api.hypixel.net/guild?key=" + i.Key + "&player=" + uuid)
	var f interface{}
	err3 := json.Unmarshal(body, &f)
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	return f.(map[string]interface{})
}

// GetKey returns information regarding given key.
func (i Init) GetKey(key string) map[string]interface{} {
	body := i.getResponse("https://api.hypixel.net/key?key=" + key)
	var f interface{}
	err3 := json.Unmarshal(body, &f)
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	return f.(map[string]interface{})
}

// GetLeaderboards returns a list of the official leaderboards and their current standings for games on the network.
func (i Init) GetLeaderboards() map[string]interface{} {
	body := i.getResponse("https://api.hypixel.net/leaderboards?key=" + i.Key)
	var f interface{}
	err3 := json.Unmarshal(body, &f)
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	return f.(map[string]interface{})
}

// GetPlayerCount returns current player count.
func (i Init) GetPlayerCount() float64 {
	gameCounts := i.GetGameCounts()
	return gameCounts["playerCount"].(float64)
}

// GetRecentGames returns recent games of a player. A maximum of 100 games are returned and recent games are only stored for up to 3 days at this time.
func (i Init) GetRecentGames(uuid string) map[string]interface{} {
	body := i.getResponse("https://api.hypixel.net/recentGames?key=" + i.Key + "&uuid=" + uuid)
	var f interface{}
	err3 := json.Unmarshal(body, &f)
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	return f.(map[string]interface{})
}

// GetResources provides an endpoint to retrieve resources which don't change often. This does not require an API key.
func (i Init) GetResources(resource string) map[string]interface{} {
	body := i.getResponse("https://api.hypixel.net/resources/" + resource)
	var f interface{}
	err3 := json.Unmarshal(body, &f)
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	return f.(map[string]interface{})
}

// GetStatus returns online status information for given player, including game, mode and map when available.
func (i Init) GetStatus(uuid string) map[string]interface{} {
	body := i.getResponse("https://api.hypixel.net/status?key=" + i.Key + "&uuid=" + uuid)
	var f interface{}
	err3 := json.Unmarshal(body, &f)
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	return f.(map[string]interface{})
}

// GetWatchdogStats returns some statistics about Watchdog & bans.
func (i Init) GetWatchdogStats() map[string]interface{} {
	body := i.getResponse("https://api.hypixel.net/watchdogstats?key=" + i.Key)
	var f interface{}
	err3 := json.Unmarshal(body, &f)
	if err3 != nil {
		fmt.Println(err3.Error())
	}
	return f.(map[string]interface{})
}
