package model

import (
	"encoding/json"
)

type Address struct {
	Id              int64  `json:"id"`
	PlayerId        int64  `json:"playerId"`
	Name            string `json:"name"`
	Sex             string `json:"sex"`
	PhoneNumber     string `json:"phoneNumber"`
	Province        string `json:"province"`
	City            string `json:"city"`
	Area            string `json:"area"`
	DetailedAddress string `json:"detailedAddress"`
	IsDefault       bool   `json:"isDefault"`
}

func (Address Address) TableName() string {
	return "Address"
}

func (Address Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":              Address.Id,
		"playerId":        Address.PlayerId,
		"name":            Address.Name,
		"sex":             Address.Sex,
		"phoneNumber":     Address.PhoneNumber,
		"province":        Address.Province,
		"city":            Address.City,
		"area":            Address.Area,
		"detailedAddress": Address.DetailedAddress,
		"isDefault":       Address.IsDefault,
	})
}

// Redis类似序列化操作
func (Address Address) MarshalBinary() ([]byte, error) {
	return json.Marshal(Address)
}

func (Address Address) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &Address)
}

type AddressResponse struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	Sex           string `json:"sex"`
	PhoneNumber   string `json:"phoneNumber"`
	Province      string `json:"province"`
	City          string `json:"city"`
	Area          string `json:"area"`
	DetailAddress string `json:"detailedAddress"`
	IsDefault     bool   `json:"isDefault"`
}

func (AddressResponse AddressResponse) TableName() string {
	return "Address"
}

func (AddressResponse AddressResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":              AddressResponse.Id,
		"name":            AddressResponse.Name,
		"sex":             AddressResponse.Sex,
		"phoneNumber":     AddressResponse.PhoneNumber,
		"province":        AddressResponse.Province,
		"city":            AddressResponse.City,
		"area":            AddressResponse.Area,
		"detailedAddress": AddressResponse.DetailAddress,
		"isDefault":       AddressResponse.IsDefault,
	})
}

// Redis类似序列化操作
func (AddressResponse AddressResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(AddressResponse)
}

func (AddressResponse AddressResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &AddressResponse)
}

// type AddressResponse struct {
// 	Id      int64           `json:"id"`
// 	Player pl `json:"playerId"`
// 	Name    string          `json:"name"`
// 	Sex    string          `json:"sex"`
// 	PhoneNumber   string          `json:"phoneNumber"`
// 	Province string           `json:"province"`
// 	City string          `json:"city"`
// 	Area  string  `json:"area"`
// 	DetailAddress string `json:"detailedAddress"`
// }

// func (AddressResponse AddressResponse) MarshalJSON() ([]byte, error) {
// 	return json.Marshal(map[string]interface{}{
// 		"id":      AddressResponse.Id,
// 		"playerId": AddressResponse.PlayerId,
// 		"name":    AddressResponse.Name,
// 		"sex":    AddressResponse.Sex,
// 		"phoneNumber":   AddressResponse.PhoneNumber,
// 		"city": AddressResponse.City,
// 		"area": AddressResponse.Area,
// 		"detailedAddress": AddressResponse.DetailAddress,
// 	})
// }

// // Redis类似序列化操作
// func (AddressResponse AddressResponse) MarshalBinary() ([]byte, error) {
// 	return json.Marshal(AddressResponse)
// }

// func (AddressResponse AddressResponse) UnmarshalBinary(data []byte) error {
// 	return json.Unmarshal(data, &AddressResponse)
// }
