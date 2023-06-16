package model

import (
	"encoding/json"
)

type Order struct {
	Id                int64   `json:"id"`
	ShippingAddressId int64   `json:"shippingAddressId"`
	PlayerId          int64   `json:"playerId"`
	CreatedTime       float64 `json:"createdTime"`
	PayedTime         float64 `json:"payedTime"`
	CompletedTime     float64 `json:"completedTime"`
	State             int64   `json:"state"`
	Payment           string  `json:"payment"`
}

func (Order Order) TableName() string {
	return "Order"
}

func (Order Order) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":                Order.Id,
		"playerId":          Order.PlayerId,
		"payment":           Order.Payment,
		"shippingAddressId": Order.ShippingAddressId,
		"createdTime":       Order.CreatedTime,
		"payTime":           Order.PayedTime,
		"completedTime":     Order.CompletedTime,
		"state":             Order.State,
	})
}

// Redis类似序列化操作
func (Order Order) MarshalBinary() ([]byte, error) {
	return json.Marshal(Order)
}

func (Order Order) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &Order)
}

type OrderResponse struct {
	Id              int64           `json:"id"`
	PlayerId        int64           `json:"playerId"`
	Bills           []BillResponse  `json:"bills"`
	ShippingAddress AddressResponse `json:"shippingAddress"`
	CreatedTime     float64         `json:"createdTime"`
	PayedTime       float64         `json:"payedTime"`
	CompletedTime   float64         `json:"completedTime"`
	Payment         string          `json:"payment"`
	State           int64           `json:"state"`
}

func (OrderResponse OrderResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":              OrderResponse.Id,
		"playerId":        OrderResponse.PlayerId,
		"bills":           OrderResponse.Bills,
		"payment":         OrderResponse.Payment,
		"shippingAddress": OrderResponse.ShippingAddress,
		"createdTime":     OrderResponse.CreatedTime,
		"payedTime":       OrderResponse.PayedTime,
		"completedTime":   OrderResponse.CompletedTime,
		"state":           OrderResponse.State,
	})
}

// Redis类似序列化操作
func (OrderResponse OrderResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(OrderResponse)
}

func (OrderResponse OrderResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &OrderResponse)
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
