package model

import (
	"encoding/json"

	"github.com/RoyJoel/LZStudyingProjectBackend/package/utils"
)

type User struct {
	Id             int64            `json:"id"`
	LoginName      string           `json:"loginName"`
	Password       string           `json:"password"`
	Name           string           `json:"name"`
	Icon           string           `json:"icon"`
	Sex            string           `json:"sex"`
	Age            int64            `json:"age"`
	Points         int64            `json:"points"`
	Friends        []PlayerResponse `json:"friends"`
	AllLikedMusic  utils.IntMatrix  `json:"allLikedMusic"`
	Addresss       utils.IntMatrix  `json:"addresss"`
	AllOrders      utils.IntMatrix  `json:"allOrders"`
	Cart           int64            `json:"cart"`
	DefaultAddress AddressResponse  `json:"defaultAddress"`
	Token          string           `json:"token"`
}

func (User User) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]interface{}{
		"id":        User.Id,
		"loginName": User.LoginName,
		"password":  User.Password,
		"name":      User.Name,
		"icon":      User.Icon,
		"sex":       User.Sex,
		"age":       User.Age,

		"points": User.Points,

		"friends":        User.Friends,
		"allLikedMusic":  User.AllLikedMusic,
		"addresss":       User.Addresss,
		"allOrders":      User.AllOrders,
		"cart":           User.Cart,
		"defaultAddress": User.DefaultAddress,
		"token":          User.Token,
	})
}

// Redis类似序列化操作
func (User User) MarshalBinary() ([]byte, error) {
	return json.Marshal(User)
}

func (User User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, &User)
}
