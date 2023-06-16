package impl

import (
	"context"
	"encoding/json"

	"github.com/RoyJoel/LZStudyingProject/package/dao/impl"
	"github.com/RoyJoel/LZStudyingProject/package/model"
	"github.com/RoyJoel/LZStudyingProject/proto"
)

type TennisMomentRPCImpl struct {
	dao *impl.TennisMomentDaoImpl
}

func NewTennisMomentControllerImpl() *TennisMomentRPCImpl {
	return &TennisMomentRPCImpl{dao: impl.NewTennisMomentDaoImpl()}
}
func (impl *TennisMomentRPCImpl) AddPlayer(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
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

func (impl *TennisMomentRPCImpl) UpdatePlayer(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
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

func (impl *TennisMomentRPCImpl) SearchPlayer(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	id := request.GetId()
	Players, _ := impl.dao.GetPlayerInfo(ctx, id)
	infos, _ := json.Marshal(Players)
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(infos)}, nil
}

func (impl *TennisMomentRPCImpl) GetPlayerInfo(ctx context.Context, req *proto.PlayerInfoRequest) (resp *proto.PlayerInfoResponse, err error) {
	id := req.GetId()
	Players, _ := impl.dao.GetPlayerInfo(ctx, id)
	infos, _ := json.Marshal(Players)
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(infos)}, nil
}

func (impl *TennisMomentRPCImpl) AddFriend(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	id := request.GetId()
	Players, _ := impl.dao.GetPlayerInfo(ctx, id)
	infos, _ := json.Marshal(Players)
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(infos)}, nil
}
func (impl *TennisMomentRPCImpl) DeleteFriend(ctx context.Context, request *proto.PlayerInfoRequest) (*proto.PlayerInfoResponse, error) {
	id := request.GetId()
	Players, _ := impl.dao.GetPlayerInfo(ctx, id)
	infos, _ := json.Marshal(Players)
	return &proto.PlayerInfoResponse{Code: 0, Msg: "", Count: 1, Data: string(infos)}, nil
}
