package model

import (
	"encoding/json"
)

type MusicUser struct {
	PlayerId int64 `json:"playerId"`
	MusicId  int64 `json:"musicId"`
}

func (MusicUser MusicUser) TableName() string {
	return "MusicUser"
}

func (MusicUser MusicUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"playerId": MusicUser.PlayerId,
		"musicId":  MusicUser.MusicId,
	})
}

// Redis类似序列化操作
func (MusicUser MusicUser) MarshalBinary() ([]byte, error) {
	return json.Marshal(MusicUser)
}

func (MusicUser MusicUser) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &MusicUser)
}
