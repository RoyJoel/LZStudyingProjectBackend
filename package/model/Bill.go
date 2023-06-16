package model

import (
	"encoding/json"
)

type Bill struct {
	Id       int64 `json:"id"`
	ComId    int64 `json:"comId"`
	Quantity int64 `json:"quantity"`
	OptionId int64 `json:"optionId"`
	OrderId  int64 `json:"orderId"`
}

func (Bill Bill) TableName() string {
	return "Bill"
}

func (Bill Bill) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":       Bill.Id,
		"comId":    Bill.ComId,
		"quantity": Bill.Quantity,
		"optionId": Bill.OptionId,
		"orderId":  Bill.OrderId,
	})
}

// Redis类似序列化操作
func (Bill Bill) MarshalBinary() ([]byte, error) {
	return json.Marshal(Bill)
}

func (Bill Bill) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &Bill)
}

type BillResponse struct {
	Id       int64             `json:"id"`
	Com      CommodityResponse `json:"com"`
	Quantity int64             `json:"quantity"`
	Option   OptionResponse    `json:"option"`
}

func (BillResponse BillResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":       BillResponse.Id,
		"com":      BillResponse.Com,
		"quantity": BillResponse.Quantity,
		"option":   BillResponse.Option,
	})
}

// Redis类似序列化操作
func (BillResponse BillResponse) MarshalBinary() ([]byte, error) {
	return json.Marshal(BillResponse)
}

func (BillResponse BillResponse) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &BillResponse)
}
