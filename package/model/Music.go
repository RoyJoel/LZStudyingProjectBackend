package model

import (
	"encoding/json"
)

type Music struct {
	Id          string `json:"id"`
	PlayUrl     string `json:"play_url"`
	Type        string `json:"type"`
	Recommend   int64  `json:"recommend"`
	Atime       int64  `json:"atime"`
	Author      string `json:"author"`
	AnimeInfoId string `json:"anime_info_id"`
}

func (Music Music) TableName() string {
	return "Music"
}

func (Music Music) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":            Music.Id,
		"play_url":      Music.PlayUrl,
		"type":          Music.Type,
		"recommend":     Music.Recommend,
		"atime":         Music.Atime,
		"author":        Music.Author,
		"anime_info_id": Music.AnimeInfoId,
	})
}

// Redis类似序列化操作
func (Music Music) MarshalBinary() ([]byte, error) {
	return json.Marshal(Music)
}

func (Music Music) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &Music)
}

type MusicResponse struct {
	Id        string    `json:"id"`
	PlayUrl   string    `json:"play_url"`
	Type      string    `json:"type"`
	Recommend int64     `json:"recommend"`
	Atime     int64     `json:"atime"`
	Author    string    `json:"author"`
	AnimeInfo AnimeInfo `json:"anime_info"`
}

func (MusicResponse MusicResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":         MusicResponse.Id,
		"play_url":   MusicResponse.PlayUrl,
		"type":       MusicResponse.Type,
		"recommend":  MusicResponse.Recommend,
		"atime":      MusicResponse.Atime,
		"author":     MusicResponse.Author,
		"anime_info": MusicResponse.AnimeInfo,
	})
}

// Redis类似序列化操作
func (MusicResponse MusicResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(MusicResponse)
}

func (MusicResponse MusicResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &MusicResponse)
}
