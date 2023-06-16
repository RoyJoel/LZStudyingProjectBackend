package model

import (
	"encoding/json"
)

type Player struct {
	Id        int64  `json:"id"`
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Sex       string `json:"sex"`
	Age       int64  `json:"age"`
	Points    int64  `json:"points"`
}

func (player Player) TableName() string {
	return "Player"
}

func (player Player) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":        player.Id,
		"loginName": player.LoginName,
		"password":  player.Password,
		"name":      player.Name,
		"icon":      player.Icon,
		"sex":       player.Sex,
		"age":       player.Age,
		"points":    player.Points,
	})
}

// Redis类似序列化操作
func (player Player) MarshalBinary() ([]byte, error) {
	return json.Marshal(player)
}

func (player Player) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &player)
}

type PlayerResponse struct {
	Id        int64  `json:"id"`
	LoginName string `json:"loginName"`
	Password  string `json:"password"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	Sex       string `json:"sex"`
	Age       int64  `json:"age"`
	Points    int64  `json:"points"`
}

func (playerResponse PlayerResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":        playerResponse.Id,
		"loginName": playerResponse.LoginName,
		"password":  playerResponse.Password,
		"name":      playerResponse.Name,
		"icon":      playerResponse.Icon,
		"sex":       playerResponse.Sex,
		"age":       playerResponse.Age,
		"points":    playerResponse.Points,
	})
}

// Redis类似序列化操作
func (playerResponse PlayerResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(playerResponse)
}

func (playerResponse PlayerResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &playerResponse)
}
