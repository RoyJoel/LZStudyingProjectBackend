package model

import (
	"encoding/json"
)

type Option struct {
	Id        int64   `json:"id"`
	Image     string  `json:"image"`
	Intro     string  `json:"intro"`
	Price     float32 `json:"price"`
	Inventory int64   `json:"inventory"`
	ComId     int64   `json:"comId"`
}

func (Option Option) TableName() string {
	return "Option"
}

func (Option Option) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":        Option.Id,
		"image":     Option.Image,
		"intro":     Option.Intro,
		"comId":     Option.ComId,
		"price":     Option.Price,
		"inventory": Option.Inventory,
	})
}

// Redis类似序列化操作
func (Option Option) MarshalBinary() ([]byte, error) {
	return json.Marshal(Option)
}

func (Option Option) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &Option)
}

type OptionResponse struct {
	Id        int64   `json:"id"`
	Image     string  `json:"image"`
	Intro     string  `json:"intro"`
	Price     float32 `json:"price"`
	Inventory int64   `json:"inventory"`
}

func (OptionResponse OptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":        OptionResponse.Id,
		"image":     OptionResponse.Image,
		"intro":     OptionResponse.Intro,
		"price":     OptionResponse.Price,
		"inventory": OptionResponse.Inventory,
	})
}

// Redis类似序列化操作
func (OptionResponse OptionResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(OptionResponse)
}

func (OptionResponse OptionResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &OptionResponse)
}
