package model

import (
	"encoding/json"
)

type AnimeInfo struct {
	Id    string `json:"id"`
	Bg    string `json:"bg"`
	Year  int64  `json:"year"`
	Month int64  `json:"month"`
	Title string `json:"title"`
	Atime string `json:"atime"`
	Desc  string `json:"desc"`
	Logo  string `json:"logo"`
}

func (AnimeInfo AnimeInfo) TableName() string {
	return "AnimeInfo"
}

func (AnimeInfo AnimeInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":    AnimeInfo.Id,
		"bg":    AnimeInfo.Bg,
		"year":  AnimeInfo.Year,
		"month": AnimeInfo.Month,
		"title": AnimeInfo.Title,
		"atime": AnimeInfo.Atime,
		"desc":  AnimeInfo.Desc,
		"logo":  AnimeInfo.Logo,
	})
}

// Redis类似序列化操作
func (AnimeInfo AnimeInfo) MarshalBinary() ([]byte, error) {
	return json.Marshal(AnimeInfo)
}

func (AnimeInfo AnimeInfo) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &AnimeInfo)
}
