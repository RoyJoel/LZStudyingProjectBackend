package model

import (
	"encoding/json"
)

type Cart struct {
	PlayerId int64 `json:"playerId"`
	OrderId  int64 `json:"orderId"`
}

func (Cart Cart) TableName() string {
	return "Cart"
}

func (Cart Cart) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"playerId": Cart.PlayerId,
		"orderId":  Cart.OrderId,
	})
}

// Redis类似序列化操作
func (Cart Cart) MarshalBinary() ([]byte, error) {
	return json.Marshal(Cart)
}

func (Cart Cart) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &Cart)
}
