package model

import (
	"encoding/json"
)

type Relationship struct {
	Player1Id       int64 `json:"player1Id"`
	Player2Id int64 `json:"Player2Id"`
}

func (relationship Relationship) TableName() string {
	return "Relationship"
}

func (relationship Relationship) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"player1Id":       relationship.Player1Id,
		"player2Id": relationship.Player2Id,
	})
}

// Redis类似序列化操作
func (relationship Relationship) MarshalBinary() ([]byte, error) {
	return json.Marshal(relationship)
}

func (relationship Relationship) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &relationship)
}
