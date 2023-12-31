package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/RoyJoel/LZStudyingProjectBackend/package/dao/impl"
	"github.com/RoyJoel/LZStudyingProjectBackend/package/middleware"
	"github.com/RoyJoel/LZStudyingProjectBackend/package/model"
	"github.com/RoyJoel/LZStudyingProjectBackend/package/utils"
	"github.com/gin-gonic/gin"
)

type LZControllerImpl struct {
	dao *impl.LZDaoImpl
}

func NewLZControllerImpl() *LZControllerImpl {
	return &LZControllerImpl{dao: impl.NewLZDaoImpl()}
}

func (impl LZControllerImpl) AddPlayer(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	player := model.Player{}
	json.Unmarshal(bytes, &player)
	if err != nil {
		panic(err)
	}
	// player.Role = 1
	res, _ := impl.dao.AddPlayer(c, player)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) SignUp(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	User := model.User{}
	json.Unmarshal(bytes, &User)
	if err != nil {
		panic(err)
	}
	// player.Role = 1
	user, res := impl.dao.SignUp(c, User)
	type addResponse struct {
		User model.User `json:"user"`
		Res  bool       `json:"res"`
	}
	result := addResponse{User: user, Res: res}
	fmt.Println(user)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": result})
}

func (impl LZControllerImpl) UpdateUser(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	User := model.User{}
	json.Unmarshal(bytes, &User)
	if err != nil {
		panic(err)
	}
	// player.Role = 1
	user := impl.dao.UpdateUser(c, User)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": user})
}

func (impl LZControllerImpl) Auth(c *gin.Context) {
	auth := c.Request.Header.Get("Authorization")
	claims, error := middleware.ParseToken(auth)
	type authResponse struct {
		LoginName string `json:"loginName"`
		Password  string `json:"password"`
	}
	res := authResponse{LoginName: claims.LoginName, Password: claims.Password}

	if error != nil {
		c.JSON(401, map[string]interface{}{"code": 0, "msg": error, "count": 0, "data": nil})
	} else {
		c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
	}
}

func (impl LZControllerImpl) SignIn(c *gin.Context) {

	type SignInRequest struct {
		LoginName string `json:"loginName"`
		Password  string `json:"password"`
	}

	var req SignInRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	user, error := impl.dao.SignIn(c, req.LoginName, req.Password)
	if error != nil {
		c.JSON(401, map[string]interface{}{"code": 0, "msg": error, "count": 0, "data": nil})
	}
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": user})

}

func (impl LZControllerImpl) ResetPassword(c *gin.Context) {

	type resetRequest struct {
		LoginName string `json:"loginName"`
		Password  string `json:"password"`
	}

	var req resetRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	res := impl.dao.ResetPassword(c, req.LoginName, req.Password)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})

}

func (impl LZControllerImpl) SearchPlayer(c *gin.Context) {

	type SearchPlayerRequest struct {
		LoginName string `json:"loginName"`
	}

	var req SearchPlayerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	player := impl.dao.SearchPlayer(c, req.LoginName)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": player})

}

func (impl LZControllerImpl) UpdatePlayer(c *gin.Context) {

	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	player := model.PlayerResponse{}
	json.Unmarshal(bytes, &player)
	if err != nil {
		panic(err)
	}

	player = impl.dao.UpdatePlayer(c, player)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": player})
}

func (impl LZControllerImpl) GetPlayerInfo(c *gin.Context) {
	type SearchPlayerRequest struct {
		LoginName string `json:"loginName"`
	}

	var req SearchPlayerRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	player, _ := impl.dao.GetPlayerInfoByLoginName(c, req.LoginName)
	if player.LoginName == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get player info"})
	} else {
		c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": player})
	}
}

func (impl LZControllerImpl) AddFriend(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	relationship := model.Relationship{}
	json.Unmarshal(bytes, &relationship)
	if err != nil {
		panic(err)
	}
	result := impl.dao.AddFriend(c, relationship)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": result})
}

func (impl LZControllerImpl) DeleteFriend(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	relationship := model.Relationship{}
	json.Unmarshal(bytes, &relationship)
	if err != nil {
		panic(err)
	}
	result := impl.dao.DeleteFriend(c, relationship)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": result})
}

func (impl LZControllerImpl) GetAllFriends(c *gin.Context) {
	type SearchPlayerRequest struct {
		Id int64 `json:"id"`
	}

	var req SearchPlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	player := impl.dao.GetAllFriends(c, req.Id)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": player})
}

func (impl LZControllerImpl) SearchFriend(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	relationship := model.Relationship{}
	json.Unmarshal(bytes, &relationship)
	if err != nil {
		panic(err)
	}

	res := impl.dao.SearchFriend(c, relationship)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) GetMusicInfos(c *gin.Context) {
	type SearchPlayerRequest struct {
		Ids utils.IntMatrix `json:"ids"`
	}

	var req SearchPlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}
	res := make([]model.MusicResponse, 0)

	for _, id := range req.Ids {
		club := impl.dao.GetMusicInfo(c, id)
		res = append(res, club)
	}
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) AddOrder(c *gin.Context) {

	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	order := model.OrderResponse{}
	json.Unmarshal(bytes, &order)
	if err != nil {
		panic(err)
	}

	res := impl.dao.AddOrder(c, order)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) UpdateOrder(c *gin.Context) {

	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	order := model.OrderResponse{}
	json.Unmarshal(bytes, &order)
	if err != nil {
		panic(err)
	}

	res := impl.dao.UpdateOrder(c, order)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) DeleteOrder(c *gin.Context) {
	type DeleteOrderRequest struct {
		Id int64 `json:"id"`
	}

	var req DeleteOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	res := impl.dao.DeleteOrder(c, req.Id)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) AddAddress(c *gin.Context) {

	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	address := model.Address{}
	json.Unmarshal(bytes, &address)
	if err != nil {
		panic(err)
	}

	res := impl.dao.AddAddress(c, address)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) UpdateAddress(c *gin.Context) {

	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	address := model.Address{}
	json.Unmarshal(bytes, &address)
	if err != nil {
		panic(err)
	}

	res := impl.dao.UpdateAddress(c, address)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) GetAddressInfos(c *gin.Context) {
	type SearchPlayerRequest struct {
		Ids utils.IntMatrix `json:"ids"`
	}

	var req SearchPlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}
	res := make([]model.AddressResponse, 0)

	for _, id := range req.Ids {
		address := impl.dao.GetAddressInfo(c, id)
		res = append(res, address)
		fmt.Println(res)
	}
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) DeleteAddress(c *gin.Context) {
	type DeleteAddressRequest struct {
		Id int64 `json:"id"`
	}

	var req DeleteAddressRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	res := impl.dao.DeleteAddress(c, req.Id)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) GetOrderInfosByUserId(c *gin.Context) {
	type SearchPlayerRequest struct {
		Id int64 `json:"id"`
	}

	var req SearchPlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	res := impl.dao.GetOrderInfosByUserId(c, req.Id)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) GetCartInfo(c *gin.Context) {
	type SearchPlayerRequest struct {
		Id int64 `json:"id"`
	}

	var req SearchPlayerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	res := impl.dao.GetOrderInfo(c, req.Id)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) AddBillToCart(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	bill := model.Bill{}
	json.Unmarshal(bytes, &bill)
	if err != nil {
		panic(err)
	}
	// Game.Role = 1
	res := impl.dao.AddBillToCart(c, &bill)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) DeleteBillInCart(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	bill := model.Bill{}
	json.Unmarshal(bytes, &bill)
	if err != nil {
		panic(err)
	}
	// Game.Role = 1
	res := impl.dao.DeleteBillInCart(c, &bill)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) AssignCartForUser(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	order := model.OrderResponse{}
	json.Unmarshal(bytes, &order)
	if err != nil {
		panic(err)
	}

	res := impl.dao.AssignCartForUser(c, order)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) AddCommodity(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	commodity := model.CommodityResponse{}
	json.Unmarshal(bytes, &commodity)
	if err != nil {
		panic(err)
	}
	// Game.Role = 1
	res := impl.dao.AddCommodity(c, commodity)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) DeleteCommodity(c *gin.Context) {
	type DeleteCommodityRequest struct {
		Id int64 `json:"id"`
	}

	var req DeleteCommodityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	res := impl.dao.DeleteCommodity(c, req.Id)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) AddOption(c *gin.Context) {
	type AddOptionRequest struct {
		Option model.OptionResponse `json:"option"`
		ComId  int64                `json:"comId"`
	}

	var req AddOptionRequest
	// Game.Role = 1
	res := impl.dao.AddOption(c, req.Option, req.ComId)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) UpdateOption(c *gin.Context) {
	type UpdateOptionRequest struct {
		Option model.OptionResponse `json:"option"`
		ComId  int64                `json:"comId"`
	}

	var req UpdateOptionRequest
	// Game.Role = 1
	res := impl.dao.UpdateOption(c, req.Option, req.ComId)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) DeleteOption(c *gin.Context) {
	type DeleteOptionRequest struct {
		Id int64 `json:"id"`
	}

	var req DeleteOptionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// 处理错误
	}

	res := impl.dao.DeleteOption(c, req.Id)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) UpdateCommodity(c *gin.Context) {
	body := c.Request.Body
	bytes, err := ioutil.ReadAll(body)
	commodity := model.CommodityResponse{}
	json.Unmarshal(bytes, &commodity)
	if err != nil {
		panic(err)
	}
	// Game.Role = 1
	res := impl.dao.UpdateCommodity(c, commodity)
	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) GetAllCommodities(c *gin.Context) {

	res := impl.dao.GetAllCommodities(c)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}

func (impl LZControllerImpl) GetAllOrders(c *gin.Context) {

	res := impl.dao.GetAllOrders(c)

	c.JSON(200, map[string]interface{}{"code": 0, "msg": "", "count": 0, "data": res})
}
