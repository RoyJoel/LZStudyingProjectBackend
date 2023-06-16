package impl

import (
	"context"
	"encoding/json"

	"github.com/RoyJoel/LZStudyingProjectBackend/package/dao/impl"
	"github.com/RoyJoel/LZStudyingProjectBackend/package/model"
	"github.com/RoyJoel/LZStudyingProjectBackend/proto"
)

type LZRPCImpl struct {
	dao *impl.LZDaoImpl
}

func NewLZControllerImpl() *LZRPCImpl {
	return &LZRPCImpl{dao: impl.NewLZDaoImpl()}
}
func (impl *LZRPCImpl) AddPlayer(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	Id := request.GetId()
	LoginName := request.GetLoginName()
	Name := request.GetName()
	Icon := request.GetIcon()
	Sex := request.GetSex()
	Age := request.GetAge()
	Points := request.GetPoints()

	impl.dao.AddPlayer(ctx, model.Player{
		Id:        Id,
		LoginName: LoginName,
		Name:      Name,
		Icon:      Icon,
		Sex:       Sex,
		Age:       Age,
		Points:    Points,
	})
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: "true"}, nil
}

// func (impl *PlayerInfoRPCImpl) FindPlayerByKey(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
// 	key := request.GetInfoKey()
// 	Player := impl.dao.GetPlayerByUid(ctx, key)
// 	info, _ := json.Marshal(Player)
// 	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(info)}, nil
// }

func (impl *LZRPCImpl) UpdatePlayer(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	Id := request.GetId()
	LoginName := request.GetLoginName()
	Name := request.GetName()
	Icon := request.GetIcon()
	Sex := request.GetSex()
	Age := request.GetAge()
	Points := request.GetPoints()
	impl.dao.UpdatePlayer(ctx, model.PlayerResponse{
		Id:        Id,
		LoginName: LoginName,
		Name:      Name,
		Icon:      Icon,
		Sex:       Sex,
		Age:       Age,
		Points:    Points,
	})
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: "true"}, nil
}

// func (impl *PlayerInfoRPCImpl) DeleteById(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
// 	id := request.GetId()
// 	impl.dao.DeletePlayerById(ctx, id)
// 	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: "true"}, nil
// }

func (impl *LZRPCImpl) SearchPlayer(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	id := request.GetId()
	Players, _ := impl.dao.GetPlayerInfo(ctx, id)
	infos, _ := json.Marshal(Players)
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(infos)}, nil
}

func (impl *LZRPCImpl) GetPlayerInfo(ctx context.Context, req *proto.PlayerInfoRequest) (resp *proto.PlayerInfoResponse, err error) {
	id := req.GetId()
	Players, _ := impl.dao.GetPlayerInfo(ctx, id)
	infos, _ := json.Marshal(Players)
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(infos)}, nil
}

func (impl *LZRPCImpl) AddFriend(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	id := request.GetId()
	Players, _ := impl.dao.GetPlayerInfo(ctx, id)
	infos, _ := json.Marshal(Players)
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(infos)}, nil
}
func (impl *LZRPCImpl) DeleteFriend(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	id := request.GetId()
	Players, _ := impl.dao.GetPlayerInfo(ctx, id)
	infos, _ := json.Marshal(Players)
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(infos)}, nil
}
