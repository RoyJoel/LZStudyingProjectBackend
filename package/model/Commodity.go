package model

import (
	"encoding/json"
)

type Commodity struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Intro string `json:"intro"`
	Cag   int64  `json:"cag"`
	State int64  `json:"state"`
}

func (Commodity Commodity) TableName() string {
	return "Commodity"
}

func (Commodity Commodity) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":    Commodity.Id,
		"name":  Commodity.Name,
		"intro": Commodity.Intro,
		"cag":   Commodity.Cag,
		"state": Commodity.State,
	})
}

// Redis类似序列化操作
func (Commodity Commodity) MarshalBinary() ([]byte, error) {
	return json.Marshal(Commodity)
}

func (Commodity Commodity) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &Commodity)
}

type CommodityResponse struct {
	Id      int64            `json:"id"`
	Name    string           `json:"name"`
	Intro   string           `json:"intro"`
	Cag     int64            `json:"cag"`
	Orders  int64            `json:"orders"`
	Options []OptionResponse `json:"options"`
	State   int64            `json:"state"`
}

func (CommodityResponse CommodityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":      CommodityResponse.Id,
		"name":    CommodityResponse.Name,
		"intro":   CommodityResponse.Intro,
		"cag":     CommodityResponse.Cag,
		"orders":  CommodityResponse.Orders,
		"options": CommodityResponse.Options,
		"state":   CommodityResponse.State,
	})
}

// Redis类似序列化操作
func (CommodityResponse CommodityResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(CommodityResponse)
}

func (CommodityResponse CommodityResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &CommodityResponse)
}
