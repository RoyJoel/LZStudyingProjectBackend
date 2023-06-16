package impl

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/RoyJoel/LZStudyingProjectBackend/package/cache"
	"github.com/RoyJoel/LZStudyingProjectBackend/package/middleware"
	"github.com/RoyJoel/LZStudyingProjectBackend/package/model"
	"github.com/RoyJoel/LZStudyingProjectBackend/package/utils"
	"github.com/RoyJoel/LZStudyingProjectBackendBackend/package/config"
	"gorm.io/gorm"
)

type LZDaoImpl struct {
	db    *gorm.DB
	cache *cache.LZCacheDAOImpl
}

func NewLZDaoImpl() *LZDaoImpl {
	return &LZDaoImpl{db: config.DB, cache: cache.NewLZCacheDAOImpl()}
}

func (impl *LZDaoImpl) AddPlayer(ctx context.Context, Player model.Player) (model.PlayerResponse, bool) {

	res := impl.SearchPlayer(ctx, Player.LoginName)
	if !res {
		newPlayer := model.Player{}
		newPlayer = model.Player{LoginName: Player.LoginName, Password: Player.Password, Name: Player.Name, Icon: Player.Icon, Sex: Player.Sex, Age: Player.Age, Points: Player.Points}
		impl.db.Create(&newPlayer)
		return model.PlayerResponse{Id: newPlayer.Id, LoginName: newPlayer.LoginName, Password: newPlayer.Password, Name: newPlayer.Name, Icon: newPlayer.Icon, Sex: newPlayer.Sex, Age: newPlayer.Age, Points: newPlayer.Points}, res
	} else {
		var player model.Player
		impl.db.Where("id = ?", Player.Id).First(&player)
		return model.PlayerResponse{Id: player.Id, LoginName: player.LoginName, Password: Player.Password, Name: player.Name, Icon: player.Icon, Sex: player.Sex, Age: player.Age, Points: player.Points}, res
	}
}

func (impl *LZDaoImpl) SignUp(ctx context.Context, user model.User) (model.User, bool) {
	playerInfo := model.Player{Id: user.Id, LoginName: user.LoginName, Password: user.Password, Name: user.Name, Icon: user.Icon, Sex: user.Sex, Age: user.Age, Points: user.Points}
	player, res := impl.AddPlayer(ctx, playerInfo)
	friends := []model.PlayerResponse{}
	if res == false {
		relationship := model.Relationship{Player1Id: player.Id, Player2Id: player.Id}
		impl.db.Create(relationship)
		friends = append(friends, player)
	} else {
		friends = user.Friends
	}
	cart := impl.AddCartForPlayer(ctx, player.Id)
	// 生成 token string
	tokenStr, _ := middleware.GenerateToken(user.LoginName, user.Password)
	user = model.User{Id: player.Id, LoginName: player.LoginName, Password: user.Password, Name: player.Name, Icon: player.Icon, Sex: player.Sex, Age: player.Age, Points: player.Points, Friends: friends, AllLikedMusic: utils.IntMatrix{}, AllOrders: utils.IntMatrix{}, Addresss: utils.IntMatrix{}, Cart: cart, DefaultAddress: model.AddressResponse{}, Token: tokenStr}
	return user, res
}
func (impl *LZDaoImpl) AddCartForPlayer(ctx context.Context, playerId int64) int64 {
	cart := model.Cart{}
	impl.db.Where("player_id = ?", playerId).Delete(&cart)
	newOrder := model.Order{PlayerId: playerId, State: 0}
	impl.db.Create(&newOrder)
	newCart := model.Cart{PlayerId: playerId, OrderId: newOrder.Id}
	impl.db.Create(&newCart)
	return newOrder.Id
}

func (impl *LZDaoImpl) UpdateUser(ctx context.Context, user model.User) model.User {
	Player := model.PlayerResponse{Id: user.Id, LoginName: user.LoginName, Password: user.Password, Name: user.Name, Icon: user.Icon, Sex: user.Sex, Age: user.Age, Points: user.Points}
	playerResponse := impl.UpdatePlayer(ctx, Player)

	return model.User{Id: playerResponse.Id, LoginName: playerResponse.LoginName, Password: playerResponse.Password, Name: playerResponse.Name, Icon: playerResponse.Icon, Sex: playerResponse.Sex, Age: playerResponse.Age, Points: playerResponse.Points, Friends: user.Friends, AllLikedMusic: user.AllLikedMusic, AllOrders: user.AllOrders, Addresss: user.Addresss, Cart: user.Cart, Token: user.Token}
}

func (impl *LZDaoImpl) SignIn(ctx context.Context, userLoginName string, userPassword string) (*model.User, error) {
	player, _ := impl.GetPlayerInfoByLoginNameAndPassword(ctx, userLoginName, userPassword)
	if player != nil {
		friends := impl.GetAllFriends(ctx, player.Id)
		allLikedMusic := impl.GetMyMusic(ctx, player.Id)
		allOrders := impl.GetMyOrders(ctx, player.Id)
		addresss := impl.GetMyAddresss(ctx, player.Id)
		cart := impl.GetMyCart(ctx, player.Id)
		defaultAddress := impl.GetMyDefaultAddress(ctx, player.Id)
		// 生成 token string
		tokenStr, error := middleware.GenerateToken(player.LoginName, player.Password)
		if error != nil {
			return nil, error
		} else {
			user := model.User{Id: player.Id, LoginName: player.LoginName, Password: player.Password, Name: player.Name, Icon: player.Icon, Sex: player.Sex, Age: player.Age, Friends: friends, Points: player.Points, AllLikedMusic: allLikedMusic, AllOrders: allOrders, Addresss: addresss, Cart: cart, DefaultAddress: defaultAddress, Token: tokenStr}
			return &user, nil
		}
	} else {
		return nil, errors.New("no such user")
	}
}

func (impl *LZDaoImpl) ResetPassword(ctx context.Context, userLoginName string, userPassword string) bool {
	var Player model.Player
	impl.db.Where("login_name = ?", userLoginName).First(&Player)
	if Player.LoginName == userLoginName {
		player := model.Player{Password: userPassword}
		impl.db.Where("login_name = ?", Player.LoginName).Updates(&player)
		return true
	}
	return false
}

func (impl *LZDaoImpl) UpdatePlayer(ctx context.Context, player model.PlayerResponse) model.PlayerResponse {
	Player := model.Player{Id: player.Id, LoginName: player.LoginName, Password: player.Password, Name: player.Name, Icon: player.Icon, Sex: player.Sex, Age: player.Age, Points: player.Points}
	impl.db.Where("id = ?", player.Id).Updates(&Player)
	impl.db.First(&Player, "id = ?", player.Id)
	return model.PlayerResponse{Id: Player.Id, LoginName: Player.LoginName, Password: Player.Password, Name: Player.Name, Icon: Player.Icon, Sex: Player.Sex, Age: Player.Age, Points: Player.Points}
}

func (impl *LZDaoImpl) SearchPlayer(ctx context.Context, loginName string) bool {
	var Player model.Player
	impl.db.Where("login_name = ?", loginName).First(&Player)
	return Player.LoginName == loginName
}

func (impl *LZDaoImpl) GetPlayerInfo(ctx context.Context, id int64) (*model.PlayerResponse, error) {
	var Player model.Player
	impl.db.Where("id = ?", id).First(&Player)
	if Player.Id == id {
		return &model.PlayerResponse{Id: Player.Id, LoginName: Player.LoginName, Password: Player.Password, Name: Player.Name, Icon: Player.Icon, Sex: Player.Sex, Age: Player.Age, Points: Player.Points}, nil
	}
	return nil, errors.New("no such player")
}

func (impl *LZDaoImpl) GetPlayerInfoByLoginName(ctx context.Context, loginName string) (*model.PlayerResponse, error) {
	var Player model.Player
	impl.db.Where("login_name = ?", loginName).First(&Player)
	if Player.LoginName == loginName {
		return &model.PlayerResponse{Id: Player.Id, LoginName: Player.LoginName, Password: Player.Password, Name: Player.Name, Icon: Player.Icon, Sex: Player.Sex, Age: Player.Age, Points: Player.Points}, nil
	}
	return nil, errors.New("no such player")
}

func (impl *LZDaoImpl) GetPlayerInfoByLoginNameAndPassword(ctx context.Context, loginName string, password string) (*model.PlayerResponse, error) {
	var Player model.Player
	fmt.Println(password)

	impl.db.Where("login_name = ?", loginName).First(&Player)
	if Player.LoginName == loginName {
		// 将密码转换为字节数组
		passwordBytes := []byte(Player.Password)

		// 计算 SHA256 哈希值
		hash := sha256.Sum256(passwordBytes)

		// 将哈希值转换为字符串并打印输出
		authedPassword := hex.EncodeToString(hash[:])
		fmt.Println(authedPassword)
		if authedPassword == password {
			return &model.PlayerResponse{Id: Player.Id, LoginName: Player.LoginName, Password: Player.Password, Name: Player.Name, Icon: Player.Icon, Sex: Player.Sex, Age: Player.Age, Points: Player.Points}, nil
		} else {
			return nil, errors.New("no such account or password")
		}
	}
	return nil, errors.New("no such user")
}

func (impl *LZDaoImpl) GetAllFriends(ctx context.Context, id int64) []model.PlayerResponse {
	var friends []model.Player
	var friendResponses []model.PlayerResponse
	var relationship1 []model.Relationship
	impl.db.Find(&relationship1, "player1_id", id)
	for _, friend := range relationship1 {
		if friend.Player1Id != friend.Player2Id {
			var player model.Player
			impl.db.First(&player, "id", friend.Player2Id)
			friends = append(friends, player)
		}
	}
	var relationship2 []model.Relationship
	impl.db.Find(&relationship2, "player2_id", id)
	for _, friend := range relationship2 {
		var player model.Player
		impl.db.First(&player, "id", friend.Player1Id)
		friends = append(friends, player)
	}

	for _, friend := range friends {
		friendResponse := model.PlayerResponse{Id: friend.Id, LoginName: friend.LoginName, Password: friend.Password, Name: friend.Name, Icon: friend.Icon, Sex: friend.Sex, Age: friend.Age, Points: friend.Points}
		friendResponses = append(friendResponses, friendResponse)
	}
	return friendResponses
}

func (impl *LZDaoImpl) AddFriend(ctx context.Context, relationship model.Relationship) []model.PlayerResponse {
	res := impl.SearchFriend(ctx, relationship)
	if !res {
		impl.db.Create(&relationship)
	}
	return impl.GetAllFriends(ctx, relationship.Player1Id)
}

func (impl *LZDaoImpl) UpdateFriend(ctx context.Context, relationship model.Relationship) []model.PlayerResponse {
	res1 := impl.SearchFriend(ctx, model.Relationship{Player1Id: relationship.Player2Id, Player2Id: relationship.Player1Id})
	res2 := impl.SearchFriend(ctx, relationship)
	if !res1 && !res2 {
		impl.db.Create(&relationship)
	}

	return impl.GetAllFriends(ctx, relationship.Player1Id)
}

func (impl *LZDaoImpl) UpdateFriendDB(ctx context.Context, userId int64, friends []model.PlayerResponse) []model.PlayerResponse {
	relationships := make([]model.Relationship, 0)
	for _, friend := range friends {
		relationship := model.Relationship{Player1Id: userId, Player2Id: friend.Id}
		relationships = append(relationships, relationship)
	}
	var results []model.Relationship

	m := make(map[model.Relationship]bool)
	impl.db.Where("player1_id = ? OR player2_id = ?", userId, userId).Find(&results)
	for _, relat := range relationships {
		rerat := model.Relationship{Player1Id: relat.Player2Id, Player2Id: relat.Player1Id}
		m[relat] = true
		m[rerat] = true
	}
	for _, res := range results {
		rerat := model.Relationship{Player1Id: res.Player2Id, Player2Id: res.Player1Id}
		if !m[res] && !m[rerat] {
			impl.db.Where("player1_id = ? AND player2_id = ?", res.Player1Id, res.Player2Id).Delete(&res)
			impl.db.Where("player1_id = ? AND player2_id = ?", rerat.Player1Id, rerat.Player2Id).Delete(&rerat)
		}
	}
	return impl.GetAllFriends(ctx, userId)
}

func (impl *LZDaoImpl) DeleteFriend(ctx context.Context, relationship model.Relationship) []model.PlayerResponse {
	if impl.SearchFriend(ctx, relationship) {
		impl.db.Where("player1_id = ? AND player2_id = ?", relationship.Player1Id, relationship.Player2Id).Delete(&relationship)
		impl.db.Where("player1_id = ? AND player2_id = ?", relationship.Player2Id, relationship.Player1Id).Delete(&relationship)
	}
	return impl.GetAllFriends(ctx, relationship.Player1Id)
}

func (impl *LZDaoImpl) SearchFriend(ctx context.Context, relationship model.Relationship) bool {
	var Relationship1 model.Relationship
	var Relationship2 model.Relationship
	impl.db.Where("player1_id = ? AND player2_id = ?", relationship.Player1Id, relationship.Player2Id).First(&Relationship1)
	impl.db.Where("player1_id = ? AND player2_id = ?", relationship.Player2Id, relationship.Player1Id).First(&Relationship2)
	return Relationship1 == relationship || Relationship2 == relationship

}

func (impl *LZDaoImpl) GetMyMusic(ctx context.Context, playerId int64) utils.IntMatrix {
	musicUsers := make([]model.MusicUser, 0)
	musicResponses := utils.IntMatrix{}
	impl.db.Where("member_id = ?", playerId).Find(&musicUsers)

	for _, musicUser := range musicUsers {
		musicResponses = append(musicResponses, musicUser.MusicId)
	}
	return musicResponses
}

func (impl *LZDaoImpl) GetMusicInfo(ctx context.Context, musicId int64) model.MusicResponse {
	music := model.Music{}
	impl.db.Where("id = ?", musicId).Find(&music)
	AnimeInfo := impl.GetAnimeInfo(ctx, music.AnimeInfoId)
	musicResponse := model.MusicResponse{Id: music.Id, PlayUrl: music.PlayUrl, Type: music.Type, Recommend: music.Recommend, Atime: music.Atime, Author: music.Author, AnimeInfo: AnimeInfo}

	return musicResponse
}
func (impl *LZDaoImpl) GetAnimeInfo(ctx context.Context, Id string) model.AnimeInfo {
	animeInfo := model.AnimeInfo{}
	impl.db.Where("id = ?", Id).First(&animeInfo)

	return animeInfo
}

func (impl *LZDaoImpl) UpdateMusic(ctx context.Context, MusicUser model.MusicUser) {
	clubMember := model.MusicUser{}
	impl.db.Where("club_id = ? And member_id = ?", MusicUser.MusicId, MusicUser.PlayerId).First(&clubMember)
	if clubMember.MusicId != MusicUser.MusicId || clubMember.PlayerId != MusicUser.PlayerId {
		impl.db.Create(MusicUser)
	}
	return
}

func (impl *LZDaoImpl) UpdateClubDB(ctx context.Context, userId int64, music utils.IntMatrix) utils.IntMatrix {

	var results utils.IntMatrix

	m := make(map[int64]bool)
	impl.db.Where("member_id = ?", userId).Find(&results)

	for _, item := range music {
		m[item] = true
	}

	for _, item := range results {
		if !m[item] {
			club := model.MusicUser{MusicId: item, PlayerId: userId}
			impl.db.Where("club_id = ?", item).Delete(&club)
		}
	}
	return impl.GetMyMusic(ctx, userId)
}

func (impl *LZDaoImpl) UpdateOrder(ctx context.Context, Order model.OrderResponse) model.OrderResponse {
	order := model.Order{Id: Order.Id, ShippingAddressId: Order.ShippingAddress.Id, PlayerId: Order.PlayerId, State: Order.State, Payment: Order.Payment, CreatedTime: Order.CreatedTime, PayedTime: Order.PayedTime, CompletedTime: Order.CompletedTime}
	impl.db.Save(&order)
	return impl.GetOrderInfo(ctx, order.Id)
}

func (impl *LZDaoImpl) DeleteOrder(ctx context.Context, id int64) bool {

	order := model.Order{}
	impl.db.Where("id = ?", id).Delete(&order)
	bill := model.Bill{}
	impl.db.Where("order_id", id).Delete(&bill)

	return true
}
func (impl *LZDaoImpl) GetOrderInfo(ctx context.Context, orderId int64) model.OrderResponse {
	order := model.Order{}
	impl.db.Where("id = ?", orderId).First(&order)

	bills := impl.GetBillInfos(ctx, order.Id)
	shippingAddress := impl.GetAddressInfo(ctx, order.ShippingAddressId)

	return model.OrderResponse{Id: order.Id, Bills: bills, ShippingAddress: shippingAddress, CreatedTime: order.CreatedTime, PayedTime: order.PayedTime, CompletedTime: order.CompletedTime, State: order.State}
}

func (impl *LZDaoImpl) GetOrderInfosByUserId(ctx context.Context, playerId int64) []model.OrderResponse {
	orders := make([]model.Order, 0)
	impl.db.Where("player_id = ?", playerId).Find(&orders)

	orderResponses := make([]model.OrderResponse, 0)
	for _, order := range orders {
		bills := impl.GetBillInfos(ctx, order.Id)
		shippingAddress := impl.GetAddressInfo(ctx, order.ShippingAddressId)

		orderResponse := model.OrderResponse{Id: order.Id, Bills: bills, ShippingAddress: shippingAddress, CreatedTime: order.CreatedTime, PayedTime: order.PayedTime, CompletedTime: order.CompletedTime, State: order.State}

		orderResponses = append(orderResponses, orderResponse)
	}
	fmt.Println(orderResponses)
	return orderResponses
}

func (impl *LZDaoImpl) AddBillToCart(ctx context.Context, bill *model.Bill) model.OrderResponse {

	Bill := model.Bill{}
	impl.db.Where("com_id = ? AND option_id = ? AND order_id = ?", bill.ComId, bill.OptionId, bill.OrderId).First(&Bill)
	newBill := model.Bill{}
	cart := impl.GetOrderInfo(ctx, bill.OrderId)
	if Bill.ComId == bill.ComId && Bill.OptionId == bill.OptionId && Bill.OrderId == bill.OrderId {
		Bill.Quantity += bill.Quantity
		impl.db.Save(&Bill)
	} else {
		impl.db.Create(&newBill)
		bill.Id = newBill.Id
		impl.db.Save(&bill)
		billInfo := impl.GetBillInfo(ctx, bill.Id)
		cart.Bills = append(cart.Bills, billInfo)
	}

	return *&cart
}

func (impl *LZDaoImpl) DeleteBillInCart(ctx context.Context, bill *model.Bill) model.OrderResponse {

	if impl.SearchBill(ctx, bill.Id) {
		impl.db.Where("id = ?", bill.Id).Delete(&bill)
	}

	cart := impl.GetOrderInfo(ctx, bill.OrderId)
	return cart
}

func (impl *LZDaoImpl) AssignCartForUser(ctx context.Context, order model.OrderResponse) int64 {

	Order := model.Order{Id: order.Id, ShippingAddressId: order.ShippingAddress.Id, PlayerId: order.PlayerId, State: order.State, Payment: order.Payment, CreatedTime: order.CreatedTime, PayedTime: order.PayedTime, CompletedTime: order.CompletedTime}
	impl.db.Save(&Order)
	for _, bill := range order.Bills {
		billResponse := model.Bill{Id: bill.Id, ComId: bill.Com.Id, Quantity: bill.Quantity, OptionId: bill.Option.Id, OrderId: order.Id}
		impl.AddBill(ctx, billResponse)
	}
	cart := impl.AddCartForPlayer(ctx, order.PlayerId)

	return cart

}

func (impl *LZDaoImpl) GetAllCommodities(ctx context.Context) []model.CommodityResponse {
	commodities := make([]model.Commodity, 0)
	impl.db.Find(&commodities)

	CommodityResponses := make([]model.CommodityResponse, 0)
	for _, commodity := range commodities {
		CommodityResponse := impl.GetCommodityInfo(ctx, commodity.Id)
		CommodityResponses = append(CommodityResponses, CommodityResponse)
	}

	return CommodityResponses
}

func (impl *LZDaoImpl) GetAllOrders(ctx context.Context) []model.OrderResponse {
	orders := make([]model.Order, 0)
	impl.db.Find(&orders)

	orderResponses := make([]model.OrderResponse, 0)
	for _, orderResponse := range orders {
		orderResponse := impl.GetOrderInfo(ctx, orderResponse.Id)
		orderResponses = append(orderResponses, orderResponse)
	}

	return orderResponses
}

func (impl *LZDaoImpl) GetBillInfos(ctx context.Context, orderId int64) []model.BillResponse {
	bills := make([]model.Bill, 0)
	impl.db.Where("order_id = ?", orderId).Find(&bills)

	billResponses := make([]model.BillResponse, 0)
	for _, bill := range bills {
		commodity := impl.GetCommodityInfo(ctx, bill.ComId)
		option := impl.GetOptionInfo(ctx, bill.OptionId)
		billResponse := model.BillResponse{Id: bill.Id, Com: commodity, Quantity: bill.Quantity, Option: option}
		billResponses = append(billResponses, billResponse)
	}

	return billResponses
}

func (impl *LZDaoImpl) GetOptionInfo(ctx context.Context, optionId int64) model.OptionResponse {
	option := model.Option{}
	impl.db.Where("id = ?", optionId).First(&option)

	return model.OptionResponse{Id: option.Id, Image: option.Image, Intro: option.Intro, Price: option.Price, Inventory: option.Inventory}
}

func (impl *LZDaoImpl) GetAddressInfo(ctx context.Context, addressId int64) model.AddressResponse {
	address := model.Address{}
	impl.db.Where("id = ?", addressId).First(&address)
	addressResponse := model.AddressResponse{Id: address.Id, Name: address.Name, Sex: address.Sex, PhoneNumber: address.PhoneNumber, Province: address.Province, City: address.City, Area: address.Area, DetailAddress: address.DetailedAddress, IsDefault: address.IsDefault}
	return addressResponse
}

func (impl *LZDaoImpl) AddOrder(ctx context.Context, Order model.OrderResponse) int64 {

	newOrder := model.Order{}
	impl.db.Create(&newOrder)
	order := model.Order{Id: newOrder.Id, ShippingAddressId: Order.ShippingAddress.Id, PlayerId: Order.PlayerId, State: Order.State, Payment: Order.Payment, CreatedTime: Order.CreatedTime, PayedTime: Order.PayedTime, CompletedTime: Order.CompletedTime}
	impl.db.Save(&order)
	for _, bill := range Order.Bills {
		billResponse := model.Bill{Id: bill.Id, ComId: bill.Com.Id, Quantity: bill.Quantity, OptionId: bill.Option.Id, OrderId: order.Id}
		impl.AddBill(ctx, billResponse)
	}
	return order.Id
}

func (impl *LZDaoImpl) SearchOrder(ctx context.Context, id int64) bool {
	var Order model.Order
	impl.db.Where("id = ?", id).First(&Order)
	return Order.Id == id
}

func (impl *LZDaoImpl) AddAddress(ctx context.Context, address model.Address) model.AddressResponse {

	newAddress := model.Address{}
	impl.db.Create(&newAddress)
	address.Id = newAddress.Id
	impl.db.Save(&address)
	return model.AddressResponse{Id: address.Id, Name: address.Name, Sex: address.Sex, PhoneNumber: address.PhoneNumber, Province: address.Province, City: address.City, Area: address.Area, DetailAddress: address.DetailedAddress, IsDefault: address.IsDefault}
}

func (impl *LZDaoImpl) SearchAddress(ctx context.Context, id int64) bool {
	var Address model.Address
	impl.db.Where("id = ?", id).First(&Address)
	return Address.Id == id
}

func (impl *LZDaoImpl) UpdateAddress(ctx context.Context, updatingAddress model.Address) model.AddressResponse {
	impl.db.Save(&updatingAddress)
	addresss := impl.GetMyAddresss(ctx, updatingAddress.PlayerId)
	for _, address := range addresss {
		if address != updatingAddress.Id {
			addressInfo := impl.GetAddressInfo(ctx, address)
			addressInfo.IsDefault = false
			AddressInfo := model.Address{Id: addressInfo.Id, PlayerId: updatingAddress.PlayerId, Name: addressInfo.Name, PhoneNumber: addressInfo.PhoneNumber, Sex: addressInfo.Sex, Province: addressInfo.Province, City: addressInfo.City, Area: addressInfo.Area, DetailedAddress: addressInfo.DetailAddress, IsDefault: addressInfo.IsDefault}
			impl.db.Save(&AddressInfo)
		}
	}
	return impl.GetAddressInfo(ctx, updatingAddress.Id)
}

func (impl *LZDaoImpl) DeleteAddress(ctx context.Context, id int64) bool {
	address := model.Address{}
	impl.db.Where("id = ?", id).Delete(&address)
	return true
}

func (impl *LZDaoImpl) AddBill(ctx context.Context, bill model.Bill) model.Bill {

	newBill := model.Bill{}
	impl.db.Create(&newBill)
	bill.Id = newBill.Id
	impl.db.Save(&bill)
	return bill
}

func (impl *LZDaoImpl) UpdateBill(ctx context.Context, userId int64, Bill model.Bill) model.BillResponse {
	impl.db.Save(&Bill)
	return impl.GetBillInfo(ctx, Bill.Id)
}

func (impl *LZDaoImpl) GetBillInfo(ctx context.Context, billId int64) model.BillResponse {
	bill := model.Bill{}
	impl.db.Where("id = ?", billId).First(&bill)
	com := impl.GetCommodityInfo(ctx, bill.ComId)
	option := impl.GetOptionInfo(ctx, bill.OptionId)
	billResponse := model.BillResponse{Id: bill.Id, Com: com, Quantity: bill.Quantity, Option: option}
	return billResponse
}
func (impl *LZDaoImpl) SearchBill(ctx context.Context, id int64) bool {
	var bill model.Bill
	impl.db.Where("id = ?", id).First(&bill)
	return bill.Id == id
}

func (impl *LZDaoImpl) AddCommodity(ctx context.Context, Commodity model.CommodityResponse) model.CommodityResponse {

	newCommodity := model.Commodity{}
	impl.db.Create(&newCommodity)
	newCommodity.Name = Commodity.Name
	newCommodity.Intro = Commodity.Intro
	newCommodity.Cag = Commodity.Cag
	newCommodity.State = Commodity.State
	for _, option := range Commodity.Options {
		impl.AddOption(ctx, option, newCommodity.Id)
	}
	impl.db.Save(&newCommodity)
	Commodity.Id = newCommodity.Id
	return Commodity
}

func (impl *LZDaoImpl) AddOption(ctx context.Context, option model.OptionResponse, comId int64) []model.OptionResponse {
	newOption := model.Option{}
	impl.db.Create(&newOption)
	newOption.Image = option.Image
	newOption.Intro = option.Intro
	newOption.Price = option.Price
	newOption.Inventory = option.Inventory
	newOption.ComId = comId
	println(newOption.ComId)
	impl.db.Save(&newOption)
	return impl.GetOptionsForCommidity(ctx, comId)
}

func (impl *LZDaoImpl) DeleteCommodity(ctx context.Context, id int64) []model.CommodityResponse {
	commodity := model.Commodity{}
	impl.db.Where("id = ?", id).Delete(&commodity)
	option := model.Option{}
	impl.db.Where("com_id = ?", id).Delete(&option)
	return impl.GetAllCommodities(ctx)
}

func (impl *LZDaoImpl) DeleteOption(ctx context.Context, id int64) []model.OptionResponse {
	Option := model.Option{}
	OptionRequest := model.Option{}
	impl.db.Where("id = ?", id).Find(&Option)
	impl.db.Where("id = ?", id).Delete(&OptionRequest)
	return impl.GetOptionsForCommidity(ctx, Option.ComId)
}

func (impl *LZDaoImpl) UpdateCommodity(ctx context.Context, CommodityResponse model.CommodityResponse) model.CommodityResponse {
	Commodity := model.Commodity{Id: CommodityResponse.Id, Name: CommodityResponse.Name, Intro: CommodityResponse.Intro, Cag: CommodityResponse.Cag, State: CommodityResponse.State}
	impl.db.Save(&Commodity)
	for _, option := range CommodityResponse.Options {
		impl.UpdateOption(ctx, option, CommodityResponse.Id)
	}
	return impl.GetCommodityInfo(ctx, Commodity.Id)
}

func (impl *LZDaoImpl) UpdateOption(ctx context.Context, OptionResponse model.OptionResponse, comId int64) []model.OptionResponse {
	Option := model.Option{Id: OptionResponse.Id, Image: OptionResponse.Image, Intro: OptionResponse.Intro, Price: OptionResponse.Price, Inventory: OptionResponse.Inventory, ComId: comId}
	impl.db.Save(&Option)
	return impl.GetOptionsForCommidity(ctx, comId)
}

func (impl *LZDaoImpl) GetCommodityInfo(ctx context.Context, CommodityId int64) model.CommodityResponse {
	Commodity := model.Commodity{}
	impl.db.Where("id = ?", CommodityId).First(&Commodity)
	orders := impl.GetOrderNumForCommodity(ctx, Commodity.Id)
	options := impl.GetOptionsForCommidity(ctx, Commodity.Id)
	return model.CommodityResponse{Id: Commodity.Id, Name: Commodity.Name, Intro: Commodity.Intro, Cag: Commodity.Cag, Orders: orders, Options: options, State: Commodity.State}
}

func (impl *LZDaoImpl) GetCommoditySimpleInfo(ctx context.Context, CommodityId int64) model.Commodity {
	Commodity := model.Commodity{}
	impl.db.Where("id = ?", CommodityId).First(&Commodity)
	return Commodity
}

func (impl *LZDaoImpl) GetOrderNumForCommodity(ctx context.Context, CommodityId int64) int64 {
	bills := make([]model.Bill, 0)
	impl.db.Where("id = ?", CommodityId).Find(&bills)
	var num int64
	for _, bill := range bills {
		num += bill.Quantity
	}
	return num
}

func (impl *LZDaoImpl) GetOptionsForCommidity(ctx context.Context, CommodityId int64) []model.OptionResponse {
	options := make([]model.Option, 0)
	impl.db.Where("com_id = ?", CommodityId).Find(&options)

	optionResponses := make([]model.OptionResponse, 0)
	for _, option := range options {
		optionResponse := model.OptionResponse{Id: option.Id, Image: option.Image, Intro: option.Intro, Price: option.Price, Inventory: option.Inventory}
		optionResponses = append(optionResponses, optionResponse)
	}
	return optionResponses
}

func (impl *LZDaoImpl) SearchCommodity(ctx context.Context, id int64) bool {
	var Commodity model.Commodity
	impl.db.Where("id = ?", id).First(&Commodity)
	return Commodity.Id == id
}

func (impl *LZDaoImpl) GetMyOrders(ctx context.Context, playerId int64) utils.IntMatrix {
	orders := make([]model.Order, 0)
	orderResponses := utils.IntMatrix{}
	impl.db.Where("player_id = ?", playerId).Find(&orders)

	for _, order := range orders {
		orderResponses = append(orderResponses, order.Id)
	}
	return orderResponses
}

func (impl *LZDaoImpl) GetMyAddresss(ctx context.Context, playerId int64) utils.IntMatrix {
	addresss := make([]model.Address, 0)
	addressResponses := utils.IntMatrix{}
	impl.db.Where("player_id = ?", playerId).Find(&addresss)

	for _, order := range addresss {
		addressResponses = append(addressResponses, order.Id)
	}
	return addressResponses
}
func (impl *LZDaoImpl) GetMyDefaultAddress(ctx context.Context, playerId int64) model.AddressResponse {
	address := model.Address{}
	impl.db.Where("player_id = ? AND is_default = ?", playerId, 1).Find(&address)
	addressResponse := model.AddressResponse{Id: address.Id, Name: address.Name, Sex: address.Sex, PhoneNumber: address.PhoneNumber, Province: address.Province, City: address.City, Area: address.Area, DetailAddress: address.DetailedAddress, IsDefault: address.IsDefault}
	return addressResponse
}

func (impl *LZDaoImpl) GetMyCart(ctx context.Context, playerId int64) int64 {
	cart := model.Cart{}
	impl.db.Where("player_id = ?", playerId).First(&cart)
	return cart.OrderId
}
